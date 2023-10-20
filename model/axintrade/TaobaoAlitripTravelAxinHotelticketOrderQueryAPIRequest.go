package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelticketorderqueryAPIRequest 阿信度假交易订单查询接口 API请求
// taobao.alitrip.travel.axin.hotelticket.order.query
//
// 阿信度假交易订单查询接口
type TaobaoalitriptravelaxinhotelticketorderqueryAPIRequest struct {
	model.Params
	// 分销商id
	_distributorTid int64
	// 采购单id
	_purchaseSubOrderId int64
}

// NewTaobaoalitriptravelaxinhotelticketorderqueryRequest 初始化TaobaoalitriptravelaxinhotelticketorderqueryAPIRequest对象
func NewTaobaoalitriptravelaxinhotelticketorderqueryRequest() *TaobaoalitriptravelaxinhotelticketorderqueryAPIRequest {
	return &TaobaoalitriptravelaxinhotelticketorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelaxinhotelticketorderqueryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotelticket.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelaxinhotelticketorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelaxinhotelticketorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoalitriptravelaxinhotelticketorderqueryAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoalitriptravelaxinhotelticketorderqueryAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetPurchaseSubOrderId is PurchaseSubOrderId Setter
// 采购单id
func (r *TaobaoalitriptravelaxinhotelticketorderqueryAPIRequest) SetPurchaseSubOrderId(_purchaseSubOrderId int64) error {
	r._purchaseSubOrderId = _purchaseSubOrderId
	r.Set("purchase_sub_order_id", _purchaseSubOrderId)
	return nil
}

// GetPurchaseSubOrderId PurchaseSubOrderId Getter
func (r TaobaoalitriptravelaxinhotelticketorderqueryAPIRequest) GetPurchaseSubOrderId() int64 {
	return r._purchaseSubOrderId
}
