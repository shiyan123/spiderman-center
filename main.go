package main

import (
	"log"
	"spiderman-center/app"

	"encoding/json"
	"net/http"
	"spiderman-center/service"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := app.GetApp().Prepare(); err != nil {
		panic(err)
	}
	s, err := service.DisCoveryService()
	if err != nil {
		panic(err)
	}
	go service.MonitorNodes(s)
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe("0.0.0.0:8001", nil)
}

//func main() {
//
//	m, err := dis.NewMaster([]string{
//		"http://10.163.50.179:2379",
//		"http://10.163.69.200:2379",
//		"http://10.163.87.47:2379",
//	}, "services/")
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	go func() {
//		http.HandleFunc("/", IndexHandler)
//		http.ListenAndServe("0.0.0.0:8001", nil)
//	}()
//
//	resp, _ := m.Client.Get(context.Background(), m.Path, clientv3.WithPrefix())
//	key := ""
//	value := ""
//	for _, v := range resp.Kvs {
//		fmt.Println(string(v.Key))
//		key = string(v.Key)
//		fmt.Println(string(v.Value))
//		value = string(v.Value)
//	}
//
//	time.Sleep(10 * time.Second)
//	fmt.Println("<><>")
//
//	var info dis.ServiceInfo
//	json.Unmarshal([]byte(value), &info)
//	task := &model.TaskInfo{
//		TaskId:   "task1",
//		TaskName: "任务1",
//	}
//	conf := &model.TaskConfig{
//		ProgramUpdateAt: 0,
//		ProgramName:     "temp_fmt",
//		Path:            "/Users/hortor/work/go/src/demos/temp_fmt",
//		CronStr:         "*/5 * * * * *",
//		Status:          model.TaskStatus_Stop,
//	}
//	task.Config = conf
//	taskMap := make(map[string]*model.TaskInfo, 0)
//	taskMap["task1"] = task
//	info.TaskMap = taskMap
//
//	body, _ := json.Marshal(&info)
//
//	fmt.Println(key)
//	fmt.Println(string(body))
//	re, err := m.Client.Grant(context.TODO(), 5)
//	if err != nil {
//		return
//	}
//	fmt.Println(key)
//	fmt.Println(string(body))
//	m.Client.Put(context.Background(), key, string(body), clientv3.WithLease(re.ID))
//
//	for {
//		for k, v := range m.Nodes {
//			fmt.Printf("node:%s, ip=%s\n", k, v.Info.IP)
//		}
//		fmt.Printf("nodes num = %d\n", len(m.Nodes))
//		time.Sleep(time.Second * 5)
//	}
//
//}
//
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
