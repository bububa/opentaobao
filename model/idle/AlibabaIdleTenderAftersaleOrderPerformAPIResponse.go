package idle

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleTenderAftersaleOrderPerformAPIResponse 闲鱼帮卖售后订单履约 API返回值
// alibaba.idle.tender.aftersale.order.perform
//
// 闲鱼帮卖售后订单履约
type AlibabaIdleTenderAftersaleOrderPerformAPIResponse struct {
	model.CommonResponse
	AlibabaIdleTenderAftersaleOrderPerformAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleTenderAftersaleOrderPerformAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleTenderAftersaleOrderPerformAPIResponseModel).Reset()
}

// AlibabaIdleTenderAftersaleOrderPerformAPIResponseModel is 闲鱼帮卖售后订单履约 成功返回结果
type AlibabaIdleTenderAftersaleOrderPerformAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_tender_aftersale_order_perform_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	PerformError string `json:"perform_error,omitempty" xml:"perform_error,omitempty"`
	// 是否履约成功
	TransformSuccess bool `json:"transform_success,omitempty" xml:"transform_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleTenderAftersaleOrderPerformAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PerformError = ""
	m.TransformSuccess = false
}

var poolAlibabaIdleTenderAftersaleOrderPerformAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleTenderAftersaleOrderPerformAPIResponse)
	},
}

// GetAlibabaIdleTenderAftersaleOrderPerformAPIResponse 从 sync.Pool 获取 AlibabaIdleTenderAftersaleOrderPerformAPIResponse
func GetAlibabaIdleTenderAftersaleOrderPerformAPIResponse() *AlibabaIdleTenderAftersaleOrderPerformAPIResponse {
	return poolAlibabaIdleTenderAftersaleOrderPerformAPIResponse.Get().(*AlibabaIdleTenderAftersaleOrderPerformAPIResponse)
}

// ReleaseAlibabaIdleTenderAftersaleOrderPerformAPIResponse 将 AlibabaIdleTenderAftersaleOrderPerformAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleTenderAftersaleOrderPerformAPIResponse(v *AlibabaIdleTenderAftersaleOrderPerformAPIResponse) {
	v.Reset()
	poolAlibabaIdleTenderAftersaleOrderPerformAPIResponse.Put(v)
}
