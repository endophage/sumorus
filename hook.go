package sumorus

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

type SumoLogicHook struct {
	endPointUrl string
	tags        []string
	host        string
	levels      []logrus.Level
}

type SumoLogicMesssage struct {
	Tags  []string    `json:"tags"`
	Host  string      `json:"host"`
	Level string      `json:"level"`
	Data  interface{} `json:"data"`
}

func NewSumoLogicHook(endPointUrl string, host string, level logrus.Level, tags ...string) *SumoLogicHook {
	levels := []logrus.Level{}
	for _, l := range []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	} {
		if l <= level {
			levels = append(levels, l)
		}
	}

	var tagList []string
	for _, tag := range tags {
		tagList = append(tagList, tag)
	}

	return &SumoLogicHook{
		host:        host,
		tags:        tagList,
		endPointUrl: endPointUrl,
		levels:      levels,
	}
}

func (hook *SumoLogicHook) Fire(entry *logrus.Entry) error {
	dataStr, _ := entry.String()
	var data interface{}
	json.Unmarshal([]byte(dataStr), &data)
	message := SumoLogicMesssage{
		Tags:  hook.tags,
		Host:  hook.host,
		Level: strings.ToUpper(entry.Level.String()),
		Data:  data,
	}
	payload, _ := json.Marshal(message)
	req, err := http.NewRequest(
		"POST",
		hook.endPointUrl,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Duration(30 * time.Second)}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (hook *SumoLogicHook) Levels() []logrus.Level {
	return hook.levels
}
