package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaonextonelogisticswarehouseupdateAPIRequest AG退货入仓状态写接口 API请求
// taobao.nextone.logistics.warehouse.update
//
// 商家上传退货入仓状态给ag
type TaobaonextonelogisticswarehouseupdateAPIRequest struct {
	model.Params
	// 退款编号
	_refundId int64
	// 退货入仓状态 1.已入仓
	_warehouseStatus int64
}

// NewTaobaonextonelogisticswarehouseupdateRequest 初始化TaobaonextonelogisticswarehouseupdateAPIRequest对象
func NewTaobaonextonelogisticswarehouseupdateRequest() *TaobaonextonelogisticswarehouseupdateAPIRequest {
	return &TaobaonextonelogisticswarehouseupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaonextonelogisticswarehouseupdateAPIRequest) GetApiMethodName() string {
	return "taobao.nextone.logistics.warehouse.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaonextonelogisticswarehouseupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaonextonelogisticswarehouseupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundId is RefundId Setter
// 退款编号
func (r *TaobaonextonelogisticswarehouseupdateAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaonextonelogisticswarehouseupdateAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// SetWarehouseStatus is WarehouseStatus Setter
// 退货入仓状态 1.已入仓
func (r *TaobaonextonelogisticswarehouseupdateAPIRequest) SetWarehouseStatus(_warehouseStatus int64) error {
	r._warehouseStatus = _warehouseStatus
	r.Set("warehouse_status", _warehouseStatus)
	return nil
}

// GetWarehouseStatus WarehouseStatus Getter
func (r TaobaonextonelogisticswarehouseupdateAPIRequest) GetWarehouseStatus() int64 {
	return r._warehouseStatus
}
