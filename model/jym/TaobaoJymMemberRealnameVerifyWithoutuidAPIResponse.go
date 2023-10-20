package jym

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaojymmemberrealnameverifywithoutuidAPIResponse 用户实名认证 API返回值
// taobao.jym.member.realname.verify.withoutuid
//
// 开放用户实名认证接口使用
type TaobaojymmemberrealnameverifywithoutuidAPIResponse struct {
	model.CommonResponse
	TaobaojymmemberrealnameverifywithoutuidAPIResponseModel
}

// TaobaojymmemberrealnameverifywithoutuidAPIResponseModel is 用户实名认证 成功返回结果
type TaobaojymmemberrealnameverifywithoutuidAPIResponseModel struct {
	XMLName xml.Name `xml:"jym_member_realname_verify_withoutuid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 实名认证结果
	Result *TaobaojymmemberrealnameverifywithoutuidResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
