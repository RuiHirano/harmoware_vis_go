package harmwoware_vis_go

import (
	"fmt"
)

type HarmowareVisGo struct{

}

func NewHarmowareVisGo() *HarmowareVisGo {
    return &HarmowareVisGo{

	}
}

func (hv *HarmowareVisGo) RunServer(backgraund bool) {
	if backgraund{
		fmt.Printf("background")
	}
	fmt.Printf("foreground")
}