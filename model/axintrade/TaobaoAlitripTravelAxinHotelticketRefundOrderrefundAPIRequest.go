package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIRequest 阿信度假业务申请退款 API请求
// taobao.alitrip.travel.axin.hotelticket.refund.orderrefund
//
// 阿信度假业务申请退款
type TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIRequest struct {
	model.Params
	// 备注
	_remark string
	// 退款原因
	_reason string
	// 分销商ID
	_distributorTid int64
	// 采购单号
	_purchaseSubOrderId int64
}

// NewTaobaoalitriptravelaxinhotelticketrefundorderrefundRequest 初始化TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIRequest对象
func NewTaobaoalitriptravelaxinhotelticketrefundorderrefundRequest() *TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIRequest {
	return &TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotelticket.refund.orderrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRemark is Remark Setter
// 备注
func (r *TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIRequest) GetRemark() string {
	return r._remark
}

// SetReason is Reason Setter
// 退款原因
func (r *TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIRequest) GetReason() string {
	return r._reason
}

// SetDistributorTid is DistributorTid Setter
// 分销商ID
func (r *TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetPurchaseSubOrderId is PurchaseSubOrderId Setter
// 采购单号
func (r *TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIRequest) SetPurchaseSubOrderId(_purchaseSubOrderId int64) error {
	r._purchaseSubOrderId = _purchaseSubOrderId
	r.Set("purchase_sub_order_id", _purchaseSubOrderId)
	return nil
}

// GetPurchaseSubOrderId PurchaseSubOrderId Getter
func (r TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIRequest) GetPurchaseSubOrderId() int64 {
	return r._purchaseSubOrderId
}
