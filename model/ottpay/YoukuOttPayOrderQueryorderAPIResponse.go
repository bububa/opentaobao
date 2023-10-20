package ottpay

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttPayOrderQueryorderAPIResponse 查询订单 API返回值
// youku.ott.pay.order.queryorder
//
// 通过订单号查询订单信息
type YoukuOttPayOrderQueryorderAPIResponse struct {
	model.CommonResponse
	YoukuOttPayOrderQueryorderAPIResponseModel
}

// Reset 清空结构体
func (m *YoukuOttPayOrderQueryorderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YoukuOttPayOrderQueryorderAPIResponseModel).Reset()
}

// YoukuOttPayOrderQueryorderAPIResponseModel is 查询订单 成功返回结果
type YoukuOttPayOrderQueryorderAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_ott_pay_order_queryorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// status
	Data *TvOrderQueryResultDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *YoukuOttPayOrderQueryorderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolYoukuOttPayOrderQueryorderAPIResponse = sync.Pool{
	New: func() any {
		return new(YoukuOttPayOrderQueryorderAPIResponse)
	},
}

// GetYoukuOttPayOrderQueryorderAPIResponse 从 sync.Pool 获取 YoukuOttPayOrderQueryorderAPIResponse
func GetYoukuOttPayOrderQueryorderAPIResponse() *YoukuOttPayOrderQueryorderAPIResponse {
	return poolYoukuOttPayOrderQueryorderAPIResponse.Get().(*YoukuOttPayOrderQueryorderAPIResponse)
}

// ReleaseYoukuOttPayOrderQueryorderAPIResponse 将 YoukuOttPayOrderQueryorderAPIResponse 保存到 sync.Pool
func ReleaseYoukuOttPayOrderQueryorderAPIResponse(v *YoukuOttPayOrderQueryorderAPIResponse) {
	v.Reset()
	poolYoukuOttPayOrderQueryorderAPIResponse.Put(v)
}
