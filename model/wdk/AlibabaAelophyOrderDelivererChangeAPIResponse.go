package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAelophyOrderDelivererChangeAPIResponse 配送员信息变更接口 API返回值
// alibaba.aelophy.order.deliverer.change
//
// 配送员信息变更接口
type AlibabaAelophyOrderDelivererChangeAPIResponse struct {
	model.CommonResponse
	AlibabaAelophyOrderDelivererChangeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAelophyOrderDelivererChangeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAelophyOrderDelivererChangeAPIResponseModel).Reset()
}

// AlibabaAelophyOrderDelivererChangeAPIResponseModel is 配送员信息变更接口 成功返回结果
type AlibabaAelophyOrderDelivererChangeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aelophy_order_deliverer_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 配送员信息变更响应
	ApiResult *TopBaseResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAelophyOrderDelivererChangeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaAelophyOrderDelivererChangeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAelophyOrderDelivererChangeAPIResponse)
	},
}

// GetAlibabaAelophyOrderDelivererChangeAPIResponse 从 sync.Pool 获取 AlibabaAelophyOrderDelivererChangeAPIResponse
func GetAlibabaAelophyOrderDelivererChangeAPIResponse() *AlibabaAelophyOrderDelivererChangeAPIResponse {
	return poolAlibabaAelophyOrderDelivererChangeAPIResponse.Get().(*AlibabaAelophyOrderDelivererChangeAPIResponse)
}

// ReleaseAlibabaAelophyOrderDelivererChangeAPIResponse 将 AlibabaAelophyOrderDelivererChangeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAelophyOrderDelivererChangeAPIResponse(v *AlibabaAelophyOrderDelivererChangeAPIResponse) {
	v.Reset()
	poolAlibabaAelophyOrderDelivererChangeAPIResponse.Put(v)
}
