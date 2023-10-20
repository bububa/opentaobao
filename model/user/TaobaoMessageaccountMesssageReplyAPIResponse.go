package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaomessageaccountmesssagereplyAPIResponse 消息号下行回复接口 API返回值
// taobao.messageaccount.messsage.reply
//
// 外部 isv 调用该进口来进行消息号消息的回复
type TaobaomessageaccountmesssagereplyAPIResponse struct {
	model.CommonResponse
	TaobaomessageaccountmesssagereplyAPIResponseModel
}

// TaobaomessageaccountmesssagereplyAPIResponseModel is 消息号下行回复接口 成功返回结果
type TaobaomessageaccountmesssagereplyAPIResponseModel struct {
	XMLName xml.Name `xml:"messageaccount_messsage_reply_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaomessageaccountmesssagereplyResult `json:"result,omitempty" xml:"result,omitempty"`
}
