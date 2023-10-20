package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangphysicsinventorybatchuploadasyncAPIRequest 实仓库存同步 API请求
// alibaba.dchain.aoxiang.physics.inventory.batch.upload.async
//
// 实仓库存同步
type AlibabadchainaoxiangphysicsinventorybatchuploadasyncAPIRequest struct {
	model.Params
	// 请求入参
	_inventoryUploadRequest *InventoryBatchUploadAsyncRequest
}

// NewAlibabadchainaoxiangphysicsinventorybatchuploadasyncRequest 初始化AlibabadchainaoxiangphysicsinventorybatchuploadasyncAPIRequest对象
func NewAlibabadchainaoxiangphysicsinventorybatchuploadasyncRequest() *AlibabadchainaoxiangphysicsinventorybatchuploadasyncAPIRequest {
	return &AlibabadchainaoxiangphysicsinventorybatchuploadasyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangphysicsinventorybatchuploadasyncAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.physics.inventory.batch.upload.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangphysicsinventorybatchuploadasyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangphysicsinventorybatchuploadasyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInventoryUploadRequest is InventoryUploadRequest Setter
// 请求入参
func (r *AlibabadchainaoxiangphysicsinventorybatchuploadasyncAPIRequest) SetInventoryUploadRequest(_inventoryUploadRequest *InventoryBatchUploadAsyncRequest) error {
	r._inventoryUploadRequest = _inventoryUploadRequest
	r.Set("inventory_upload_request", _inventoryUploadRequest)
	return nil
}

// GetInventoryUploadRequest InventoryUploadRequest Getter
func (r AlibabadchainaoxiangphysicsinventorybatchuploadasyncAPIRequest) GetInventoryUploadRequest() *InventoryBatchUploadAsyncRequest {
	return r._inventoryUploadRequest
}
