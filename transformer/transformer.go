package transformer

import (
	"net"
	"strconv"
	"strings"

	model "github.com/Mautu/altermangetokafka/modle"
)

func hostaddr(instance string, host string) string {
	addr := strings.Split(instance, ":")[0]
	hostip := strings.Split(host, ":")[0]
	if net.ParseIP(addr) != nil && net.ParseIP(hostip) == nil {
		return addr
	}
	return hostip
}

//AlertToMessagekafka altermanger messge transform to kafka
func AlertToMessagekafka(alert model.Alert) (messagekafka *model.Messgaekafka) {
	var status map[string]string
	status = map[string]string{"resolved": "OK", "firing": "CRITICAL"}
	messagekafka.Uuid = alert.StartsAt.Format("20060102150405") + strconv.FormatInt(alert.StartsAt.Unix(), 10)
	messagekafka.State = status[alert.Status]
	messagekafka.Alert_type = ""
	messagekafka.Service_name = alert.Labels["alertname"]
	messagekafka.Content = alert.Annotations["description"]
	messagekafka.Notify_type = "PROBLEM"
	messagekafka.Alert_time = strconv.FormatInt(alert.StartsAt.Unix(), 10)
	messagekafka.Host = hostaddr(alert.Labels["instance"], alert.Labels["host"])
	messagekafka.Address = hostaddr(alert.Labels["instance"], alert.Labels["host"])
	return
}
