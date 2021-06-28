package ticket

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【门票API2.0】景点门票规则维护接口 APIResponse
alitrip.ticket.rule.upload

景点门票规则维护接口。该接口同时支持新发规则和编辑现有规则，如果out_rule_id下没有发布过规则，则系统将判断为新发一个规则，否则认为是编辑现有规则。
对于新发布规则的情况，有些参数是必填的，请仔细查看各字段说明。对于编辑的情况，除out_rule_id外都是可选，编辑情况支持增量更新（某个参数不传则使用该规则上原有值）
*/
type AlitripTicketRuleUploadAPIResponse struct {
    model.CommonResponse
    AlitripTicketRuleUploadResponse
}

type AlitripTicketRuleUploadResponse struct {
    XMLName xml.Name `xml:"alitrip_ticket_rule_upload_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AlitripTicketRuleUploadResultSet `json:"result,omitempty" xml:"result,omitempty"`

    
}
