package v1

import (
	"spiderman-center/common/routes/base"
)

type V1Router struct {
	base.BaseRouter
}

func NewV1Router() *V1Router {
	return (&V1Router{}).init()
}

func (r *V1Router) init() *V1Router {
	r.Register("/etcd", NewEtcdRouter())
	return r
}
