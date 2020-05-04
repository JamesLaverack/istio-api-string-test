package main

import (
	"encoding/json"
	"fmt"
	istionet "github.com/improbable/istio-api/networking/v1alpha3"
)

func main() {
	fmt.Println("Test")

	gw := istionet.Gateway{
		Servers: []*istionet.Server{
			{
				Hosts: []string{
					"example.com",
				},
				Port: &istionet.Port{
					Number:   15443,
					Protocol: "TLS",
					Name:     "TLS",
				},
				Tls: &istionet.Server_TLSOptions{
					Mode: istionet.Server_TLSOptions_AUTO_PASSTHROUGH,
				},
			},
		},
		Selector: map[string]string{
			"istio": "ingressgateway",
		},
	}

	// When this Gateway gets marshalled it gets the field `servers.tls.mode: AUTO_PASSTHROUGH`
	gwBytes, _ := json.Marshal(gw)
	fmt.Println(string(gwBytes))

	se := istionet.ServiceEntry{
		Hosts: []string{
			"example.com",
		},
		Location: istionet.ServiceEntry_MESH_INTERNAL,
		Resolution: istionet.ServiceEntry_DNS,
		Endpoints: []*istionet.ServiceEntry_Endpoint{
			{
				Address: "192.168.1.1",
				Ports: map[string]uint32{
					"http1": 15443,
				},
			},
		},
	}

	// But when this ServiceEntry gets marshalled it gets the fields `location` and `resolution` as ints.
	seBytes, _ := json.Marshal(se)
	fmt.Println(string(seBytes))
}
