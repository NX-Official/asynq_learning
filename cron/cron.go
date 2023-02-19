package main

import (
	"github.com/hibiken/asynq"
	"log"
)

const redisAddr = "127.0.0.1:6379"

func main() {
	scheduler := asynq.NewScheduler(asynq.RedisClientOpt{Addr: redisAddr}, nil)

	task := asynq.NewTask("example_task", nil)

	// You can use cron spec string to specify the schedule.
	entryID, err := scheduler.Register("* * * * *", task)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("registered an entry: %q\n", entryID)

	// You can use "@every <duration>" to specify the interval.
	entryID, err = scheduler.Register("@every 30s", task)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("registered an entry: %q\n", entryID)

	// You can also pass options.
	entryID, err = scheduler.Register("@every 24h", task, asynq.Queue("myqueue"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("registered an entry: %q\n", entryID)

	if err := scheduler.Run(); err != nil {
		log.Fatal(err)
	}
}
