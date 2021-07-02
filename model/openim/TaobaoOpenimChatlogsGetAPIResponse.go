package openim

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimChatlogsGetAPIResponse openim聊天记录查询接口 API返回值
// taobao.openim.chatlogs.get
//
// 查询openim账号聊天记录
type TaobaoOpenimChatlogsGetAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimChatlogsGetAPIResponseModel
}

// TaobaoOpenimChatlogsGetAPIResponseModel is openim聊天记录查询接口 成功返回结果
type TaobaoOpenimChatlogsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_chatlogs_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 聊天记录查询结果
	Result *RoamingMessageResult `json:"result,omitempty" xml:"result,omitempty"`
}
