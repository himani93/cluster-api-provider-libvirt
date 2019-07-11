package libvirt

import (
	"fmt"
	"log"

	libvirt "github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

func CreateDomain(name string, vcpu int, memory_in_gb uint, image_uri string, user_data_uri string) error {
	domainDef, err := defineDomain(name, vcpu, memory_in_gb, image_uri, user_data_uri)
	if err != nil {
		return fmt.Errorf("Error defining domain %v XML: %v", name, err)
	}

	log.Printf("Domain Definition XML: %v", domainDef)

	conn, err := libvirt.NewConnect("qemu+tcp://192.168.122.1:16509/system")
	if err != nil {
		log.Printf("Error creating a new Libvirt connection: %v", err)
		return fmt.Errorf("Error creating a new Libvirt connection: %v", err)
	}
	defer conn.Close()

	_, err = conn.DomainCreateXML(domainDef, 0)
	if err != nil {
		return fmt.Errorf("Error creating a new domain: %v", err)
	}

	return nil
}

func DomainExists(domainName string) (bool, error) {
	conn, err := libvirt.NewConnect("qemu+tcp://192.168.122.1:16509/system")
	if err != nil {
		log.Printf("Error creating a new Libvirt connection: %v", err)
		return false, fmt.Errorf("Error creating a new Libvirt connection: %v", err)
	}
	defer conn.Close()

	domains, _ := conn.ListAllDomains(1)

	for _, d := range domains {
		name, err := d.GetName()
		if err != nil {
			log.Printf("Error getting domain name: %v", err)
			return false, fmt.Errorf("Error getting domain name: %v", err)
		}
		if name == domainName {
			return true, nil
		}
	}

	return false, nil
}

func defineDomain(name string, vcpu int, memory_in_gb uint, image_uri string, user_data_uri string) (string, error) {
	domcfg := &libvirtxml.Domain{}

	bootOrder := uint(1)
	domcfg = &libvirtxml.Domain{
		Type: "kvm",
		Name: name,
		Memory: &libvirtxml.DomainMemory{
			Value: memory_in_gb,
			Unit:  "GB",
		},
		VCPU: &libvirtxml.DomainVCPU{
			Value: vcpu,
		},
		OS: &libvirtxml.DomainOS{
			Type: &libvirtxml.DomainOSType{
				Type: "hvm",
			},
		},
		Devices: &libvirtxml.DomainDeviceList{
			Emulator: "/usr/bin/kvm-spice",
			Disks: []libvirtxml.DomainDisk{
				libvirtxml.DomainDisk{
					Device: "disk",
					Source: &libvirtxml.DomainDiskSource{
						File: &libvirtxml.DomainDiskSourceFile{
							File: image_uri,
						},
					},
					Target: &libvirtxml.DomainDiskTarget{
						Dev: "vda",
						Bus: "virtio",
					},
					Boot: &libvirtxml.DomainDeviceBoot{
						Order: bootOrder,
					},
					Driver: &libvirtxml.DomainDiskDriver{
						Name: "qemu",
						Type: "qcow2",
					},
				},
				libvirtxml.DomainDisk{
					Device: "cdrom",
					Source: &libvirtxml.DomainDiskSource{
						File: &libvirtxml.DomainDiskSourceFile{
							File: user_data_uri,
						},
					},
					Target: &libvirtxml.DomainDiskTarget{
						Dev: "hda",
						Bus: "ide",
					},
				},
			},
			Interfaces: []libvirtxml.DomainInterface{
				libvirtxml.DomainInterface{
					Source: &libvirtxml.DomainInterfaceSource{
						Network: &libvirtxml.DomainInterfaceSourceNetwork{
							Network: "default",
						},
						Bridge: &libvirtxml.DomainInterfaceSourceBridge{
							Bridge: "virbr0",
						},
					},
					Model: &libvirtxml.DomainInterfaceModel{
						Type: "virtio",
					},
				},
			},
			Serials: []libvirtxml.DomainSerial{
				libvirtxml.DomainSerial{
					Protocol: &libvirtxml.DomainChardevProtocol{
						Type: "serial",
					},
					Target: &libvirtxml.DomainSerialTarget{
						Port: new(uint),
					},
				},
			},
			Consoles: []libvirtxml.DomainConsole{
				libvirtxml.DomainConsole{
					Target: &libvirtxml.DomainConsoleTarget{
						Type: "serial",
						Port: new(uint),
					},
				},
				libvirtxml.DomainConsole{
					Target: &libvirtxml.DomainConsoleTarget{
						Type: "virtio",
					},
				},
			},
			Graphics: []libvirtxml.DomainGraphic{
				libvirtxml.DomainGraphic{
					Spice: &libvirtxml.DomainGraphicSpice{
						AutoPort: "yes",
					},
				},
			},
		},
	}
	return domcfg.Marshal()
}
