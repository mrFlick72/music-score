package main

import (
	application "github.com/mrflick72/cloud-native-golang-framework"
	"github.com/mrflick72/cloud-native-golang-framework/configuration"
	"sync"
)

func main() {
	initConfigurationManager()
	wg := &sync.WaitGroup{}
	wg.Add(1)

	app := application.NewApplicationServer()
	application.StartApplicationServer(wg, app)
	wg.Wait()
}

func initConfigurationManager() {
	configurationWg := &sync.WaitGroup{}
	configurationWg.Add(1)
	manager := configuration.GetConfigurationManagerInstance()

	go manager.Init(configurationWg)
	configurationWg.Wait()
}
