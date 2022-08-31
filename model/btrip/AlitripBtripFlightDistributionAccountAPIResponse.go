package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightDistributionAccountAPIResponse 机票分销企业或者tmc企业预存or月结账户查询接口 API返回值
// alitrip.btrip.flight.distribution.account
//
// 机票分销企业或者tmc企业预存or月结账户查询
type AlitripBtripFlightDistributionAccountAPIResponse struct {
	model.CommonResponse
	AlitripBtripFlightDistributionAccountAPIResponseModel
}

// AlitripBtripFlightDistributionAccountAPIResponseModel is 机票分销企业或者tmc企业预存or月结账户查询接口 成功返回结果
type AlitripBtripFlightDistributionAccountAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_flight_distribution_account_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应结果
	Result *BtripAccountPrestoreRs `json:"result,omitempty" xml:"result,omitempty"`
}
