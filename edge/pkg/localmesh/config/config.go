package config

import (
	"sync"

	"github.com/kubeedge/kubeedge/pkg/apis/componentconfig/edgecore/v1alpha1"
)

var config Configure
var once sync.Once

type Configure struct {
	v1alpha1.LocalMesh
	NodeName string
}

func InitConfigure(localMesh *v1alpha1.LocalMesh, nodeName string) {
	once.Do(func() {
		config = Configure{
			LocalMesh: *localMesh,
			NodeName:   nodeName,
		}
	})
}

func Get() *Configure {
	return &config
}
