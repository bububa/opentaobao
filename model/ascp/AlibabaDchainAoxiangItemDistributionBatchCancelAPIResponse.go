package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse 取消商品分销 API返回值
// alibaba.dchain.aoxiang.item.distribution.batch.cancel
//
// 取消商品分销
type AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponseModel
}

// AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponseModel is 取消商品分销 成功返回结果
type AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_item_distribution_batch_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	CancelDistributeResponse *TopResponse `json:"cancel_distribute_response,omitempty" xml:"cancel_distribute_response,omitempty"`
}
