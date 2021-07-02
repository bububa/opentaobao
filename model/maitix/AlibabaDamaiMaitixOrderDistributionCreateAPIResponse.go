package maitix

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixOrderDistributionCreateAPIResponse 大麦-新分销下单 API返回值
// alibaba.damai.maitix.order.distribution.create
//
// createDistributionOrder
type AlibabaDamaiMaitixOrderDistributionCreateAPIResponse struct {
	model.CommonResponse
	AlibabaDamaiMaitixOrderDistributionCreateAPIResponseModel
}

// AlibabaDamaiMaitixOrderDistributionCreateAPIResponseModel is 大麦-新分销下单 成功返回结果
type AlibabaDamaiMaitixOrderDistributionCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_order_distribution_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *MxResult `json:"result,omitempty" xml:"result,omitempty"`
}
