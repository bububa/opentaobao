package openim

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimImmsgPushAPIResponse openim标准消息发送 API返回值
// taobao.openim.immsg.push
//
// 服务端对openim用户发送标准消息，包括文字、语音、图片等。
type TaobaoOpenimImmsgPushAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimImmsgPushAPIResponseModel
}

// TaobaoOpenimImmsgPushAPIResponseModel is openim标准消息发送 成功返回结果
type TaobaoOpenimImmsgPushAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_immsg_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 消息id，方便定位问题
	Msgid int64 `json:"msgid,omitempty" xml:"msgid,omitempty"`
}
