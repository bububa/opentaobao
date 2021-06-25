package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询会员卡信息 APIResponse
alibaba.wdk.member.card.get

根据会员卡查询会员信息
*/
type AlibabaWdkMemberCardGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMemberCardGetResponse `json:"alibaba_wdk_member_card_get_response,omitempty"`
}

type AlibabaWdkMemberCardGetResponse struct {

    // 结果
    ApiResult   *AlibabaWdkMemberCardGetApiResult `json:"api_result,omitempty"`

}
