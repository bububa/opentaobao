package maitix

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
加密招商一网能支付入参 APIResponse
alibaba.damai.maitix.distribution.cmb.paramencrypt

encryptParam4Cmb
*/
type AlibabaDamaiMaitixDistributionCmbParamencryptAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMaitixDistributionCmbParamencryptResponse
}

type AlibabaDamaiMaitixDistributionCmbParamencryptResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_maitix_distribution_cmb_paramencrypt_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *OpenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
