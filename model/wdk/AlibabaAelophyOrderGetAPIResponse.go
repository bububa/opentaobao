package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAelophyOrderGetAPIResponse 翱象拉取订单接口 API返回值
// alibaba.aelophy.order.get
//
// 翱象拉取订单接口
type AlibabaAelophyOrderGetAPIResponse struct {
	model.CommonResponse
	AlibabaAelophyOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAelophyOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAelophyOrderGetAPIResponseModel).Reset()
}

// AlibabaAelophyOrderGetAPIResponseModel is 翱象拉取订单接口 成功返回结果
type AlibabaAelophyOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aelophy_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应对象
	ApiResult *TopBaseResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAelophyOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaAelophyOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAelophyOrderGetAPIResponse)
	},
}

// GetAlibabaAelophyOrderGetAPIResponse 从 sync.Pool 获取 AlibabaAelophyOrderGetAPIResponse
func GetAlibabaAelophyOrderGetAPIResponse() *AlibabaAelophyOrderGetAPIResponse {
	return poolAlibabaAelophyOrderGetAPIResponse.Get().(*AlibabaAelophyOrderGetAPIResponse)
}

// ReleaseAlibabaAelophyOrderGetAPIResponse 将 AlibabaAelophyOrderGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAelophyOrderGetAPIResponse(v *AlibabaAelophyOrderGetAPIResponse) {
	v.Reset()
	poolAlibabaAelophyOrderGetAPIResponse.Put(v)
}
