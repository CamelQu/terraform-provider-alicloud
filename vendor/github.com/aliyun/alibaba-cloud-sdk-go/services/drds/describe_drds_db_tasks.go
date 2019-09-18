package drds

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

// DescribeDrdsDbTasks invokes the drds.DescribeDrdsDbTasks API synchronously
// api document: https://help.aliyun.com/api/drds/describedrdsdbtasks.html
func (client *Client) DescribeDrdsDbTasks(request *DescribeDrdsDbTasksRequest) (response *DescribeDrdsDbTasksResponse, err error) {
	response = CreateDescribeDrdsDbTasksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDrdsDbTasksWithChan invokes the drds.DescribeDrdsDbTasks API asynchronously
// api document: https://help.aliyun.com/api/drds/describedrdsdbtasks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDrdsDbTasksWithChan(request *DescribeDrdsDbTasksRequest) (<-chan *DescribeDrdsDbTasksResponse, <-chan error) {
	responseChan := make(chan *DescribeDrdsDbTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDrdsDbTasks(request)
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

// DescribeDrdsDbTasksWithCallback invokes the drds.DescribeDrdsDbTasks API asynchronously
// api document: https://help.aliyun.com/api/drds/describedrdsdbtasks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDrdsDbTasksWithCallback(request *DescribeDrdsDbTasksRequest, callback func(response *DescribeDrdsDbTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDrdsDbTasksResponse
		var err error
		defer close(result)
		response, err = client.DescribeDrdsDbTasks(request)
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

// DescribeDrdsDbTasksRequest is the request struct for api DescribeDrdsDbTasks
type DescribeDrdsDbTasksRequest struct {
	*requests.RpcRequest
	TaskType       string `position:"Query" name:"TaskType"`
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// DescribeDrdsDbTasksResponse is the response struct for api DescribeDrdsDbTasks
type DescribeDrdsDbTasksResponse struct {
	*responses.BaseResponse
	RequestId string                     `json:"RequestId" xml:"RequestId"`
	Success   bool                       `json:"Success" xml:"Success"`
	Tasks     TasksInDescribeDrdsDbTasks `json:"Tasks" xml:"Tasks"`
}

// CreateDescribeDrdsDbTasksRequest creates a request to invoke DescribeDrdsDbTasks API
func CreateDescribeDrdsDbTasksRequest() (request *DescribeDrdsDbTasksRequest) {
	request = &DescribeDrdsDbTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeDrdsDbTasks", "drds", "openAPI")
	return
}

// CreateDescribeDrdsDbTasksResponse creates a response to parse from DescribeDrdsDbTasks response
func CreateDescribeDrdsDbTasksResponse() (response *DescribeDrdsDbTasksResponse) {
	response = &DescribeDrdsDbTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}