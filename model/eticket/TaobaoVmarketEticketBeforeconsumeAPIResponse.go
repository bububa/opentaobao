package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketBeforeconsumeAPIResponse 电子凭证验码前置确认 API返回值
// taobao.vmarket.eticket.beforeconsume
//
// 商家验码之前的调用接口，用来同步到最新的订单状态并判断是否可以进行验码操作
type TaobaoVmarketEticketBeforeconsumeAPIResponse struct {
	model.CommonResponse
	TaobaoVmarketEticketBeforeconsumeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketBeforeconsumeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoVmarketEticketBeforeconsumeAPIResponseModel).Reset()
}

// TaobaoVmarketEticketBeforeconsumeAPIResponseModel is 电子凭证验码前置确认 成功返回结果
type TaobaoVmarketEticketBeforeconsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"vmarket_eticket_beforeconsume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 有效期结束时间
	ValidEnds string `json:"valid_ends,omitempty" xml:"valid_ends,omitempty"`
	// 有效期开始时间
	ValidStart string `json:"valid_start,omitempty" xml:"valid_start,omitempty"`
	// 扩展字段，暂时预留为0，没有任何意义
	LeftAmount string `json:"left_amount,omitempty" xml:"left_amount,omitempty"`
	// 商品标题
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// 1:可以进行核销码操作
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 当前订单剩余可核销数量
	LeftNum int64 `json:"left_num,omitempty" xml:"left_num,omitempty"`
	// 订单ID
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 扩展字段，暂时预留为0，没有任何意义
	ItemType int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
	// 当前码剩余可核销数量
	CodeLeftNum int64 `json:"code_left_num,omitempty" xml:"code_left_num,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketBeforeconsumeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ValidEnds = ""
	m.ValidStart = ""
	m.LeftAmount = ""
	m.ItemTitle = ""
	m.RetCode = 0
	m.LeftNum = 0
	m.OrderId = 0
	m.ItemType = 0
	m.CodeLeftNum = 0
}

var poolTaobaoVmarketEticketBeforeconsumeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoVmarketEticketBeforeconsumeAPIResponse)
	},
}

// GetTaobaoVmarketEticketBeforeconsumeAPIResponse 从 sync.Pool 获取 TaobaoVmarketEticketBeforeconsumeAPIResponse
func GetTaobaoVmarketEticketBeforeconsumeAPIResponse() *TaobaoVmarketEticketBeforeconsumeAPIResponse {
	return poolTaobaoVmarketEticketBeforeconsumeAPIResponse.Get().(*TaobaoVmarketEticketBeforeconsumeAPIResponse)
}

// ReleaseTaobaoVmarketEticketBeforeconsumeAPIResponse 将 TaobaoVmarketEticketBeforeconsumeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoVmarketEticketBeforeconsumeAPIResponse(v *TaobaoVmarketEticketBeforeconsumeAPIResponse) {
	v.Reset()
	poolTaobaoVmarketEticketBeforeconsumeAPIResponse.Put(v)
}
