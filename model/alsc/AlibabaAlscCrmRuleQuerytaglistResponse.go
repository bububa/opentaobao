package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询标签列表 APIResponse
alibaba.alsc.crm.rule.querytaglist

查询标签列表
*/
type AlibabaAlscCrmRuleQuerytaglistAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRuleQuerytaglistResponse
}

type AlibabaAlscCrmRuleQuerytaglistResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_rule_querytaglist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 分页返回模型
    
    Result   *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
