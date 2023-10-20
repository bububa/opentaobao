package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponse 实仓库存同步 API返回值
// alibaba.dchain.aoxiang.physics.inventory.batch.upload.async
//
// 实仓库存同步
type AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponseModel).Reset()
}

// AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponseModel is 实仓库存同步 成功返回结果
type AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_physics_inventory_batch_upload_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求出参
	InventoryUploadResponse *InventoryUploadAsyncResponse `json:"inventory_upload_response,omitempty" xml:"inventory_upload_response,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.InventoryUploadResponse = nil
}

var poolAlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponse)
	},
}

// GetAlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponse 从 sync.Pool 获取 AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponse
func GetAlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponse() *AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponse {
	return poolAlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponse.Get().(*AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponse)
}

// ReleaseAlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponse 将 AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponse(v *AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponse) {
	v.Reset()
	poolAlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponse.Put(v)
}
