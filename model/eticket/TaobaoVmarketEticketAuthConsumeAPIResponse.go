package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketAuthConsumeAPIResponse 核销放行的核销接口 API返回值
// taobao.vmarket.eticket.auth.consume
//
// 针对O2O电子凭证核销放行业务，为满足码商能够核销淘宝码而开放的核销接口
type TaobaoVmarketEticketAuthConsumeAPIResponse struct {
	model.CommonResponse
	TaobaoVmarketEticketAuthConsumeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketAuthConsumeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoVmarketEticketAuthConsumeAPIResponseModel).Reset()
}

// TaobaoVmarketEticketAuthConsumeAPIResponseModel is 核销放行的核销接口 成功返回结果
type TaobaoVmarketEticketAuthConsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"vmarket_eticket_auth_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品标题
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// 淘宝卖家旺旺名称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 1:可以进行核销码操作
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 淘宝卖家ID
	TaobaoSid int64 `json:"taobao_sid,omitempty" xml:"taobao_sid,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketAuthConsumeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ItemTitle = ""
	m.SellerNick = ""
	m.RetCode = 0
	m.OrderId = 0
	m.TaobaoSid = 0
}

var poolTaobaoVmarketEticketAuthConsumeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoVmarketEticketAuthConsumeAPIResponse)
	},
}

// GetTaobaoVmarketEticketAuthConsumeAPIResponse 从 sync.Pool 获取 TaobaoVmarketEticketAuthConsumeAPIResponse
func GetTaobaoVmarketEticketAuthConsumeAPIResponse() *TaobaoVmarketEticketAuthConsumeAPIResponse {
	return poolTaobaoVmarketEticketAuthConsumeAPIResponse.Get().(*TaobaoVmarketEticketAuthConsumeAPIResponse)
}

// ReleaseTaobaoVmarketEticketAuthConsumeAPIResponse 将 TaobaoVmarketEticketAuthConsumeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoVmarketEticketAuthConsumeAPIResponse(v *TaobaoVmarketEticketAuthConsumeAPIResponse) {
	v.Reset()
	poolTaobaoVmarketEticketAuthConsumeAPIResponse.Put(v)
}
