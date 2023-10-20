package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthbabybaseinfoordersyncAPIRequest alibaba.alihealth.baby.baseinfo.order.sync API请求
// alibaba.alihealth.baby.baseinfo.order.sync
//
// 育学园将订单信息回传给我们
type AlibabaalihealthbabybaseinfoordersyncAPIRequest struct {
	model.Params
	// 商品id
	_commodityId string
	// 商品名称
	_commodityName string
	// 订单时间
	_orderTime string
	// 订单id
	_orderId string
	// 健康id
	_tpUserId int64
	// 价钱
	_amount float64
	// 状态，1是已支付，2是已退款
	_status int64
}

// NewAlibabaalihealthbabybaseinfoordersyncRequest 初始化AlibabaalihealthbabybaseinfoordersyncAPIRequest对象
func NewAlibabaalihealthbabybaseinfoordersyncRequest() *AlibabaalihealthbabybaseinfoordersyncAPIRequest {
	return &AlibabaalihealthbabybaseinfoordersyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthbabybaseinfoordersyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.baby.baseinfo.order.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthbabybaseinfoordersyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthbabybaseinfoordersyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCommodityId is CommodityId Setter
// 商品id
func (r *AlibabaalihealthbabybaseinfoordersyncAPIRequest) SetCommodityId(_commodityId string) error {
	r._commodityId = _commodityId
	r.Set("commodity_id", _commodityId)
	return nil
}

// GetCommodityId CommodityId Getter
func (r AlibabaalihealthbabybaseinfoordersyncAPIRequest) GetCommodityId() string {
	return r._commodityId
}

// SetCommodityName is CommodityName Setter
// 商品名称
func (r *AlibabaalihealthbabybaseinfoordersyncAPIRequest) SetCommodityName(_commodityName string) error {
	r._commodityName = _commodityName
	r.Set("commodity_name", _commodityName)
	return nil
}

// GetCommodityName CommodityName Getter
func (r AlibabaalihealthbabybaseinfoordersyncAPIRequest) GetCommodityName() string {
	return r._commodityName
}

// SetOrderTime is OrderTime Setter
// 订单时间
func (r *AlibabaalihealthbabybaseinfoordersyncAPIRequest) SetOrderTime(_orderTime string) error {
	r._orderTime = _orderTime
	r.Set("order_time", _orderTime)
	return nil
}

// GetOrderTime OrderTime Getter
func (r AlibabaalihealthbabybaseinfoordersyncAPIRequest) GetOrderTime() string {
	return r._orderTime
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaalihealthbabybaseinfoordersyncAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaalihealthbabybaseinfoordersyncAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetTpUserId is TpUserId Setter
// 健康id
func (r *AlibabaalihealthbabybaseinfoordersyncAPIRequest) SetTpUserId(_tpUserId int64) error {
	r._tpUserId = _tpUserId
	r.Set("tp_user_id", _tpUserId)
	return nil
}

// GetTpUserId TpUserId Getter
func (r AlibabaalihealthbabybaseinfoordersyncAPIRequest) GetTpUserId() int64 {
	return r._tpUserId
}

// SetAmount is Amount Setter
// 价钱
func (r *AlibabaalihealthbabybaseinfoordersyncAPIRequest) SetAmount(_amount float64) error {
	r._amount = _amount
	r.Set("amount", _amount)
	return nil
}

// GetAmount Amount Getter
func (r AlibabaalihealthbabybaseinfoordersyncAPIRequest) GetAmount() float64 {
	return r._amount
}

// SetStatus is Status Setter
// 状态，1是已支付，2是已退款
func (r *AlibabaalihealthbabybaseinfoordersyncAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaalihealthbabybaseinfoordersyncAPIRequest) GetStatus() int64 {
	return r._status
}
