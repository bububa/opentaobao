package alilabs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabuserprofilegetAPIResponse 查询用户信息 API返回值
// alibaba.ailab.user.profile.get
//
// 提供天猫精灵用户头像、昵称的查询接口，供本田车载天猫精灵使用
type AlibabaailabuserprofilegetAPIResponse struct {
	model.CommonResponse
	AlibabaailabuserprofilegetAPIResponseModel
}

// AlibabaailabuserprofilegetAPIResponseModel is 查询用户信息 成功返回结果
type AlibabaailabuserprofilegetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailab_user_profile_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口model
	Result *AlibabaailabuserprofilegetResult `json:"result,omitempty" xml:"result,omitempty"`
}
