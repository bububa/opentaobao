package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
用户标签判断接口 APIResponse
tmall.promotag.taguser.judge

查询用户是否有标签
*/
type TmallPromotagTaguserJudgeAPIResponse struct {
    model.CommonResponse
    Response *TmallPromotagTaguserJudgeResponse `json:"tmall_promotag_taguser_judge_response,omitempty"`
}

type TmallPromotagTaguserJudgeResponse struct {

    // 服务调用是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 用户是否有标签
    HasTag   bool `json:"has_tag,omitempty"`

}
