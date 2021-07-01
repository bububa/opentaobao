package alihealthcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
减重计划--同步减重计划 API返回值 
alibaba.fmhealth.weight.lossplan.synclossplan

减重计划--三方同步用户初始化减重计划给我们
*/
type AlibabaFmhealthWeightLossplanSynclossplanAPIResponse struct {
    model.CommonResponse
    AlibabaFmhealthWeightLossplanSynclossplanAPIResponseModel
}

// 减重计划--同步减重计划 成功返回结果
type AlibabaFmhealthWeightLossplanSynclossplanAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_fmhealth_weight_lossplan_synclossplan_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
