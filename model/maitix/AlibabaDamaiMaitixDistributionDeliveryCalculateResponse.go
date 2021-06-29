package maitix

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
计算渠道用户下单快递费 API返回值 
alibaba.damai.maitix.distribution.delivery.calculate

计算渠道用户下单快递费
*/
type AlibabaDamaiMaitixDistributionDeliveryCalculateAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMaitixDistributionDeliveryCalculateResponse
}

// 计算渠道用户下单快递费 成功返回结果
type AlibabaDamaiMaitixDistributionDeliveryCalculateResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_maitix_distribution_delivery_calculate_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaDamaiMaitixDistributionDeliveryCalculateResult `json:"result,omitempty" xml:"result,omitempty"`
}
