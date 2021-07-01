package iot

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TmallDeviceMemberIdentityGetAPIResponse
智能硬件会员判断 API返回值
tmall.device.member.identity.get

用来识别该用户是否是商家会员· */
type TmallDeviceMemberIdentityGetAPIResponse struct {
	model.CommonResponse
	TmallDeviceMemberIdentityGetAPIResponseModel
}

// TmallDeviceMemberIdentityGetAPIResponseModel is 智能硬件会员判断 成功返回结果
type TmallDeviceMemberIdentityGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_device_member_identity_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallDeviceMemberIdentityGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
