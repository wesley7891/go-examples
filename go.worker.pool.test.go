package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Id int
	RandNum int
}

type Result struct {
	job *Job // Job Struct
	sum int
}

// 创建工作池
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < num; i ++ {
		go func(jobChan chan *Job, resultChan chan *Result){
			for job := range(jobChan) {
				r_num := job.RandNum

				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}

				r := &Result {
					job: job,
					sum : sum,
				}

				// 得到的结果放到Result Chan里面
				resultChan <- r
			}

		}(jobChan, resultChan)
	}
}

func main() {
	
	// 1.job管道
	jobChan := make(chan *Job, 128)
	
	// 2.Result管道
	resultChan := make(chan *Result, 128)
	
	// 3.创建工作池
	createPool(64, jobChan, resultChan)

	// 4.打开打印的协程
	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("Job.Id=%v randNum:%v result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}(resultChan)

	var id int
	// 创建一个循环Job, 输入到管道中
	for {
		id ++
		r_num := rand.Int()
		job := &Job {
			Id: id,
			RandNum: r_num,
		}

		jobChan <- job
	}
}
