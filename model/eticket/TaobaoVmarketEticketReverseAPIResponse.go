package eticket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketReverseAPIResponse
电子凭证冲正接口 API返回值
taobao.vmarket.eticket.reverse

电子凭证平台冲正接口 */
type TaobaoVmarketEticketReverseAPIResponse struct {
	model.CommonResponse
	TaobaoVmarketEticketReverseAPIResponseModel
}

// TaobaoVmarketEticketReverseAPIResponseModel is 电子凭证冲正接口 成功返回结果
type TaobaoVmarketEticketReverseAPIResponseModel struct {
	XMLName xml.Name `xml:"vmarket_eticket_reverse_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 0:失败，1:成功
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 宝贝标题
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// 整个订单的剩余可核销数量
	LeftNum int64 `json:"left_num,omitempty" xml:"left_num,omitempty"`
}
