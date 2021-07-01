package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest
alibaba.alihealth.baby.baseinfo.order.sync API请求
alibaba.alihealth.baby.baseinfo.order.sync

育学园将订单信息回传给我们 */
type AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest struct {
	model.Params
	// 健康id
	_tpUserId int64
	// 商品id
	_commodityId string
	// 商品名称
	_commodityName string
	// 价钱
	_amount *BigDecimal
	// 状态，1是已支付，2是已退款
	_status int64
	// 订单时间
	_orderTime string
	// 订单id
	_orderId string
}

// NewAlibabaAlihealthBabyBaseinfoOrderSyncRequest 初始化AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest对象
func NewAlibabaAlihealthBabyBaseinfoOrderSyncRequest() *AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest {
	return &AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.baby.baseinfo.order.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TpUserId Setter
// 健康id
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest) SetTpUserId(_tpUserId int64) error {
	r._tpUserId = _tpUserId
	r.Set("tp_user_id", _tpUserId)
	return nil
}

// Get TpUserId Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest) GetTpUserId() int64 {
	return r._tpUserId
}

// Set is CommodityId Setter
// 商品id
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest) SetCommodityId(_commodityId string) error {
	r._commodityId = _commodityId
	r.Set("commodity_id", _commodityId)
	return nil
}

// Get CommodityId Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest) GetCommodityId() string {
	return r._commodityId
}

// Set is CommodityName Setter
// 商品名称
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest) SetCommodityName(_commodityName string) error {
	r._commodityName = _commodityName
	r.Set("commodity_name", _commodityName)
	return nil
}

// Get CommodityName Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest) GetCommodityName() string {
	return r._commodityName
}

// Set is Amount Setter
// 价钱
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest) SetAmount(_amount *BigDecimal) error {
	r._amount = _amount
	r.Set("amount", _amount)
	return nil
}

// Get Amount Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest) GetAmount() *BigDecimal {
	return r._amount
}

// Set is Status Setter
// 状态，1是已支付，2是已退款
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest) GetStatus() int64 {
	return r._status
}

// Set is OrderTime Setter
// 订单时间
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest) SetOrderTime(_orderTime string) error {
	r._orderTime = _orderTime
	r.Set("order_time", _orderTime)
	return nil
}

// Get OrderTime Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest) GetOrderTime() string {
	return r._orderTime
}

// Set is OrderId Setter
// 订单id
func (r *AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AlibabaAlihealthBabyBaseinfoOrderSyncAPIRequest) GetOrderId() string {
	return r._orderId
}
