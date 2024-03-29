package xhoteloffline

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderAlipayfaceCreateAPIResponse 信用住支付创建接口 API返回值
// taobao.xhotel.order.alipayface.create
//
// 用于创建一笔信用住支付，主要应用场景是线下信用住
type TaobaoXhotelOrderAlipayfaceCreateAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderAlipayfaceCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderAlipayfaceCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelOrderAlipayfaceCreateAPIResponseModel).Reset()
}

// TaobaoXhotelOrderAlipayfaceCreateAPIResponseModel is 信用住支付创建接口 成功返回结果
type TaobaoXhotelOrderAlipayfaceCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_alipayface_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 酒店订单号, 和入参中传入一致
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 用于当传入多个入住人时, 将阿里旅行最终用于结算的入住人信息回传
	MainIdNumber string `json:"main_id_number,omitempty" xml:"main_id_number,omitempty"`
	// 用于签约和扣款的买家淘宝账号
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 阿里旅行生成的备注信息,用于提示用户一些注意事宜. 请将该字段的信息打印到客人的入住单上. 如果为空代表没有阿里旅行方面的特殊备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 阿里旅行生成的订单id，该字段很重要，后续结账等操作都要使用tid
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 阿里旅行为该笔订单提供的最大杂费(不含房费)担保金额,单位为分. 注意该金额指客人除了房费以外可消费的金额上限
	GuaranteeAmout int64 `json:"guarantee_amout,omitempty" xml:"guarantee_amout,omitempty"`
	// 阿里旅行平台提供的优惠金额,单位为分
	AlitripDiscount int64 `json:"alitrip_discount,omitempty" xml:"alitrip_discount,omitempty"`
	// 商家自身提供给该订单的优惠金额，单位为分
	SellerDiscount int64 `json:"seller_discount,omitempty" xml:"seller_discount,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderAlipayfaceCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OutOrderId = ""
	m.MainIdNumber = ""
	m.BuyerNick = ""
	m.Remark = ""
	m.Tid = 0
	m.GuaranteeAmout = 0
	m.AlitripDiscount = 0
	m.SellerDiscount = 0
}

var poolTaobaoXhotelOrderAlipayfaceCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelOrderAlipayfaceCreateAPIResponse)
	},
}

// GetTaobaoXhotelOrderAlipayfaceCreateAPIResponse 从 sync.Pool 获取 TaobaoXhotelOrderAlipayfaceCreateAPIResponse
func GetTaobaoXhotelOrderAlipayfaceCreateAPIResponse() *TaobaoXhotelOrderAlipayfaceCreateAPIResponse {
	return poolTaobaoXhotelOrderAlipayfaceCreateAPIResponse.Get().(*TaobaoXhotelOrderAlipayfaceCreateAPIResponse)
}

// ReleaseTaobaoXhotelOrderAlipayfaceCreateAPIResponse 将 TaobaoXhotelOrderAlipayfaceCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelOrderAlipayfaceCreateAPIResponse(v *TaobaoXhotelOrderAlipayfaceCreateAPIResponse) {
	v.Reset()
	poolTaobaoXhotelOrderAlipayfaceCreateAPIResponse.Put(v)
}
