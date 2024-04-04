package wifi

import (
	"net"
	"reflect"
	"testing"

	"github.com/mdlayher/wifi"
)

type MockWiFi struct{}

func (m MockWiFi) Interfaces() ([]*wifi.Interface, error) {
	return []*wifi.Interface{
		{
			HardwareAddr: net.HardwareAddr{0xDE, 0xAD, 0xBE, 0xEF, 0xFE, 0xED},
			Name:         "mockInterface",
		},
	}, nil
}

func TestWiFiService_GetAddresses(t *testing.T) {
	mockWiFi := MockWiFi{}
	wifiService := New(mockWiFi)

	addrs, err := wifiService.GetAddresses()

	expected := []net.HardwareAddr{{0xDE, 0xAD, 0xBE, 0xEF, 0xFE, 0xED}}
	if !reflect.DeepEqual(addrs, expected) {
		t.Errorf("Expected addresses to be %v, but got %v", expected, addrs)
	}

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
