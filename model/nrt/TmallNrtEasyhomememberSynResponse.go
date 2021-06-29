package nrt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
会员信息同 APIResponse
tmall.nrt.easyhomemember.syn

居然之家将会员信息同步到零售中台 包含基本的会员信息
*/
type TmallNrtEasyhomememberSynAPIResponse struct {
    model.CommonResponse
    TmallNrtEasyhomememberSynResponse
}

type TmallNrtEasyhomememberSynResponse struct {
    XMLName xml.Name `xml:"tmall_nrt_easyhomemember_syn_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ResultDO `json:"result,omitempty" xml:"result,omitempty"`

    
}
