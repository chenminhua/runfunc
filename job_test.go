package runfunc

import (
	"testing"
)

type printJob struct {}

//Do() JobResult
func (pj *printJob) Do() JobResult {

	println("hello run ticker")
	return Job_Result_Success

}

func TestRun(t *testing.T) {
	runWithInterval(&printJob{},3)
}
