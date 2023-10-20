package ottpay

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttPayOrderCreateorderAPIResponse 创建订单 API返回值
// youku.ott.pay.order.createorder
//
// ottpay创建订单
type YoukuOttPayOrderCreateorderAPIResponse struct {
	model.CommonResponse
	YoukuOttPayOrderCreateorderAPIResponseModel
}

// Reset 清空结构体
func (m *YoukuOttPayOrderCreateorderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YoukuOttPayOrderCreateorderAPIResponseModel).Reset()
}

// YoukuOttPayOrderCreateorderAPIResponseModel is 创建订单 成功返回结果
type YoukuOttPayOrderCreateorderAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_ott_pay_order_createorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *TvOrderResultDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *YoukuOttPayOrderCreateorderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolYoukuOttPayOrderCreateorderAPIResponse = sync.Pool{
	New: func() any {
		return new(YoukuOttPayOrderCreateorderAPIResponse)
	},
}

// GetYoukuOttPayOrderCreateorderAPIResponse 从 sync.Pool 获取 YoukuOttPayOrderCreateorderAPIResponse
func GetYoukuOttPayOrderCreateorderAPIResponse() *YoukuOttPayOrderCreateorderAPIResponse {
	return poolYoukuOttPayOrderCreateorderAPIResponse.Get().(*YoukuOttPayOrderCreateorderAPIResponse)
}

// ReleaseYoukuOttPayOrderCreateorderAPIResponse 将 YoukuOttPayOrderCreateorderAPIResponse 保存到 sync.Pool
func ReleaseYoukuOttPayOrderCreateorderAPIResponse(v *YoukuOttPayOrderCreateorderAPIResponse) {
	v.Reset()
	poolYoukuOttPayOrderCreateorderAPIResponse.Put(v)
}
