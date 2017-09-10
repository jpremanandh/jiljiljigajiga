package core

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	"../utils"
)

type serviceConfig struct {
	Environment string  `json:"environment"`
	ModelPath   string  `json:"model_path"`
	Dsps        []*Dsps `json:"dsps"`
}

var ServiceConfigSingleton *serviceConfig

func init() {
	environment := flag.String("environment", "development", "")
	modelPath := flag.String("modelPath", "", "Path for the models. Currently it is inside model_repo")
	if *environment == "" || *modelPath == "" {
		panic("Cannot start service as enough start parameters were not provided.")
	}
	ServiceConfigSingleton = &serviceConfig{Environment: *environment, ModelPath: *modelPath}
	ServiceConfigSingleton.Dsps = LoadDsps()
}

func LoadDsps() []*Dsps {
	var dspsConfig = []Dsps{}
	modelContent, err := ioutil.ReadFile(ServiceConfigSingleton.ModelPath)
	if err != nil {
		panic("Start failed as the model file could not be read")
	}
	var allModels = map[string]interface{}{}
	if err = json.Unmarshal(modelContent, &allModels); err != nil {
		panic("Loading all model contents failed.")
	}
	dspsAsInterface := utils.JSONValue(allModels, ServiceConfigSingleton.Environment)

	var dspsBytes []byte
	if dspsBytes, err = json.Marshal(dspsAsInterface); err != nil {
		panic("Failed to marshal environment dsps")
	}

	if err = json.Unmarshal(dspsBytes, &dspsConfig); err != nil {
		panic("Failed to unmarshal dsps config.")
	}
	var dspsConfigPointers = []*Dsps{}
	for _, dsp := range dspsConfig {
		if dsp.limiter, err = NewRateLimiter(dsp.QPS); err != nil {
			panic("Invalid QPS configuration provided.")
		}
		dspsConfigPointers = append(dspsConfigPointers, &dsp)
	}
	return dspsConfigPointers
}
