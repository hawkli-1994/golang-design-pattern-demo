package design

import "fmt"


type Seller interface {
	sell(name string)
}

type Station struct {
	stock int
}

func (station *Station) sell(name string) {
	if station.stock > 0 {
		station.stock--
		fmt.Printf("%s get paper, stock: %d\n", name, station.stock)
	} else {
		fmt.Println("over")
	}
}

type StationProxy struct {
	s *Station
}

func (proxy *StationProxy) sell(name string) {
	if proxy.s.stock > 0 {
		proxy.s.stock--
		fmt.Printf("%s, %d\n", name, proxy.s.stock)
	} else {
		fmt.Println("over")
	}
}


