package logistic

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEleFengniaoOrderPushAPIResponse 推送订单 API返回值
// alibaba.ele.fengniao.order.push
//
// 推送淘宝订单至蜂鸟开放平台配送
type AlibabaEleFengniaoOrderPushAPIResponse struct {
	model.CommonResponse
	AlibabaEleFengniaoOrderPushAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEleFengniaoOrderPushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEleFengniaoOrderPushAPIResponseModel).Reset()
}

// AlibabaEleFengniaoOrderPushAPIResponseModel is 推送订单 成功返回结果
type AlibabaEleFengniaoOrderPushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_fengniao_order_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// msg
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEleFengniaoOrderPushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
}

var poolAlibabaEleFengniaoOrderPushAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEleFengniaoOrderPushAPIResponse)
	},
}

// GetAlibabaEleFengniaoOrderPushAPIResponse 从 sync.Pool 获取 AlibabaEleFengniaoOrderPushAPIResponse
func GetAlibabaEleFengniaoOrderPushAPIResponse() *AlibabaEleFengniaoOrderPushAPIResponse {
	return poolAlibabaEleFengniaoOrderPushAPIResponse.Get().(*AlibabaEleFengniaoOrderPushAPIResponse)
}

// ReleaseAlibabaEleFengniaoOrderPushAPIResponse 将 AlibabaEleFengniaoOrderPushAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEleFengniaoOrderPushAPIResponse(v *AlibabaEleFengniaoOrderPushAPIResponse) {
	v.Reset()
	poolAlibabaEleFengniaoOrderPushAPIResponse.Put(v)
}
