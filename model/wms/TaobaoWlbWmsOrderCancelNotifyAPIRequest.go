package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbwmsordercancelnotifyAPIRequest 单据取消接口 API请求
// taobao.wlb.wms.order.cancel.notify
//
// 单据取消接口
type TaobaowlbwmsordercancelnotifyAPIRequest struct {
	model.Params
	// 订单类型
	_orderCode string
	// 单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、301 调拨出库单、901普通出库单、903 其他出库单、201 一般交易出库单 202 B2B交易出库单 502 换货出库单 503 补发出库单、1102 仓内加工作业单、 701 盘亏单、702盘盈单、711 库存状态调整单
	_orderType string
	// 仓库编码
	_storeCode string
}

// NewTaobaowlbwmsordercancelnotifyRequest 初始化TaobaowlbwmsordercancelnotifyAPIRequest对象
func NewTaobaowlbwmsordercancelnotifyRequest() *TaobaowlbwmsordercancelnotifyAPIRequest {
	return &TaobaowlbwmsordercancelnotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbwmsordercancelnotifyAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.order.cancel.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbwmsordercancelnotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbwmsordercancelnotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 订单类型
func (r *TaobaowlbwmsordercancelnotifyAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaowlbwmsordercancelnotifyAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetOrderType is OrderType Setter
// 单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、301 调拨出库单、901普通出库单、903 其他出库单、201 一般交易出库单 202 B2B交易出库单 502 换货出库单 503 补发出库单、1102 仓内加工作业单、 701 盘亏单、702盘盈单、711 库存状态调整单
func (r *TaobaowlbwmsordercancelnotifyAPIRequest) SetOrderType(_orderType string) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r TaobaowlbwmsordercancelnotifyAPIRequest) GetOrderType() string {
	return r._orderType
}

// SetStoreCode is StoreCode Setter
// 仓库编码
func (r *TaobaowlbwmsordercancelnotifyAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaowlbwmsordercancelnotifyAPIRequest) GetStoreCode() string {
	return r._storeCode
}
