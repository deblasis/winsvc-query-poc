package main

import (
	"log"
	"syscall"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc/mgr"
)

func main() {
	const svcName = "Dhcp"

	h, err := OpenSCManagerLeastPrivilege()
	if err != nil {
		log.Fatalf("failed to connect to Windows service manager: %v", err)
	}
	m := &mgr.Mgr{Handle: h}
	defer m.Disconnect()

	service, err := OpenServiceLeastPrivilege(m, svcName)
	if err != nil {
		log.Fatalf("could not access service %v: %v", svcName, err)
	}
	defer service.Close()

	status, err := service.Query()
	if err != nil {
		log.Fatalf("could not query service %v: %v", svcName, err)
	}
	log.Printf("service %v has status %v", svcName, status)

	return
}

func OpenSCManagerLeastPrivilege() (windows.Handle, error) {
	var s *uint16
	return windows.OpenSCManager(s, nil, windows.SC_MANAGER_CONNECT)
}

func OpenServiceLeastPrivilege(m *mgr.Mgr, name string) (*mgr.Service, error) {
	h, err := windows.OpenService(m.Handle, syscall.StringToUTF16Ptr(name), windows.SC_MANAGER_ENUMERATE_SERVICE)
	if err != nil {
		return nil, err
	}
	return &mgr.Service{Name: name, Handle: h}, nil
}
