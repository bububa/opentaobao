package westcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询是否是亲橙里会员 APIResponse
alibaba.westcrm.is.member.get

根据淘宝Id查询是否是亲橙里会员
*/
type AlibabaWestcrmIsMemberGetAPIResponse struct {
    model.CommonResponse
    AlibabaWestcrmIsMemberGetResponse
}

type AlibabaWestcrmIsMemberGetResponse struct {
    XMLName xml.Name `xml:"alibaba_westcrm_is_member_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // true:是亲橙里会员，false：不是亲橙里会员
    
    IsHaveMember   bool `json:"is_have_member,omitempty" xml:"is_have_member,omitempty"`

    
}
