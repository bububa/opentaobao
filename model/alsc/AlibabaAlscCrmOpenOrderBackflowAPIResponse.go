package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmOpenOrderBackflowAPIResponse 订单回流接口 API返回值
// alibaba.alsc.crm.open.order.backflow
//
// 回流isv订单接口
type AlibabaAlscCrmOpenOrderBackflowAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmOpenOrderBackflowAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscCrmOpenOrderBackflowAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscCrmOpenOrderBackflowAPIResponseModel).Reset()
}

// AlibabaAlscCrmOpenOrderBackflowAPIResponseModel is 订单回流接口 成功返回结果
type AlibabaAlscCrmOpenOrderBackflowAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_open_order_backflow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscCrmOpenOrderBackflowAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscCrmOpenOrderBackflowAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscCrmOpenOrderBackflowAPIResponse)
	},
}

// GetAlibabaAlscCrmOpenOrderBackflowAPIResponse 从 sync.Pool 获取 AlibabaAlscCrmOpenOrderBackflowAPIResponse
func GetAlibabaAlscCrmOpenOrderBackflowAPIResponse() *AlibabaAlscCrmOpenOrderBackflowAPIResponse {
	return poolAlibabaAlscCrmOpenOrderBackflowAPIResponse.Get().(*AlibabaAlscCrmOpenOrderBackflowAPIResponse)
}

// ReleaseAlibabaAlscCrmOpenOrderBackflowAPIResponse 将 AlibabaAlscCrmOpenOrderBackflowAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscCrmOpenOrderBackflowAPIResponse(v *AlibabaAlscCrmOpenOrderBackflowAPIResponse) {
	v.Reset()
	poolAlibabaAlscCrmOpenOrderBackflowAPIResponse.Put(v)
}
