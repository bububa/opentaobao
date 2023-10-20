package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscPurchaseServicedefinitionParamQueryAPIResponse 查询采购服务定义参数信息 API返回值
// alibaba.ssc.purchase.servicedefinition.param.query
//
// 查询采购服务定义参数信息
type AlibabaSscPurchaseServicedefinitionParamQueryAPIResponse struct {
	model.CommonResponse
	AlibabaSscPurchaseServicedefinitionParamQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSscPurchaseServicedefinitionParamQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSscPurchaseServicedefinitionParamQueryAPIResponseModel).Reset()
}

// AlibabaSscPurchaseServicedefinitionParamQueryAPIResponseModel is 查询采购服务定义参数信息 成功返回结果
type AlibabaSscPurchaseServicedefinitionParamQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_purchase_servicedefinition_param_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务标准定义list
	ServiceDefinitions []ServiceDefinition `json:"service_definitions,omitempty" xml:"service_definitions>service_definition,omitempty"`
	// 错误描述信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSscPurchaseServicedefinitionParamQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ServiceDefinitions = m.ServiceDefinitions[:0]
	m.DisplayMsg = ""
	m.IsSuccess = false
}

var poolAlibabaSscPurchaseServicedefinitionParamQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSscPurchaseServicedefinitionParamQueryAPIResponse)
	},
}

// GetAlibabaSscPurchaseServicedefinitionParamQueryAPIResponse 从 sync.Pool 获取 AlibabaSscPurchaseServicedefinitionParamQueryAPIResponse
func GetAlibabaSscPurchaseServicedefinitionParamQueryAPIResponse() *AlibabaSscPurchaseServicedefinitionParamQueryAPIResponse {
	return poolAlibabaSscPurchaseServicedefinitionParamQueryAPIResponse.Get().(*AlibabaSscPurchaseServicedefinitionParamQueryAPIResponse)
}

// ReleaseAlibabaSscPurchaseServicedefinitionParamQueryAPIResponse 将 AlibabaSscPurchaseServicedefinitionParamQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSscPurchaseServicedefinitionParamQueryAPIResponse(v *AlibabaSscPurchaseServicedefinitionParamQueryAPIResponse) {
	v.Reset()
	poolAlibabaSscPurchaseServicedefinitionParamQueryAPIResponse.Put(v)
}
