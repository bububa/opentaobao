package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponse 经纪人接待状态变更 API返回值
// alibaba.alihouse.existinghome.sync.broker.is.receive.customer
//
// 经纪人接待状态变更
type AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponseModel).Reset()
}

// AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponseModel is 经纪人接待状态变更 成功返回结果
type AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_sync_broker_is_receive_customer_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 1
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// 1
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 1
	ReturnSuccess bool `json:"return_success,omitempty" xml:"return_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ReturnCode = ""
	m.ReturnMessage = ""
	m.Data = false
	m.ReturnSuccess = false
}

var poolAlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponse)
	},
}

// GetAlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponse 从 sync.Pool 获取 AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponse
func GetAlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponse() *AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponse {
	return poolAlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponse.Get().(*AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponse)
}

// ReleaseAlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponse 将 AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponse(v *AlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseExistinghomeSyncBrokerIsReceiveCustomerAPIResponse.Put(v)
}
