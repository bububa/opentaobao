package ottpay

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttPayOrderAuthpayAPIResponse 委托代扣服务 API返回值
// youku.ott.pay.order.authpay
//
// 应用中心sdk连续包月委托代扣服务
type YoukuOttPayOrderAuthpayAPIResponse struct {
	model.CommonResponse
	YoukuOttPayOrderAuthpayAPIResponseModel
}

// Reset 清空结构体
func (m *YoukuOttPayOrderAuthpayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YoukuOttPayOrderAuthpayAPIResponseModel).Reset()
}

// YoukuOttPayOrderAuthpayAPIResponseModel is 委托代扣服务 成功返回结果
type YoukuOttPayOrderAuthpayAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_ott_pay_order_authpay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TvOrderResultDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *YoukuOttPayOrderAuthpayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolYoukuOttPayOrderAuthpayAPIResponse = sync.Pool{
	New: func() any {
		return new(YoukuOttPayOrderAuthpayAPIResponse)
	},
}

// GetYoukuOttPayOrderAuthpayAPIResponse 从 sync.Pool 获取 YoukuOttPayOrderAuthpayAPIResponse
func GetYoukuOttPayOrderAuthpayAPIResponse() *YoukuOttPayOrderAuthpayAPIResponse {
	return poolYoukuOttPayOrderAuthpayAPIResponse.Get().(*YoukuOttPayOrderAuthpayAPIResponse)
}

// ReleaseYoukuOttPayOrderAuthpayAPIResponse 将 YoukuOttPayOrderAuthpayAPIResponse 保存到 sync.Pool
func ReleaseYoukuOttPayOrderAuthpayAPIResponse(v *YoukuOttPayOrderAuthpayAPIResponse) {
	v.Reset()
	poolYoukuOttPayOrderAuthpayAPIResponse.Put(v)
}
