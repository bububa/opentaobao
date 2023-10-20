package tbuser

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUserAvatarGetAPIResponse 淘宝用户头像查询 API返回值
// taobao.user.avatar.get
//
// 根据混淆nick查询用户头像
type TaobaoUserAvatarGetAPIResponse struct {
	model.CommonResponse
	TaobaoUserAvatarGetAPIResponseModel
}

// TaobaoUserAvatarGetAPIResponseModel is 淘宝用户头像查询 成功返回结果
type TaobaoUserAvatarGetAPIResponseModel struct {
	XMLName xml.Name `xml:"user_avatar_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户头像地址
	Avatar string `json:"avatar,omitempty" xml:"avatar,omitempty"`
}
