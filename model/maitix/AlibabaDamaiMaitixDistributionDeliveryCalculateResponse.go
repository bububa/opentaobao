package maitix

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
计算渠道用户下单快递费 APIResponse
alibaba.damai.maitix.distribution.delivery.calculate

计算渠道用户下单快递费
*/
type AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMaitixDistributionDeliveryCalculateResponse
}

type AlibabaDamaiMaitixDistributionDeliveryCalculateResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_maitix_distribution_delivery_calculate_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaDamaiMaitixDistributionDeliveryCalculateResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
