package servicefinder

import (
	"github.com/anton-107/learn-go-gateway/request"
	"github.com/anton-107/learn-go-gateway/response"
	"github.com/hashicorp/consul/api"
	"log"
	"strings"
	"time"
)

type ServiceFinder struct {
	serviceDiscoveryClient *api.Client
	knownServices          map[string]*api.AgentService
}

func (finder *ServiceFinder) Handle(req *request.Request, res *response.Response, next func()) {
	for _, service := range finder.knownServices {
		log.Printf("Original request: %s", req.GetPath())
		log.Printf("Compare to service name: %s", service.Service)
		if strings.HasPrefix(req.GetPath(), service.Service) {
			req.AddMatchedService(service)
		}
	}
	next()
}

func (finder *ServiceFinder) fetchServices() {
	services, err := finder.serviceDiscoveryClient.Agent().Services()
	if err != nil {
		log.Printf("error fetching services: %v", err)
	}

	finder.knownServices = services
}

func initTimer(r ServiceFinder) {
	ticker := time.NewTicker(3 * time.Second)

	for {
		select {
		case <-ticker.C:
			go r.fetchServices()
		}
	}
}

func NewServiceFinder() *ServiceFinder {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	r := ServiceFinder{
		serviceDiscoveryClient: client,
	}

	go initTimer(r)
	go r.fetchServices()

	return &r
}
