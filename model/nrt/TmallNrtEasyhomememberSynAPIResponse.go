package nrt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
会员信息同 API返回值 
tmall.nrt.easyhomemember.syn

居然之家将会员信息同步到零售中台 包含基本的会员信息
*/
type TmallNrtEasyhomememberSynAPIResponse struct {
    model.CommonResponse
    TmallNrtEasyhomememberSynAPIResponseModel
}

// 会员信息同 成功返回结果
type TmallNrtEasyhomememberSynAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_nrt_easyhomemember_syn_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
