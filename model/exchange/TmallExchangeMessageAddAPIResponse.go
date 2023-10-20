package exchange

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallexchangemessageaddAPIResponse 卖家创建换货留言 API返回值
// tmall.exchange.message.add
//
// 卖家创建换货留言
type TmallexchangemessageaddAPIResponse struct {
	model.CommonResponse
	TmallexchangemessageaddAPIResponseModel
}

// TmallexchangemessageaddAPIResponseModel is 卖家创建换货留言 成功返回结果
type TmallexchangemessageaddAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_exchange_message_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TmallexchangemessageaddResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
