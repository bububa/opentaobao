package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询标签列表 APIResponse
alibaba.alsc.crm.rule.querytaglist

查询标签列表
*/
type AlibabaAlscCrmRuleQuerytaglistAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlscCrmRuleQuerytaglistResponse `json:"alibaba_alsc_crm_rule_querytaglist_response,omitempty"`
}

type AlibabaAlscCrmRuleQuerytaglistResponse struct {

    // 分页返回模型
    Result   *CommonPageResult `json:"result,omitempty"`

}
