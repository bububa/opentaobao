package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemDistributionDescribeAPIResponse 分销商品文描 API返回值
// alibaba.dchain.aoxiang.item.distribution.describe
//
// 分销商品文描
type AlibabaDchainAoxiangItemDistributionDescribeAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangItemDistributionDescribeAPIResponseModel
}

// AlibabaDchainAoxiangItemDistributionDescribeAPIResponseModel is 分销商品文描 成功返回结果
type AlibabaDchainAoxiangItemDistributionDescribeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_item_distribution_describe_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分销文描结果
	CreateItemDistributionRequest *TopResponse `json:"create_item_distribution_request,omitempty" xml:"create_item_distribution_request,omitempty"`
}
