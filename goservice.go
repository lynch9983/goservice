package goservice

import (
	"fmt"
	"os"

	"github.com/kardianos/service"
)

type HandlerFunc func() error

func Service(name string, displayName string, description string, runSvr HandlerFunc, closeSvr HandlerFunc) error {
	svcConfig := &service.Config{
		Name:        name,
		DisplayName: displayName,
		Description: description,
	}

	prg := &program{
		RunSvr:   runSvr,
		CloseSvr: closeSvr,
	}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		return err
	}

	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			return err
		} else {
			fmt.Println(os.Args[1] + " 成功")
			return nil
		}
	}

	return s.Run()
}

type program struct {
	RunSvr   HandlerFunc
	CloseSvr HandlerFunc
}

func (p *program) Start(s service.Service) error {
	go p.RunSvr()
	return nil
}

func (p *program) Stop(s service.Service) error {
	return p.CloseSvr()
}
