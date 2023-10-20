package tbuser

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoUserAvatarGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUserAvatarGetAPIResponseModel).Reset()
}

// TaobaoUserAvatarGetAPIResponseModel is 淘宝用户头像查询 成功返回结果
type TaobaoUserAvatarGetAPIResponseModel struct {
	XMLName xml.Name `xml:"user_avatar_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户头像地址
	Avatar string `json:"avatar,omitempty" xml:"avatar,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUserAvatarGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Avatar = ""
}

var poolTaobaoUserAvatarGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUserAvatarGetAPIResponse)
	},
}

// GetTaobaoUserAvatarGetAPIResponse 从 sync.Pool 获取 TaobaoUserAvatarGetAPIResponse
func GetTaobaoUserAvatarGetAPIResponse() *TaobaoUserAvatarGetAPIResponse {
	return poolTaobaoUserAvatarGetAPIResponse.Get().(*TaobaoUserAvatarGetAPIResponse)
}

// ReleaseTaobaoUserAvatarGetAPIResponse 将 TaobaoUserAvatarGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUserAvatarGetAPIResponse(v *TaobaoUserAvatarGetAPIResponse) {
	v.Reset()
	poolTaobaoUserAvatarGetAPIResponse.Put(v)
}
