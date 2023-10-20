package tvpay

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTvpayPartnerOrderQueryAPIResponse 商户查询订单 API返回值
// taobao.tvpay.partner.order.query
//
// 给商户提供的查询订单状态的API
type TaobaoTvpayPartnerOrderQueryAPIResponse struct {
	model.CommonResponse
	TaobaoTvpayPartnerOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTvpayPartnerOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTvpayPartnerOrderQueryAPIResponseModel).Reset()
}

// TaobaoTvpayPartnerOrderQueryAPIResponseModel is 商户查询订单 成功返回结果
type TaobaoTvpayPartnerOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tvpay_partner_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Top返回对象
	Result *TopResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTvpayPartnerOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTvpayPartnerOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTvpayPartnerOrderQueryAPIResponse)
	},
}

// GetTaobaoTvpayPartnerOrderQueryAPIResponse 从 sync.Pool 获取 TaobaoTvpayPartnerOrderQueryAPIResponse
func GetTaobaoTvpayPartnerOrderQueryAPIResponse() *TaobaoTvpayPartnerOrderQueryAPIResponse {
	return poolTaobaoTvpayPartnerOrderQueryAPIResponse.Get().(*TaobaoTvpayPartnerOrderQueryAPIResponse)
}

// ReleaseTaobaoTvpayPartnerOrderQueryAPIResponse 将 TaobaoTvpayPartnerOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTvpayPartnerOrderQueryAPIResponse(v *TaobaoTvpayPartnerOrderQueryAPIResponse) {
	v.Reset()
	poolTaobaoTvpayPartnerOrderQueryAPIResponse.Put(v)
}
