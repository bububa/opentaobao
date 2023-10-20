package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAelophyOrderWorkCallbackAPIResponse 仓配作业结果回传接口 API返回值
// alibaba.aelophy.order.work.callback
//
// 仓配作业结果回传接口
type AlibabaAelophyOrderWorkCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaAelophyOrderWorkCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAelophyOrderWorkCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAelophyOrderWorkCallbackAPIResponseModel).Reset()
}

// AlibabaAelophyOrderWorkCallbackAPIResponseModel is 仓配作业结果回传接口 成功返回结果
type AlibabaAelophyOrderWorkCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aelophy_order_work_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回传响应
	ApiResult *TopBaseResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAelophyOrderWorkCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaAelophyOrderWorkCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAelophyOrderWorkCallbackAPIResponse)
	},
}

// GetAlibabaAelophyOrderWorkCallbackAPIResponse 从 sync.Pool 获取 AlibabaAelophyOrderWorkCallbackAPIResponse
func GetAlibabaAelophyOrderWorkCallbackAPIResponse() *AlibabaAelophyOrderWorkCallbackAPIResponse {
	return poolAlibabaAelophyOrderWorkCallbackAPIResponse.Get().(*AlibabaAelophyOrderWorkCallbackAPIResponse)
}

// ReleaseAlibabaAelophyOrderWorkCallbackAPIResponse 将 AlibabaAelophyOrderWorkCallbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAelophyOrderWorkCallbackAPIResponse(v *AlibabaAelophyOrderWorkCallbackAPIResponse) {
	v.Reset()
	poolAlibabaAelophyOrderWorkCallbackAPIResponse.Put(v)
}
