package timestamp

import (
	"log"
	"sync"
	"time"
)

var once sync.Once
var instance time.Time

// CurrentTime : get the current time
func CurrentTime() time.Time {

	once.Do(func() {
		location, err := time.LoadLocation("America/Toronto")
		if err != nil {
			log.Fatal(err)
		}
		instance = time.Now().In(location)
	})

	return instance
}
