package nrt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新零售会员同步接口 APIResponse
tmall.nrt.member.synchronize

新零售会员上翻接口，商家的会员信息同步至阿里侧
*/
type TmallNrtMemberSynchronizeAPIResponse struct {
    model.CommonResponse
    TmallNrtMemberSynchronizeResponse
}

type TmallNrtMemberSynchronizeResponse struct {
    XMLName xml.Name `xml:"tmall_nrt_member_synchronize_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    Data   *MemberSynResponse `json:"data,omitempty" xml:"data,omitempty"`

    
}
