package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询会员等级规则 APIResponse
alibaba.alsc.crm.rule.level.querylevelrule

查询会员等级规则
*/
type AlibabaAlscCrmRuleLevelQuerylevelruleAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alsc_crm_rule_level_querylevelrule_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 会员等级规则
    
    Result   *CommonPageResult `json:"result,omitempty" xml:"