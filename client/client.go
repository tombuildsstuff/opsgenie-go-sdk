// Copyright 2015 OpsGenie. All rights reserved.
// Use of this source code is governed by a Apache Software 
// license that can be found in the LICENSE file.

/*
Package client manages the creation of API clients. 
API user first creates a pointer of type OpsGenieClient. Following that
he/she can set some configurations for HTTP communication layer by setting 
a proxy definition and/or transport layer options. 

Introduction

The most fundamental and general use case is being able to access the 
OpsGenie Web API by coding a Go program.
The program -by mean of a client application- can send OpsGenie Web API 
the requests using the 'client' package in a higher level. For the programmer 
of the client application, that reduces the number of LoCs.
Besides it will result a less error-prone application and reduce 
the complexity by hiding the low-level networking, error-handling and 
byte-processing calls.

Package client has ports for all entry points to the Web API. 
The OpsGenie Web API is structured in JSON-bodied 
calls (except the file attachment).
*/
package client

import "errors"

// OpsGenie Go SDK performs HTTP calls to the Web API.
// The Web API is designated by a URL so called an endpoint
const ENDPOINT_URL string = "https://api.opsgenie.com" 
// OpsGenieClient is a general data type used for:
// 	- authenticating callers through their api keys and 
// 	- instanciating "alert" and "heartbeat" clients
//	- setting HTTP transport layer configurations
type OpsGenieClient struct {
	proxy *ClientProxyConfiguration
	httpTransportSettings *HttpTransportSettings
	apiKey string
}
// Setters:
//	- proxy
//	- http transport layer conf
//	- api key
func (cli *OpsGenieClient) SetClientProxyConfiguration(conf *ClientProxyConfiguration) {
	cli.proxy = conf
}

func (cli *OpsGenieClient) SetHttpTransportSettings(settings *HttpTransportSettings) {
	cli.httpTransportSettings = settings
}

func (cli *OpsGenieClient) SetApiKey(key string) error {
	if key == "" {
		return errors.New("API Key can not be empty")
	}
	cli.apiKey = key
	return nil
}

// Instanciates a new OpsGenieAlertClient
// and sets the api key to be used alongside the execution.
func (cli *OpsGenieClient) Alert() (*OpsGenieAlertClient, error) {
	if cli.apiKey == "" {
		return nil, errors.New("API Key should be set first")
	}
	alertClient := new (OpsGenieAlertClient)
	alertClient.apiKey = cli.apiKey
	if cli.proxy != nil {
		alertClient.proxy = cli.proxy.ToString()	
	}
	return alertClient, nil
}
// Instanciates a new OpsGenieHeartbeatClient
// and sets the api key to be used alongside the execution.
func (cli *OpsGenieClient) Heartbeat() (*OpsGenieHeartbeatClient, error) {
	if cli.apiKey == "" {
		return nil, errors.New("API Key should be set first")
	}
	heartbeatClient := new (OpsGenieHeartbeatClient)
	heartbeatClient.apiKey = cli.apiKey
	if cli.proxy != nil {
		heartbeatClient.proxy = cli.proxy.ToString()	
	}
	return heartbeatClient, nil
}
// Instanciates a new OpsGenieIntegrationClient
// and sets the api key to be used alongside the execution.
func (cli *OpsGenieClient) Integration() (*OpsGenieIntegrationClient, error) {
	if cli.apiKey == "" {
		return nil, errors.New("API Key should be set first")
	}
	integrationClient := new (OpsGenieIntegrationClient)
	integrationClient.apiKey = cli.apiKey
	if cli.proxy != nil {
		integrationClient.proxy = cli.proxy.ToString()	
	}
	return integrationClient, nil
}

// Instanciates a new OpsGeniePolicyClient
// and sets the api key to be used alongside the execution.
func (cli *OpsGenieClient) Policy() (*OpsGeniePolicyClient, error) {
	if cli.apiKey == "" {
		return nil, errors.New("API Key should be set first")
	}
	policyClient := new (OpsGeniePolicyClient)
	policyClient.apiKey = cli.apiKey
	if cli.proxy != nil {
		policyClient.proxy = cli.proxy.ToString()	
	}
	return policyClient, nil
}
