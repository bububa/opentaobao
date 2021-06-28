package mei

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
支付码获取会员信息 APIResponse
tmall.mei.crm.member.getbypaycode

通过支付码获取会员信息
*/
type TmallMeiCrmMemberGetbypaycodeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_mei_crm_member_getbypaycode_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *ResultDTO `json:"result,omitempty" xml:"