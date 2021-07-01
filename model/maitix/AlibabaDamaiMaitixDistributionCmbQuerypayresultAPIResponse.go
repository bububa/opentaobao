package maitix

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询招行支付状态api API返回值 
alibaba.damai.maitix.distribution.cmb.querypayresult

queryPayResult
*/
type AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponseModel
}

// 查询招行支付状态api 成功返回结果
type AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_damai_maitix_distribution_cmb_querypayresult_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
