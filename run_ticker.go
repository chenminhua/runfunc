package runfunc

import (
	"time"
)


type JobResult int

const (
	Job_Result_Success JobResult = iota
	Job_Result_Failure
)

// 每隔多久运行一次
func runWithInterval(job Job, interval int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	for {
		ticker.Reset(time.Duration(interval) * time.Second)
		job.Do()
		<-ticker.C
	}
}

type Job interface {
	Do() JobResult
}
