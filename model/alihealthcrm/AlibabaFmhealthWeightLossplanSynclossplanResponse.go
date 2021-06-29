package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
减重计划--同步减重计划 APIResponse
alibaba.fmhealth.weight.lossplan.synclossplan

减重计划--三方同步用户初始化减重计划给我们
*/
type AlibabaFmhealthWeightLossplanSynclossplanAPIResponse struct {
    model.CommonResponse
    AlibabaFmhealthWeightLossplanSynclossplanResponse
}

type AlibabaFmhealthWeightLossplanSynclossplanResponse struct {
    XMLName xml.Name `xml:"alibaba_fmhealth_weight_lossplan_synclossplan_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
