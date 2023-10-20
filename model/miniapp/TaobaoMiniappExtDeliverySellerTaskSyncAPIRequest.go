package miniapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappextdeliverysellertasksyncAPIRequest 同步外投任务 API请求
// taobao.miniapp.ext.delivery.seller.task.sync
//
// 同步外投任务
type TaobaominiappextdeliverysellertasksyncAPIRequest struct {
	model.Params
	// 入参
	_sellerDeliveryTask *SellerDeliveryTaskDto
}

// NewTaobaominiappextdeliverysellertasksyncRequest 初始化TaobaominiappextdeliverysellertasksyncAPIRequest对象
func NewTaobaominiappextdeliverysellertasksyncRequest() *TaobaominiappextdeliverysellertasksyncAPIRequest {
	return &TaobaominiappextdeliverysellertasksyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappextdeliverysellertasksyncAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.ext.delivery.seller.task.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappextdeliverysellertasksyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappextdeliverysellertasksyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerDeliveryTask is SellerDeliveryTask Setter
// 入参
func (r *TaobaominiappextdeliverysellertasksyncAPIRequest) SetSellerDeliveryTask(_sellerDeliveryTask *SellerDeliveryTaskDto) error {
	r._sellerDeliveryTask = _sellerDeliveryTask
	r.Set("seller_delivery_task", _sellerDeliveryTask)
	return nil
}

// GetSellerDeliveryTask SellerDeliveryTask Getter
func (r TaobaominiappextdeliverysellertasksyncAPIRequest) GetSellerDeliveryTask() *SellerDeliveryTaskDto {
	return r._sellerDeliveryTask
}
