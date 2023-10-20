package maitix

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse 查询分销物流单 API返回值
// alibaba.damai.maitix.distribution.delivery.query
//
// 渠道查询物流订单
type AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponseModel
}

// AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponseModel is 查询分销物流单 成功返回结果
type AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_distribution_delivery_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
