package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的会员价规则 APIResponse
alibaba.alsc.crm.rule.querympricerule

查询品牌下的会员价规则
*/
type AlibabaAlscCrmRuleQuerympriceruleAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alsc_crm_rule_querympricerule_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"