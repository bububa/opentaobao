package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询运营计划 APIResponse
alibaba.alsc.crm.rule.queryoptplan

查询运营计划
*/
type AlibabaAlscCrmRuleQueryoptplanAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alsc_crm_rule_queryoptplan_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 请求参数
    
    Result   *CommonPageResult `json:"result,omitempty" xml:"