package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询会员卡信息 APIResponse
alibaba.wdk.member.card.get

根据会员卡查询会员信息
*/
type AlibabaWdkMemberCardGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_member_card_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    ApiResult   *AlibabaWdkMemberCardGetApiResult `json:"api_result,omitempty" xml:"