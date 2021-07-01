package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
喵街会员是否绑定 API返回值 
alibaba.mj.member.hasbind

喵街检测用户是否为数字化会员
*/
type AlibabaMjMemberHasbindAPIResponse struct {
    model.CommonResponse
    AlibabaMjMemberHasbindAPIResponseModel
}

// 喵街会员是否绑定 成功返回结果
type AlibabaMjMemberHasbindAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_mj_member_hasbind_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *SingleResult `json:"result,omitempty" xml:"result,omitempty"`
}
