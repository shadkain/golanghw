package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
)

func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}

	in := make(chan interface{}, 1)
	for _, fun := range jobs {
		wg.Add(1)
		out := make(chan interface{}, 1)
		go runWorker(fun, in, out, wg)
		in = out
	}

	wg.Wait()
}

func runWorker(fun job, in, out chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	fun(in, out)
	close(out)
}

func SingleHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	for val := range in {
		data := strconv.Itoa(val.(int))

		chan1 := make(chan string, 1)
		chan2 := make(chan string, 1)

		go signCrc32(data, chan1)
		go signCrc32(DataSignerMd5(data), chan2)

		wg.Add(1)
		go func(chan1, chan2 chan string, out chan interface{}) {
			defer wg.Done()

			crc32 := <-chan1
			crc32md5 := <-chan2
			res := crc32 + "~" + crc32md5

			out <- res
		}(chan1, chan2, out)
	}

	wg.Wait()
}

func signCrc32(data string, res chan string) {
	res <- DataSignerCrc32(data)
}

func MultiHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	for val := range in {
		data := val.(string)

		chans := make([]chan string, 0, 6)

		for i := 0; i < 6; i++ {
			chans = append(chans, make(chan string, 1))
			go signCrc32(strconv.Itoa(i)+data, chans[i])
		}

		wg.Add(1)
		go func(chans []chan string, out chan interface{}) {
			defer wg.Done()

			res := ""
			for i := 0; i < 6; i++ {
				res += <-chans[i]
			}

			out <- res
		}(chans, out)
	}

	wg.Wait()
}

func CombineResults(in, out chan interface{}) {
	results := make([]string, 0, 100)

	for val := range in {
		data := val.(string)
		results = append(results, data)
	}

	sort.Strings(results)

	res := strings.Join(results, "_")

	out <- res
}
