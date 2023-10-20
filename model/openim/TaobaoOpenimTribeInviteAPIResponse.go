package openim

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimtribeinviteAPIResponse OPENIM群邀请加入 API返回值
// taobao.openim.tribe.invite
//
// OPENIM群邀请加入接口
type TaobaoopenimtribeinviteAPIResponse struct {
	model.CommonResponse
	TaobaoopenimtribeinviteAPIResponseModel
}

// TaobaoopenimtribeinviteAPIResponseModel is OPENIM群邀请加入 成功返回结果
type TaobaoopenimtribeinviteAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribe_invite_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 群服务code
	TribeCode int64 `json:"tribe_code,omitempty" xml:"tribe_code,omitempty"`
}
