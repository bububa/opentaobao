package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
绑定会员 API返回值 
alibaba.mj.member.bindmember

用于绑定喵街数字化会员
*/
type AlibabaMjMemberBindmemberAPIResponse struct {
    model.CommonResponse
    AlibabaMjMemberBindmemberAPIResponseModel
}

// 绑定会员 成功返回结果
type AlibabaMjMemberBindmemberAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_mj_member_bindmember_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 1
    Result   *SingleResult `json:"result,omitempty" xml:"result,omitempty"`
}
