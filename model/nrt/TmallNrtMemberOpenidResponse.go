package nrt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据会员手机查询openId APIResponse
tmall.nrt.member.openid

根据会员手机查询openId
*/
type TmallNrtMemberOpenidAPIResponse struct {
    model.CommonResponse
    TmallNrtMemberOpenidResponse
}

type TmallNrtMemberOpenidResponse struct {
    XMLName xml.Name `xml:"tmall_nrt_member_openid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *TmallNrtMemberOpenidResultDo `json:"result,omitempty" xml:"result,omitempty"`

    
}
