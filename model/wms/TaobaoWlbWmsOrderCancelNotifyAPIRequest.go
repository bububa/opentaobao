package wms

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsOrderCancelNotifyAPIRequest 单据取消接口 API请求
// taobao.wlb.wms.order.cancel.notify
//
// 单据取消接口
type TaobaoWlbWmsOrderCancelNotifyAPIRequest struct {
	model.Params
	// 订单类型
	_orderCode string
	// 单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、301 调拨出库单、901普通出库单、903 其他出库单、201 一般交易出库单 202 B2B交易出库单 502 换货出库单 503 补发出库单、1102 仓内加工作业单、 701 盘亏单、702盘盈单、711 库存状态调整单
	_orderType string
	// 仓库编码
	_storeCode string
}

// NewTaobaoWlbWmsOrderCancelNotifyRequest 初始化TaobaoWlbWmsOrderCancelNotifyAPIRequest对象
func NewTaobaoWlbWmsOrderCancelNotifyRequest() *TaobaoWlbWmsOrderCancelNotifyAPIRequest {
	return &TaobaoWlbWmsOrderCancelNotifyAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbWmsOrderCancelNotifyAPIRequest) Reset() {
	r._orderCode = ""
	r._orderType = ""
	r._storeCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsOrderCancelNotifyAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.order.cancel.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsOrderCancelNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbWmsOrderCancelNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 订单类型
func (r *TaobaoWlbWmsOrderCancelNotifyAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaoWlbWmsOrderCancelNotifyAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetOrderType is OrderType Setter
// 单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、301 调拨出库单、901普通出库单、903 其他出库单、201 一般交易出库单 202 B2B交易出库单 502 换货出库单 503 补发出库单、1102 仓内加工作业单、 701 盘亏单、702盘盈单、711 库存状态调整单
func (r *TaobaoWlbWmsOrderCancelNotifyAPIRequest) SetOrderType(_orderType string) error {
	r._orderType = _orderType
	r.Set("order_type", _orderType)
	return nil
}

// GetOrderType OrderType Getter
func (r TaobaoWlbWmsOrderCancelNotifyAPIRequest) GetOrderType() string {
	return r._orderType
}

// SetStoreCode is StoreCode Setter
// 仓库编码
func (r *TaobaoWlbWmsOrderCancelNotifyAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoWlbWmsOrderCancelNotifyAPIRequest) GetStoreCode() string {
	return r._storeCode
}

var poolTaobaoWlbWmsOrderCancelNotifyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbWmsOrderCancelNotifyRequest()
	},
}

// GetTaobaoWlbWmsOrderCancelNotifyRequest 从 sync.Pool 获取 TaobaoWlbWmsOrderCancelNotifyAPIRequest
func GetTaobaoWlbWmsOrderCancelNotifyAPIRequest() *TaobaoWlbWmsOrderCancelNotifyAPIRequest {
	return poolTaobaoWlbWmsOrderCancelNotifyAPIRequest.Get().(*TaobaoWlbWmsOrderCancelNotifyAPIRequest)
}

// ReleaseTaobaoWlbWmsOrderCancelNotifyAPIRequest 将 TaobaoWlbWmsOrderCancelNotifyAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbWmsOrderCancelNotifyAPIRequest(v *TaobaoWlbWmsOrderCancelNotifyAPIRequest) {
	v.Reset()
	poolTaobaoWlbWmsOrderCancelNotifyAPIRequest.Put(v)
}
