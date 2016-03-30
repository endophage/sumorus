# SumoLogic Hooks for [Logrus](https://github.com/Sirupsen/logrus) <img src="http://i.imgur.com/hTeVwmJ.png" width="40" height="40" alt=":walrus:" class="emoji" title=":walrus:"/>

## Usage

```go
package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/doublefree/sumorus"
)

var endpoint string = "YOUR_SUMOLOGIC_HTTP_HOSTED_ENDPOINT"
var host = "YOUR_HOST_NAME"

func main() {
	log := logrus.New()
	sumoLogicHook := sumorus.NewSumoLogicHook(endpoint, host, logrus.InfoLevel, "tag1", "tag2")
	log.Hooks.Add(sumoLogicHook)

	log.WithFields(logrus.Fields{
		"name": "joe",
		"age":  42,
	}).Error("Hello world!")
}
```