package mei

import (
    "github.com/bububa/opentaobao/model"
)

/* 
支付码获取会员信息 APIResponse
tmall.mei.crm.member.getbypaycode

通过支付码获取会员信息
*/
type TmallMeiCrmMemberGetbypaycodeAPIResponse struct {
    model.CommonResponse
    Response *TmallMeiCrmMemberGetbypaycodeResponse `json:"tmall_mei_crm_member_getbypaycode_response,omitempty"`
}

type TmallMeiCrmMemberGetbypaycodeResponse struct {

    // result
    Result   *ResultDTO `json:"result,omitempty"`

}
