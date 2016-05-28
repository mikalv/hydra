package pkg

import (
	"time"
	"github.com/Sirupsen/logrus"
	"github.com/go-errors/errors"
)

func Retry(maxWait time.Duration, failAfter time.Duration, f func() error) (err error) {
	var lastStart time.Time
	err = errors.New("Did not connect.")
	loopWait := time.Millisecond * 100
	retryStart := time.Now()
	for retryStart.Add(failAfter).After(time.Now()) {
		lastStart = time.Now()
		if err = f(); err == nil {
			return nil
		}

		if lastStart.Add(maxWait * 2).Before(time.Now()) {
			retryStart = time.Now()
		}

		logrus.Error(err)
		logrus.Infof("Retrying in %f seconds...", loopWait.Seconds())
		time.Sleep(loopWait)
		loopWait = loopWait * loopWait
		if loopWait > maxWait {
			loopWait = maxWait
		}
	}
	return err
}
