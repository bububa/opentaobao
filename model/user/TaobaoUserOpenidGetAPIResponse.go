package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouseropenidgetAPIResponse 查询用户openId API返回值
// taobao.user.openid.get
//
// 查询用户openId
type TaobaouseropenidgetAPIResponse struct {
	model.CommonResponse
	TaobaouseropenidgetAPIResponseModel
}

// TaobaouseropenidgetAPIResponseModel is 查询用户openId 成功返回结果
type TaobaouseropenidgetAPIResponseModel struct {
	XMLName xml.Name `xml:"user_openid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 对应账号的OpenUID
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
}
