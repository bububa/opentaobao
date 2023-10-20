package media

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPictureUserinfoGetAPIResponse 查询图片空间用户的信息 API返回值
// taobao.picture.userinfo.get
//
// 查询用户的图片空间使用信息，包括：订购量，已使用容量，免费容量，总的可使用容量，订购有效期，剩余容量
type TaobaoPictureUserinfoGetAPIResponse struct {
	model.CommonResponse
	TaobaoPictureUserinfoGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoPictureUserinfoGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoPictureUserinfoGetAPIResponseModel).Reset()
}

// TaobaoPictureUserinfoGetAPIResponseModel is 查询图片空间用户的信息 成功返回结果
type TaobaoPictureUserinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"picture_userinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户使用图片空间的信息
	UserInfo *UserInfo `json:"user_info,omitempty" xml:"user_info,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoPictureUserinfoGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.UserInfo = nil
}

var poolTaobaoPictureUserinfoGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoPictureUserinfoGetAPIResponse)
	},
}

// GetTaobaoPictureUserinfoGetAPIResponse 从 sync.Pool 获取 TaobaoPictureUserinfoGetAPIResponse
func GetTaobaoPictureUserinfoGetAPIResponse() *TaobaoPictureUserinfoGetAPIResponse {
	return poolTaobaoPictureUserinfoGetAPIResponse.Get().(*TaobaoPictureUserinfoGetAPIResponse)
}

// ReleaseTaobaoPictureUserinfoGetAPIResponse 将 TaobaoPictureUserinfoGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoPictureUserinfoGetAPIResponse(v *TaobaoPictureUserinfoGetAPIResponse) {
	v.Reset()
	poolTaobaoPictureUserinfoGetAPIResponse.Put(v)
}
