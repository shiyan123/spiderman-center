package service

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"context"
)

func MonitorNodes(s *Service) {
	rch := s.Client.Watch(context.Background(), s.Path, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case clientv3.EventTypePut:
				fmt.Printf("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			case clientv3.EventTypeDelete:
				fmt.Printf("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}
}

