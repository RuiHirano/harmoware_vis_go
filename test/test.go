package main

import (
	"log"
	"strconv"
	"sync"
	"time"

	hv "github.com/RuiHirano/harmoware_vis_go"
)

func createMockAgents(num int) []*hv.Agent{
	agents := []*hv.Agent{}
	for i := 0; i < num; i++ {
		agents = append(agents, &hv.Agent{
			ID: strconv.Itoa(i),
			Type: hv.AgentType_PERSON,
			Latitude: 35.888,
			Longitude: 135.444,
		})
	}
	return agents
}

func main() {
	wg := sync.WaitGroup{} // for syncing other goroutines
	wg.Add(1)

	agents := createMockAgents(100)
	hvg := hv.NewHarmowareVisGo()
	go hvg.RunServer("localhost:5000")
	
	time.Sleep(3*time.Second)
	log.Printf("send agents")
	hvg.SendAgents(agents)

	wg.Wait()

}