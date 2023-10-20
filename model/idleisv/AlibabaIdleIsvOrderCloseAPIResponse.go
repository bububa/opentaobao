package idleisv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvOrderCloseAPIResponse 服务商闲鱼卖家主动关闭订单 API返回值
// alibaba.idle.isv.order.close
//
// 供外部服务商 isv 提供卖家主动关闭交易订单的功能
type AlibabaIdleIsvOrderCloseAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvOrderCloseAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleIsvOrderCloseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleIsvOrderCloseAPIResponseModel).Reset()
}

// AlibabaIdleIsvOrderCloseAPIResponseModel is 服务商闲鱼卖家主动关闭订单 成功返回结果
type AlibabaIdleIsvOrderCloseAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_order_close_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果
	Result *AlibabaIdleIsvOrderCloseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleIsvOrderCloseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleIsvOrderCloseAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvOrderCloseAPIResponse)
	},
}

// GetAlibabaIdleIsvOrderCloseAPIResponse 从 sync.Pool 获取 AlibabaIdleIsvOrderCloseAPIResponse
func GetAlibabaIdleIsvOrderCloseAPIResponse() *AlibabaIdleIsvOrderCloseAPIResponse {
	return poolAlibabaIdleIsvOrderCloseAPIResponse.Get().(*AlibabaIdleIsvOrderCloseAPIResponse)
}

// ReleaseAlibabaIdleIsvOrderCloseAPIResponse 将 AlibabaIdleIsvOrderCloseAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleIsvOrderCloseAPIResponse(v *AlibabaIdleIsvOrderCloseAPIResponse) {
	v.Reset()
	poolAlibabaIdleIsvOrderCloseAPIResponse.Put(v)
}
