package main

import (
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"context"
	"log"
	"net/http"
	dis "spiderman-center/discovery"
	"time"
)

func main() {

	m, err := dis.NewMaster([]string{
		"http://10.163.50.179:2379",
		"http://10.163.69.200:2379",
		"http://10.163.87.47:2379",
	}, "services/")

	if err != nil {
		log.Fatal(err)
	}
	go func() {
		http.HandleFunc("/", IndexHandler)
		http.ListenAndServe("0.0.0.0:8001", nil)
	}()

	resp, _ := m.Client.Get(context.Background(), m.Path, clientv3.WithPrefix())
	for _,v := range resp.Kvs {
		fmt.Println(string(v.Key))
		fmt.Println(string(v.Value))
	}

	for {
		for k, v := range m.Nodes {
			fmt.Printf("node:%s, ip=%s\n", k, v.Info.IP)
		}
		fmt.Printf("nodes num = %d\n", len(m.Nodes))
		time.Sleep(time.Second * 5)
	}

}

type IpList struct {
	IpList []string `json:"ipList"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var ips IpList
	ips.IpList = []string{"10.163.50.179:2379", "10.163.69.200:2379", "10.163.87.47:2379"}
	body, _ := json.Marshal(&ips)
	w.Write(body)
	return
}
