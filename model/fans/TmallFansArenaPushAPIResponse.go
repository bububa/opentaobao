package fans

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallfansarenapushAPIResponse 消息推送 API返回值
// tmall.fans.arena.push
//
// 超级擂台消息推送
type TmallfansarenapushAPIResponse struct {
	model.CommonResponse
	TmallfansarenapushAPIResponseModel
}

// TmallfansarenapushAPIResponseModel is 消息推送 成功返回结果
type TmallfansarenapushAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_fans_arena_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *FansResult `json:"result,omitempty" xml:"result,omitempty"`
}
