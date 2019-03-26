package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"spiderman-center/app"
	"spiderman-center/common/model"
	"sync"
	"time"
)

type Service struct {
	Path         string
	RunningGroup map[string][]*model.Node
	Client       *clientv3.Client
	ClientTTL    int64
}

var (
	serviceOnce sync.Once
	service     *Service
)

func GetService() *Service {
	serviceOnce.Do(func() {
		service = disCoveryService()
		if service == nil {
			panic(errors.New("service is error"))
		}
	})

	return service
}

func disCoveryService() *Service {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   genServerIps(),
		DialTimeout: 2 * time.Second,
	})

	if err != nil {
		return nil
	}

	s := &Service{
		Path:         app.GetApp().Config.EtcdServer.Path,
		RunningGroup: make(map[string][]*model.Node, 0),
		Client:       cli,
		ClientTTL:    app.GetApp().Config.Server.ClientTTL,
	}

	err = s.getRunningGroup()
	if err != nil {
		return nil
	}
	return s
}

// loading running group
func (s *Service) getRunningGroup() (err error) {
	resp, err := s.Client.Get(context.Background(), s.Path, clientv3.WithPrefix())
	if err != nil {
		return
	}

	for _, v := range resp.Kvs {
		node := s.node(string(v.Key), string(v.Value))
		if node != nil {
			s.RunningGroup[node.Name] = append(s.RunningGroup[node.Name], node)
		}
	}
	return
}

func (s *Service) node(key, value string) *model.Node {
	fmt.Printf("key-->: %s \n", key)
	var n model.Node
	err := json.Unmarshal([]byte(value), &n)
	if err != nil {
		return nil
	}
	return &n
}

func (s *Service) TestService() string {
	return "asdasd"
}

func genServerIps() (ips []string) {
	for _, v := range app.GetApp().Config.EtcdServer.Urls {
		ip := fmt.Sprintf("http://%s:%d",
			v, app.GetApp().Config.EtcdServer.Port)
		ips = append(ips, ip)
	}
	return
}
