package ottpay

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttPayOrderDeleteorderAPIResponse 退订应用中心支付订单 API返回值
// youku.ott.pay.order.deleteorder
//
// 应用中心sdk连续包月退订接口
type YoukuOttPayOrderDeleteorderAPIResponse struct {
	model.CommonResponse
	YoukuOttPayOrderDeleteorderAPIResponseModel
}

// Reset 清空结构体
func (m *YoukuOttPayOrderDeleteorderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YoukuOttPayOrderDeleteorderAPIResponseModel).Reset()
}

// YoukuOttPayOrderDeleteorderAPIResponseModel is 退订应用中心支付订单 成功返回结果
type YoukuOttPayOrderDeleteorderAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_ott_pay_order_deleteorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TvOrderResultDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *YoukuOttPayOrderDeleteorderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolYoukuOttPayOrderDeleteorderAPIResponse = sync.Pool{
	New: func() any {
		return new(YoukuOttPayOrderDeleteorderAPIResponse)
	},
}

// GetYoukuOttPayOrderDeleteorderAPIResponse 从 sync.Pool 获取 YoukuOttPayOrderDeleteorderAPIResponse
func GetYoukuOttPayOrderDeleteorderAPIResponse() *YoukuOttPayOrderDeleteorderAPIResponse {
	return poolYoukuOttPayOrderDeleteorderAPIResponse.Get().(*YoukuOttPayOrderDeleteorderAPIResponse)
}

// ReleaseYoukuOttPayOrderDeleteorderAPIResponse 将 YoukuOttPayOrderDeleteorderAPIResponse 保存到 sync.Pool
func ReleaseYoukuOttPayOrderDeleteorderAPIResponse(v *YoukuOttPayOrderDeleteorderAPIResponse) {
	v.Reset()
	poolYoukuOttPayOrderDeleteorderAPIResponse.Put(v)
}
