package helper

import (
	"fmt"
	"spiderman-center/app"
)

func GenServerIps() (ips []string) {
	for _, v := range app.GetApp().Config.EtcdServer.Urls {
		ip := fmt.Sprintf("http://%s:%d",
			v, app.GetApp().Config.EtcdServer.Port)
		ips = append(ips, ip)
	}
	return
}