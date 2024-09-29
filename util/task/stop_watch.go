package task

import (
	"fmt"
	"github.com/mobaijaavaer/golang-tools/collections/impl"
	"github.com/mobaijaavaer/golang-tools/util/strings"
	strings2 "strings"
	"time"
)

type StopWatch struct {
	Id           string
	KeepTaskList bool
	TaskList     impl.LinkList[TaskInfo]
	CurrentTask  string
	StartTime    int
	LastTask     *TaskInfo
	TaskCount    int
	TotalTime    int
}

func NewStopWatch(id string) *StopWatch {
	return &StopWatch{
		Id:           id,
		KeepTaskList: true,
		TaskList:     *impl.NewLinkList[TaskInfo](),
	}
}

func (this *StopWatch) Start(taskName string) {
	if !strings.IsBlank(this.CurrentTask) {
		panic("task has been running")
	}
	this.CurrentTask = taskName
	this.StartTime = time.Now().Nanosecond()
}

func (this *StopWatch) Stop() {
	if strings.IsBlank(this.CurrentTask) {
		panic("task name cannot be blank")
	}
	lastTime := time.Now().Nanosecond() - this.StartTime
	this.TotalTime += lastTime
	this.LastTask = &TaskInfo{
		Name: this.CurrentTask,
		time: lastTime,
	}
	if this.KeepTaskList {
		this.TaskList.Add(*this.LastTask)
	}
	this.TaskCount++
	this.CurrentTask = ""
}

func (this *StopWatch) IsRunning() bool {
	return strings.IsNotBlank(this.CurrentTask)
}

func (this *StopWatch) shortSummary() string {
	return fmt.Sprintf("StopWatch '%s': %d total time in %d tasks", this.Id, this.TotalTime, this.TaskCount)
}
func (this *StopWatch) PrettyPrint() string {
	var sb strings2.Builder
	sb.WriteString(this.shortSummary())
	sb.WriteRune('\n')

	if !this.KeepTaskList {
		sb.WriteString("No task info kept")
	} else {
		sb.WriteString("---------------------------------------------\n")
		sb.WriteString("ns         %     Task name\n")
		sb.WriteString("---------------------------------------------\n")

		for _, task := range this.TaskList.GetData() {
			sb.WriteString(fmt.Sprintf("%9d  ", task.time))
			sb.WriteString(fmt.Sprintf("%.3f  ", float64(task.time)/float64(this.TotalTime)))
			sb.WriteString(task.Name)
			sb.WriteRune('\n')
		}
	}

	return sb.String()
}
