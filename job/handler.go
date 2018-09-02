package job

import (
	"fmt"

	cron "gopkg.in/robfig/cron.v2"
)

type JobsHandler struct {
	cron *cron.Cron
}

func NewHandler() JobsHandler {
	jh := JobsHandler{
		cron: cron.New(),
	}
	return jh
}

func (jh *JobsHandler) CreateJobsFromServices(store Store, services []Service) error {
	for _, service := range services {
		if err := service.CheckJob.IsValid(); err != nil {
			return fmt.Errorf("Error with service: %s : %s", service.Name, err.Error())
		}
		jh.cron.AddFunc(service.CheckJob.Cron, service.CheckJob.ExecuteAndStore(service.Name, store))
	}
	return nil
}

func (jh *JobsHandler) Start() {
	jh.cron.Start()
}

func (jh *JobsHandler) Stop() error {
	jh.cron.Stop()
	return nil
}
