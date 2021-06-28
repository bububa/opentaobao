package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
用户标签判断接口 APIResponse
tmall.promotag.taguser.judge

查询用户是否有标签
*/
type TmallPromotagTaguserJudgeAPIResponse struct {
    model.CommonResponse
    TmallPromotagTaguserJudgeResponse
}

type TmallPromotagTaguserJudgeResponse struct {
    XMLName xml.Name `xml:"tmall_promotag_taguser_judge_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 服务调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 用户是否有标签
    
    HasTag   bool `json:"has_tag,omitempty" xml:"has_tag,omitempty"`

    
}
