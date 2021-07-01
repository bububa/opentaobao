package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询品牌下的会员日规则 API返回值 
alibaba.alsc.crm.rule.querymdayerule

查询品牌下的会员日规则
*/
type AlibabaAlscCrmRuleQuerymdayeruleAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmRuleQuerymdayeruleAPIResponseModel
}

// 查询品牌下的会员日规则 成功返回结果
type AlibabaAlscCrmRuleQuerymdayeruleAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_rule_querymdayerule_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口结果
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
