package ottpay

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttPayOrderAuthpaywithpriceAPIResponse 委托代扣可配定价服务 API返回值
// youku.ott.pay.order.authpaywithprice
//
// 应用中心sdk连续包月委托代扣服务，次月可配置营销价
type YoukuOttPayOrderAuthpaywithpriceAPIResponse struct {
	model.CommonResponse
	YoukuOttPayOrderAuthpaywithpriceAPIResponseModel
}

// Reset 清空结构体
func (m *YoukuOttPayOrderAuthpaywithpriceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YoukuOttPayOrderAuthpaywithpriceAPIResponseModel).Reset()
}

// YoukuOttPayOrderAuthpaywithpriceAPIResponseModel is 委托代扣可配定价服务 成功返回结果
type YoukuOttPayOrderAuthpaywithpriceAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_ott_pay_order_authpaywithprice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Data *TvOrderResultDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *YoukuOttPayOrderAuthpaywithpriceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolYoukuOttPayOrderAuthpaywithpriceAPIResponse = sync.Pool{
	New: func() any {
		return new(YoukuOttPayOrderAuthpaywithpriceAPIResponse)
	},
}

// GetYoukuOttPayOrderAuthpaywithpriceAPIResponse 从 sync.Pool 获取 YoukuOttPayOrderAuthpaywithpriceAPIResponse
func GetYoukuOttPayOrderAuthpaywithpriceAPIResponse() *YoukuOttPayOrderAuthpaywithpriceAPIResponse {
	return poolYoukuOttPayOrderAuthpaywithpriceAPIResponse.Get().(*YoukuOttPayOrderAuthpaywithpriceAPIResponse)
}

// ReleaseYoukuOttPayOrderAuthpaywithpriceAPIResponse 将 YoukuOttPayOrderAuthpaywithpriceAPIResponse 保存到 sync.Pool
func ReleaseYoukuOttPayOrderAuthpaywithpriceAPIResponse(v *YoukuOttPayOrderAuthpaywithpriceAPIResponse) {
	v.Reset()
	poolYoukuOttPayOrderAuthpaywithpriceAPIResponse.Put(v)
}
