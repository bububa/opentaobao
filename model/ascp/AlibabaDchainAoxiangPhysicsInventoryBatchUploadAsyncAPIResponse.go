package ascp

import (
	"encoding/xml"

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

// AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponseModel is 实仓库存同步 成功返回结果
type AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_physics_inventory_batch_upload_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求出参
	InventoryUploadResponse *InventoryUploadAsyncResponse `json:"inventory_upload_response,omitempty" xml:"inventory_upload_response,omitempty"`
}
