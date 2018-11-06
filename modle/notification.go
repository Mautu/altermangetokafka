package modle

import "time"

/*
{
  "version": "4",
  "groupKey": <string>,    // key identifying the group of alerts (e.g. to deduplicate)
  "status": "<resolved|firing>",
  "receiver": <string>,
  "groupLabels": <object>,
  "commonLabels": <object>,
  "commonAnnotations": <object>,
  "externalURL": <string>,  // backlink to the Alertmanager.
  "alerts": [
    {
      "status": "<resolved|firing>",
      "labels": <object>,
      "annotations": <object>,
      "startsAt": "<rfc3339>",
      "endsAt": "<rfc3339>",
      "generatorURL": <string> // identifies the entity that caused the alert
    },
    ...
  ]
}
*/
type Alert struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:annotations`
	StartsAt    time.Time         `json:"startsAt"`
	EndsAt      time.Time         `json:"endsAt"`
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

/*
{
    "__module__": "__main__",
    "alert_type": "",
    "uuid": "20181105151856",
    "service_name": "load",
    "__class__": "alert",
    "content": "CRITICAL - load average: 65.64, 61.13, 57.84",
    "state": "CRITICAL",
    "alert_time": 1541402336,
    "address": "192.168.111.91",
    "host": "192.168.111.91",
    "notify_type": "PROBLEM"
}
*/
type messgaekafka struct {
	//__module__   string `json:"__module__"`
	//__class__    string `json:"__class__"`
	uuid         string `json:"uuid"`
	service_name string `json:"service_name"`
	content      string `json:"content"`
	alert_type   string `json:"alert_type"`
	state        string `json:"state"`
	alert_time   string `json:"alert_time"` //unixtime
	address      string `json:"address"`
	host         string `json:"host"`
	notify_type  string `json:"notify_type"`
}
