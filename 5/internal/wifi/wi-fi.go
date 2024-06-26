package wifi

import (
	"net"

	"github.com/mdlayher/wifi"
)

type WiFi interface {
	Interfaces() ([]*wifi.Interface, error)
}

type WiFiService struct {
	WiFi WiFi
}

func New(wifi WiFi) WiFiService {
	return WiFiService{WiFi: wifi}
}

func (service WiFiService) GetAddresses() ([]net.HardwareAddr, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, err
	}

	var addrs []net.HardwareAddr = make([]net.HardwareAddr, 0, len(interfaces))

	for _, iface := range interfaces {
		addrs = append(addrs, iface.HardwareAddr)
	}

	return addrs, nil
}

func (service WiFiService) GetNames() ([]string, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, err
	}

	var name_list []string = make([]string, 0, len(interfaces))

	for _, iface := range interfaces {
		name_list = append(name_list, iface.Name)
	}

	return name_list, nil
}
