package ottpay

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttPayOrderQuerycporderAPIResponse 查询支付订单对应cp订单号 API返回值
// youku.ott.pay.order.querycporder
//
// 根据支付订单查询对应cp订单号
type YoukuOttPayOrderQuerycporderAPIResponse struct {
	model.CommonResponse
	YoukuOttPayOrderQuerycporderAPIResponseModel
}

// Reset 清空结构体
func (m *YoukuOttPayOrderQuerycporderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YoukuOttPayOrderQuerycporderAPIResponseModel).Reset()
}

// YoukuOttPayOrderQuerycporderAPIResponseModel is 查询支付订单对应cp订单号 成功返回结果
type YoukuOttPayOrderQuerycporderAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_ott_pay_order_querycporder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TvOrderResultDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *YoukuOttPayOrderQuerycporderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolYoukuOttPayOrderQuerycporderAPIResponse = sync.Pool{
	New: func() any {
		return new(YoukuOttPayOrderQuerycporderAPIResponse)
	},
}

// GetYoukuOttPayOrderQuerycporderAPIResponse 从 sync.Pool 获取 YoukuOttPayOrderQuerycporderAPIResponse
func GetYoukuOttPayOrderQuerycporderAPIResponse() *YoukuOttPayOrderQuerycporderAPIResponse {
	return poolYoukuOttPayOrderQuerycporderAPIResponse.Get().(*YoukuOttPayOrderQuerycporderAPIResponse)
}

// ReleaseYoukuOttPayOrderQuerycporderAPIResponse 将 YoukuOttPayOrderQuerycporderAPIResponse 保存到 sync.Pool
func ReleaseYoukuOttPayOrderQuerycporderAPIResponse(v *YoukuOttPayOrderQuerycporderAPIResponse) {
	v.Reset()
	poolYoukuOttPayOrderQuerycporderAPIResponse.Put(v)
}
