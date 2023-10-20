package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelorderpayAPIRequest 阿信酒店分销订单支付 API请求
// taobao.alitrip.travel.axin.hotel.order.pay
//
// 阿信酒店分销订单支付
type TaobaoalitriptravelaxinhotelorderpayAPIRequest struct {
	model.Params
	// 分销商id
	_distributorTid int64
	// 采购单号
	_purchaseOrderId int64
}

// NewTaobaoalitriptravelaxinhotelorderpayRequest 初始化TaobaoalitriptravelaxinhotelorderpayAPIRequest对象
func NewTaobaoalitriptravelaxinhotelorderpayRequest() *TaobaoalitriptravelaxinhotelorderpayAPIRequest {
	return &TaobaoalitriptravelaxinhotelorderpayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelaxinhotelorderpayAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.order.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelaxinhotelorderpayAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelaxinhotelorderpayAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoalitriptravelaxinhotelorderpayAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoalitriptravelaxinhotelorderpayAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetPurchaseOrderId is PurchaseOrderId Setter
// 采购单号
func (r *TaobaoalitriptravelaxinhotelorderpayAPIRequest) SetPurchaseOrderId(_purchaseOrderId int64) error {
	r._purchaseOrderId = _purchaseOrderId
	r.Set("purchase_order_id", _purchaseOrderId)
	return nil
}

// GetPurchaseOrderId PurchaseOrderId Getter
func (r TaobaoalitriptravelaxinhotelorderpayAPIRequest) GetPurchaseOrderId() int64 {
	return r._purchaseOrderId
}
