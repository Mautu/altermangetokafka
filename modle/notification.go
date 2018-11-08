package modle

import "time"

type Alert struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:annotations`
	StartsAt    time.Time         `json:"startsAt"`
	EndsAt      time.Time         `json:"endsAt"`
	Status      string            `json:"status"`
}

type Notification struct {
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	Status            string            `json:"status"`
	Receiver          string            `json:receiver`
	GroupLabels       map[string]string `json:groupLabels`
	CommonLabels      map[string]string `json:commonLabels`
	CommonAnnotations map[string]string `json:commonAnnotations`
	ExternalURL       string            `json:externalURL`
	Alerts            []Alert           `json:alerts`
}

type Messgaekafka struct {
	//__module__   string `json:"__module__"`
	//__class__    string `json:"__class__"`
	Uuid         string `json:"uuid"`
	Service_name string `json:"service_name"`
	Content      string `json:"content"`
	Alert_type   string `json:"alert_type"`
	State        string `json:"state"`
	Alert_time   string `json:"alert_time"` //unixtime
	Address      string `json:"address"`
	Host         string `json:"host"`
	Notify_type  string `json:"notify_type"`
}
