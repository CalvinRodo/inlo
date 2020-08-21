package timestamp

import (
	"log"
	"sync"
	"time"
)

var once sync.Once
var instance time.Time

const (
	layoutISO = "2006-01-02"
)

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

// IsoFormat the current time in Iso Format (The one canada uses)
// Todo look up what ISO Number Canada uses
func IsoFormat() string {
	return CurrentTime().Format(layoutISO)
}
