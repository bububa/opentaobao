package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest 实仓库存同步 API请求
// alibaba.dchain.aoxiang.physics.inventory.batch.upload.async
//
// 实仓库存同步
type AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest struct {
	model.Params
	// 请求入参
	_inventoryUploadRequest *InventoryBatchUploadAsyncRequest
}

// NewAlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncRequest 初始化AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest对象
func NewAlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncRequest() *AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest {
	return &AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest) Reset() {
	r._inventoryUploadRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.physics.inventory.batch.upload.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInventoryUploadRequest is InventoryUploadRequest Setter
// 请求入参
func (r *AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest) SetInventoryUploadRequest(_inventoryUploadRequest *InventoryBatchUploadAsyncRequest) error {
	r._inventoryUploadRequest = _inventoryUploadRequest
	r.Set("inventory_upload_request", _inventoryUploadRequest)
	return nil
}

// GetInventoryUploadRequest InventoryUploadRequest Getter
func (r AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest) GetInventoryUploadRequest() *InventoryBatchUploadAsyncRequest {
	return r._inventoryUploadRequest
}

var poolAlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncRequest()
	},
}

// GetAlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncRequest 从 sync.Pool 获取 AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest
func GetAlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest() *AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest {
	return poolAlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest.Get().(*AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest)
}

// ReleaseAlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest 将 AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest(v *AlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangPhysicsInventoryBatchUploadAsyncAPIRequest.Put(v)
}
