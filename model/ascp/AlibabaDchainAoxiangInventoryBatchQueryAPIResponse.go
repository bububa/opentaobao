package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangInventoryBatchQueryAPIResponse 批量查询库存 API返回值
// alibaba.dchain.aoxiang.inventory.batch.query
//
// 批量查询库存
type AlibabaDchainAoxiangInventoryBatchQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangInventoryBatchQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangInventoryBatchQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangInventoryBatchQueryAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangInventoryBatchQueryAPIResponseModel is 批量查询库存 成功返回结果
type AlibabaDchainAoxiangInventoryBatchQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_inventory_batch_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	BatchQueryInventoryResponse *BatchQueryInventoryResponse `json:"batch_query_inventory_response,omitempty" xml:"batch_query_inventory_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangInventoryBatchQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BatchQueryInventoryResponse = nil
}

var poolAlibabaDchainAoxiangInventoryBatchQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangInventoryBatchQueryAPIResponse)
	},
}

// GetAlibabaDchainAoxiangInventoryBatchQueryAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangInventoryBatchQueryAPIResponse
func GetAlibabaDchainAoxiangInventoryBatchQueryAPIResponse() *AlibabaDchainAoxiangInventoryBatchQueryAPIResponse {
	return poolAlibabaDchainAoxiangInventoryBatchQueryAPIResponse.Get().(*AlibabaDchainAoxiangInventoryBatchQueryAPIResponse)
}

// ReleaseAlibabaDchainAoxiangInventoryBatchQueryAPIResponse 将 AlibabaDchainAoxiangInventoryBatchQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangInventoryBatchQueryAPIResponse(v *AlibabaDchainAoxiangInventoryBatchQueryAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangInventoryBatchQueryAPIResponse.Put(v)
}
