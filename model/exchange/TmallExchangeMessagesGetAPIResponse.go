package exchange

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallexchangemessagesgetAPIResponse 查询换货订单留言列表 API返回值
// tmall.exchange.messages.get
//
// 查询换货订单留言列表
type TmallexchangemessagesgetAPIResponse struct {
	model.CommonResponse
	TmallexchangemessagesgetAPIResponseModel
}

// TmallexchangemessagesgetAPIResponseModel is 查询换货订单留言列表 成功返回结果
type TmallexchangemessagesgetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_exchange_messages_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *RefundMessageResult `json:"result,omitempty" xml:"result,omitempty"`
}
