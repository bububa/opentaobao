package ascp

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponseModel is 取消商品分销 成功返回结果
type AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_item_distribution_batch_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	CancelDistributeResponse *TopResponse `json:"cancel_distribute_response,omitempty" xml:"cancel_distribute_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CancelDistributeResponse = nil
}

var poolAlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse)
	},
}

// GetAlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse
func GetAlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse() *AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse {
	return poolAlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse.Get().(*AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse)
}

// ReleaseAlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse 将 AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse(v *AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse.Put(v)
}
