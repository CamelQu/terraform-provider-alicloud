package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// PreCheckCreateOrderForClone invokes the rds.PreCheckCreateOrderForClone API synchronously
// api document: https://help.aliyun.com/api/rds/precheckcreateorderforclone.html
func (client *Client) PreCheckCreateOrderForClone(request *PreCheckCreateOrderForCloneRequest) (response *PreCheckCreateOrderForCloneResponse, err error) {
	response = CreatePreCheckCreateOrderForCloneResponse()
	err = client.DoAction(request, response)
	return
}

// PreCheckCreateOrderForCloneWithChan invokes the rds.PreCheckCreateOrderForClone API asynchronously
// api document: https://help.aliyun.com/api/rds/precheckcreateorderforclone.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PreCheckCreateOrderForCloneWithChan(request *PreCheckCreateOrderForCloneRequest) (<-chan *PreCheckCreateOrderForCloneResponse, <-chan error) {
	responseChan := make(chan *PreCheckCreateOrderForCloneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PreCheckCreateOrderForClone(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// PreCheckCreateOrderForCloneWithCallback invokes the rds.PreCheckCreateOrderForClone API asynchronously
// api document: https://help.aliyun.com/api/rds/precheckcreateorderforclone.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) PreCheckCreateOrderForCloneWithCallback(request *PreCheckCreateOrderForCloneRequest, callback func(response *PreCheckCreateOrderForCloneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PreCheckCreateOrderForCloneResponse
		var err error
		defer close(result)
		response, err = client.PreCheckCreateOrderForClone(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// PreCheckCreateOrderForCloneRequest is the request struct for api PreCheckCreateOrderForClone
type PreCheckCreateOrderForCloneRequest struct {
	*requests.RpcRequest
	OwnerId                   requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount      string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId           requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken               string           `position:"Query" name:"ClientToken"`
	CommodityCode             string           `position:"Query" name:"CommodityCode"`
	DBInstanceId              string           `position:"Query" name:"DBInstanceId"`
	DBInstanceClass           string           `position:"Query" name:"DBInstanceClass"`
	DBInstanceStorage         requests.Integer `position:"Query" name:"DBInstanceStorage"`
	DBInstanceStorageType     string           `position:"Query" name:"DBInstanceStorageType"`
	PayType                   string           `position:"Query" name:"PayType"`
	InstanceNetworkType       string           `position:"Query" name:"InstanceNetworkType"`
	UsedTime                  string           `position:"Query" name:"UsedTime"`
	TimeType                  string           `position:"Query" name:"TimeType"`
	Quantity                  requests.Integer `position:"Query" name:"Quantity"`
	AutoPay                   requests.Boolean `position:"Query" name:"AutoPay"`
	InstanceUsedType          requests.Integer `position:"Query" name:"InstanceUsedType"`
	Resource                  string           `position:"Query" name:"Resource"`
	VPCId                     string           `position:"Query" name:"VPCId"`
	VSwitchId                 string           `position:"Query" name:"VSwitchId"`
	PrivateIpAddress          string           `position:"Query" name:"PrivateIpAddress"`
	CountryCode               string           `position:"Query" name:"CountryCode"`
	CurrencyCode              string           `position:"Query" name:"CurrencyCode"`
	AutoRenew                 string           `position:"Query" name:"AutoRenew"`
	AgentId                   string           `position:"Query" name:"AgentId"`
	ZoneId                    string           `position:"Query" name:"ZoneId"`
	PromotionCode             string           `position:"Query" name:"PromotionCode"`
	BusinessInfo              string           `position:"Query" name:"BusinessInfo"`
	CloneInstanceDefaultValue string           `position:"Query" name:"CloneInstanceDefaultValue"`
	DBInstanceDescription     string           `position:"Query" name:"DBInstanceDescription"`
	DBNames                   string           `position:"Query" name:"DBNames"`
	BackupId                  string           `position:"Query" name:"BackupId"`
	RestoreTime               string           `position:"Query" name:"RestoreTime"`
	OwnerAccount              string           `position:"Query" name:"OwnerAccount"`
	ResourceGroupId           string           `position:"Query" name:"ResourceGroupId"`
	TableMeta                 string           `position:"Query" name:"TableMeta"`
	RestoreTable              string           `position:"Query" name:"RestoreTable"`
	NodeType                  string           `position:"Query" name:"NodeType"`
}

// PreCheckCreateOrderForCloneResponse is the response struct for api PreCheckCreateOrderForClone
type PreCheckCreateOrderForCloneResponse struct {
	*responses.BaseResponse
	RequestId      string                                `json:"RequestId" xml:"RequestId"`
	PreCheckResult bool                                  `json:"PreCheckResult" xml:"PreCheckResult"`
	Failures       FailuresInPreCheckCreateOrderForClone `json:"Failures" xml:"Failures"`
}

// CreatePreCheckCreateOrderForCloneRequest creates a request to invoke PreCheckCreateOrderForClone API
func CreatePreCheckCreateOrderForCloneRequest() (request *PreCheckCreateOrderForCloneRequest) {
	request = &PreCheckCreateOrderForCloneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "PreCheckCreateOrderForClone", "rds", "openAPI")
	return
}

// CreatePreCheckCreateOrderForCloneResponse creates a response to parse from PreCheckCreateOrderForClone response
func CreatePreCheckCreateOrderForCloneResponse() (response *PreCheckCreateOrderForCloneResponse) {
	response = &PreCheckCreateOrderForCloneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
