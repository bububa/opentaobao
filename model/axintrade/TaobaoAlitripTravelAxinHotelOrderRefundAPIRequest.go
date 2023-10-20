package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelorderrefundAPIRequest 阿信酒店分销订单退款API API请求
// taobao.alitrip.travel.axin.hotel.order.refund
//
// 阿信酒店分销订单退款
type TaobaoalitriptravelaxinhotelorderrefundAPIRequest struct {
	model.Params
	// 退款原因
	_reason string
	// 备注
	_remark string
	// 资源渠道
	_resourceChannel string
	// 分销商id
	_distributorTid int64
	// 采购单号
	_purchaseOrderId int64
}

// NewTaobaoalitriptravelaxinhotelorderrefundRequest 初始化TaobaoalitriptravelaxinhotelorderrefundAPIRequest对象
func NewTaobaoalitriptravelaxinhotelorderrefundRequest() *TaobaoalitriptravelaxinhotelorderrefundAPIRequest {
	return &TaobaoalitriptravelaxinhotelorderrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelaxinhotelorderrefundAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.order.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelaxinhotelorderrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelaxinhotelorderrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReason is Reason Setter
// 退款原因
func (r *TaobaoalitriptravelaxinhotelorderrefundAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r TaobaoalitriptravelaxinhotelorderrefundAPIRequest) GetReason() string {
	return r._reason
}

// SetRemark is Remark Setter
// 备注
func (r *TaobaoalitriptravelaxinhotelorderrefundAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoalitriptravelaxinhotelorderrefundAPIRequest) GetRemark() string {
	return r._remark
}

// SetResourceChannel is ResourceChannel Setter
// 资源渠道
func (r *TaobaoalitriptravelaxinhotelorderrefundAPIRequest) SetResourceChannel(_resourceChannel string) error {
	r._resourceChannel = _resourceChannel
	r.Set("resource_channel", _resourceChannel)
	return nil
}

// GetResourceChannel ResourceChannel Getter
func (r TaobaoalitriptravelaxinhotelorderrefundAPIRequest) GetResourceChannel() string {
	return r._resourceChannel
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoalitriptravelaxinhotelorderrefundAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoalitriptravelaxinhotelorderrefundAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetPurchaseOrderId is PurchaseOrderId Setter
// 采购单号
func (r *TaobaoalitriptravelaxinhotelorderrefundAPIRequest) SetPurchaseOrderId(_purchaseOrderId int64) error {
	r._purchaseOrderId = _purchaseOrderId
	r.Set("purchase_order_id", _purchaseOrderId)
	return nil
}

// GetPurchaseOrderId PurchaseOrderId Getter
func (r TaobaoalitriptravelaxinhotelorderrefundAPIRequest) GetPurchaseOrderId() int64 {
	return r._purchaseOrderId
}
