package main

import (
	"fmt"
	"time"

	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println(time.Now())

	fmt.Println("I am runnning task.")

	//fmt.Println(checkStatus())

	x := checkStatus()

	// v := make([]bookingData, 0, len(checkStatus()))

	for a := 0; a < len(x); a++ {

		//fmt.Printf("value of a: %d\n", a)
		fmt.Println("Second  ", x[a])
		EmailAction(x[a], x[a].CustomerEmail)

	}
	//fmt.Println(len(x))

	//fmt.Println("Second  ", x[1])
	//EmailAction()
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func main() {
	// Do jobs with params
	//gocron.Every(1).Second().Do(taskWithParams, 1, "hello")

	// Do jobs without params
	// gocron.Every(1).Second().Do(task)
	//gocron.Every(2).Seconds().Do(task)
	gocron.Every(1).Minute().Do(task)
	//gocron.Every(2).Minutes().Do(task)
	//gocron.Every(1).Hour().Do(task)
	// gocron.Every(2).Hours().Do(task)
	// gocron.Every(1).Day().Do(task)
	// gocron.Every(2).Days().Do(task)

	// // Do jobs on specific weekday
	// gocron.Every(1).Monday().Do(task)
	//gocron.Every(1).Thursday().Do(task)
	// // function At() take a string like 'hour:min'
	//gocron.Every(1).Day().At("13:16").Do(task)
	// gocron.Every(1).Monday().At("18:30").Do(task)

	// remove, clear and next_run
	// _, time := gocron.NextRun()
	// fmt.Println(time)

	// gocron.Remove(task)
	// gocron.Clear()

	// function Start start all the pending jobs
	<-gocron.Start()

	// also , you can create a your new scheduler,
	// to run two scheduler concurrently
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(task)
	<-s.Start()

}
