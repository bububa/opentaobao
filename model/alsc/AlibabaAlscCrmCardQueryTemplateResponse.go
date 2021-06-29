package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询卡模板详情 APIResponse
alibaba.alsc.crm.card.query.template

查询卡模板详情
*/
type AlibabaAlscCrmCardQueryTemplateAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmCardQueryTemplateResponse
}

type AlibabaAlscCrmCardQueryTemplateResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_card_query_template_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口结果
    
    Result   *CommonResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
