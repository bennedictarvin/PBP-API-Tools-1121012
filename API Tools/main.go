package main

import (
	"context"
	"fmt"
	"time"

	"github.com/blizzy78/goroutine"
	"github.com/go-gomail/gomail"
	"github.com/jasonlvhit/gocron"
	"github.com/redis/go-redis/v9"
)

func task() {
	fmt.Println("I am running task.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func ExampleClient() *redis.Client {
	url := "redis://user:password@localhost:3306%2F0?protocol=3"
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}
	return redis.NewClient(opts)
}

func main() {
	// GoMail
	m := gomail.NewMessage()
	m.SetHeader("From", "alex@example.com")
	m.SetHeader("To", "bob@example.com", "cora@example.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.example.com", 587, "user", "123456")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	// GoCRON
	// Do jobs without params
	gocron.Every(1).Second().Do(task)
	gocron.Every(2).Seconds().Do(task)
	gocron.Every(1).Minute().Do(task)
	gocron.Every(2).Minutes().Do(task)
	gocron.Every(1).Hour().Do(task)
	gocron.Every(2).Hours().Do(task)
	gocron.Every(1).Day().Do(task)
	gocron.Every(2).Days().Do(task)
	gocron.Every(1).Week().Do(task)
	gocron.Every(2).Weeks().Do(task)

	// Do jobs with params
	gocron.Every(1).Second().Do(taskWithParams, 1, "hello")

	// Do jobs on specific weekday
	gocron.Every(1).Monday().Do(task)
	gocron.Every(1).Thursday().Do(task)

	// Do a job at a specific time - 'hour:min:sec' - seconds optional
	gocron.Every(1).Day().At("10:30").Do(task)
	gocron.Every(1).Monday().At("18:30").Do(task)
	gocron.Every(1).Tuesday().At("18:30:59").Do(task)

	// Begin job immediately upon start
	gocron.Every(1).Hour().From(gocron.NextTick()).Do(task)

	// Begin job at a specific date/time
	t := time.Date(2019, time.November, 10, 15, 0, 0, 0, time.Local)
	gocron.Every(1).Hour().From(&t).Do(task)

	// NextRun gets the next running time
	_, time := gocron.NextRun()
	fmt.Println(time)

	// Remove a specific job
	gocron.Remove(task)

	// Clear all scheduled jobs
	gocron.Clear()

	// Start all the pending jobs
	<-gocron.Start()

	// also, you can create a new scheduler
	// to run two schedulers concurrently
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(task)
	<-s.Start()

	// GoRoutine
	worker := func(_ context.Context) {
		time.Sleep(100 * time.Millisecond)
	}

	goroutines := goroutine.New()

	// Start a new goroutine
	goroutine.Go(context.Background(), worker)

	// Cancel all goroutines' contexts, and wait for them to finish
	_ = goroutines.CancelAll(context.Background(), true)
}

// GoRedis
var ctx = context.Background()

func exampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:3306",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: Key value
	// Key2 does not exist
}
