package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappuserphonegetAPIResponse 获取当前授权用户手机号码 API返回值
// taobao.miniapp.user.phone.get
//
// 在商家应用中，获取当前授权用户手机号码
type TaobaominiappuserphonegetAPIResponse struct {
	model.CommonResponse
	TaobaominiappuserphonegetAPIResponseModel
}

// TaobaominiappuserphonegetAPIResponseModel is 获取当前授权用户手机号码 成功返回结果
type TaobaominiappuserphonegetAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_user_phone_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户手机号码
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
}
