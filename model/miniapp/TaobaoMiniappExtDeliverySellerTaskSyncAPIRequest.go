package miniapp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappExtDeliverySellerTaskSyncAPIRequest 同步外投任务 API请求
// taobao.miniapp.ext.delivery.seller.task.sync
//
// 同步外投任务
type TaobaoMiniappExtDeliverySellerTaskSyncAPIRequest struct {
	model.Params
	// 入参
	_sellerDeliveryTask *SellerDeliveryTaskDto
}

// NewTaobaoMiniappExtDeliverySellerTaskSyncRequest 初始化TaobaoMiniappExtDeliverySellerTaskSyncAPIRequest对象
func NewTaobaoMiniappExtDeliverySellerTaskSyncRequest() *TaobaoMiniappExtDeliverySellerTaskSyncAPIRequest {
	return &TaobaoMiniappExtDeliverySellerTaskSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappExtDeliverySellerTaskSyncAPIRequest) Reset() {
	r._sellerDeliveryTask = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappExtDeliverySellerTaskSyncAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.ext.delivery.seller.task.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappExtDeliverySellerTaskSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappExtDeliverySellerTaskSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerDeliveryTask is SellerDeliveryTask Setter
// 入参
func (r *TaobaoMiniappExtDeliverySellerTaskSyncAPIRequest) SetSellerDeliveryTask(_sellerDeliveryTask *SellerDeliveryTaskDto) error {
	r._sellerDeliveryTask = _sellerDeliveryTask
	r.Set("seller_delivery_task", _sellerDeliveryTask)
	return nil
}

// GetSellerDeliveryTask SellerDeliveryTask Getter
func (r TaobaoMiniappExtDeliverySellerTaskSyncAPIRequest) GetSellerDeliveryTask() *SellerDeliveryTaskDto {
	return r._sellerDeliveryTask
}

var poolTaobaoMiniappExtDeliverySellerTaskSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappExtDeliverySellerTaskSyncRequest()
	},
}

// GetTaobaoMiniappExtDeliverySellerTaskSyncRequest 从 sync.Pool 获取 TaobaoMiniappExtDeliverySellerTaskSyncAPIRequest
func GetTaobaoMiniappExtDeliverySellerTaskSyncAPIRequest() *TaobaoMiniappExtDeliverySellerTaskSyncAPIRequest {
	return poolTaobaoMiniappExtDeliverySellerTaskSyncAPIRequest.Get().(*TaobaoMiniappExtDeliverySellerTaskSyncAPIRequest)
}

// ReleaseTaobaoMiniappExtDeliverySellerTaskSyncAPIRequest 将 TaobaoMiniappExtDeliverySellerTaskSyncAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappExtDeliverySellerTaskSyncAPIRequest(v *TaobaoMiniappExtDeliverySellerTaskSyncAPIRequest) {
	v.Reset()
	poolTaobaoMiniappExtDeliverySellerTaskSyncAPIRequest.Put(v)
}
