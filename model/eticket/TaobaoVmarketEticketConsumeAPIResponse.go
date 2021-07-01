package eticket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketConsumeAPIResponse
电子票券消费通知 API返回值
taobao.vmarket.eticket.consume

外部合作商家电子票券消费回调接口 */
type TaobaoVmarketEticketConsumeAPIResponse struct {
	model.CommonResponse
	TaobaoVmarketEticketConsumeAPIResponseModel
}

// TaobaoVmarketEticketConsumeAPIResponseModel is 电子票券消费通知 成功返回结果
type TaobaoVmarketEticketConsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"vmarket_eticket_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 0:失败，1:成功
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 宝贝标题
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// 整个订单剩余的可核销数量
	LeftNum int64 `json:"left_num,omitempty" xml:"left_num,omitempty"`
	// 返回码消费后，需要发送的短信的模版
	SmsTpl string `json:"sms_tpl,omitempty" xml:"sms_tpl,omitempty"`
	// 服务内容，用在凭证验证成功后pos机打印小票给消费者
	PrintTpl string `json:"print_tpl,omitempty" xml:"print_tpl,omitempty"`
	// 核销流水号,可以通过该流水号来撤销对应的核销操作
	ConsumeSecialNum string `json:"consume_secial_num,omitempty" xml:"consume_secial_num,omitempty"`
	// 该核销码在核销后剩余的可核销份数，如果传了new_code来重新生成码，那么这些可核销份数会累积到新的码上
	CodeLeftNum int64 `json:"code_left_num,omitempty" xml:"code_left_num,omitempty"`
}
