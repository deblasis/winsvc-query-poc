package main

import (
	"log"

	"golang.org/x/sys/windows/svc/mgr"
)

func main() {
	const svcName = "Dhcp"

	m, err := mgr.Connect()
	if err != nil {
		log.Fatalf("failed to connect to Windows service manager: %v", err)
	}
	defer m.Disconnect()

	s, err := m.OpenService(svcName)
	if err != nil {
		log.Fatalf("could not access service %v: %v", svcName, err)
	}
	defer s.Close()

	status, err := s.Query()
	if err != nil {
		log.Fatalf("could not query service %v: %v", svcName, err)
	}
	log.Printf("service %v has status %v", svcName, status)

	return
}
