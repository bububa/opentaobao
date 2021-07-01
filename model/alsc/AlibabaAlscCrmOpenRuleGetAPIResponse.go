package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询规则 API返回值 
alibaba.alsc.crm.open.rule.get

查询会员规则
*/
type AlibabaAlscCrmOpenRuleGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmOpenRuleGetAPIResponseModel
}

// 查询规则 成功返回结果
type AlibabaAlscCrmOpenRuleGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_open_rule_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口结果
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
