package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest 阿信酒店分销订单退款API API请求
// taobao.alitrip.travel.axin.hotel.order.refund
//
// 阿信酒店分销订单退款
type TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest struct {
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

// NewTaobaoAlitripTravelAxinHotelOrderRefundRequest 初始化TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelOrderRefundRequest() *TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest {
	return &TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.order.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReason is Reason Setter
// 退款原因
func (r *TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest) GetReason() string {
	return r._reason
}

// SetRemark is Remark Setter
// 备注
func (r *TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest) GetRemark() string {
	return r._remark
}

// SetResourceChannel is ResourceChannel Setter
// 资源渠道
func (r *TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest) SetResourceChannel(_resourceChannel string) error {
	r._resourceChannel = _resourceChannel
	r.Set("resource_channel", _resourceChannel)
	return nil
}

// GetResourceChannel ResourceChannel Getter
func (r TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest) GetResourceChannel() string {
	return r._resourceChannel
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetPurchaseOrderId is PurchaseOrderId Setter
// 采购单号
func (r *TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest) SetPurchaseOrderId(_purchaseOrderId int64) error {
	r._purchaseOrderId = _purchaseOrderId
	r.Set("purchase_order_id", _purchaseOrderId)
	return nil
}

// GetPurchaseOrderId PurchaseOrderId Getter
func (r TaobaoAlitripTravelAxinHotelOrderRefundAPIRequest) GetPurchaseOrderId() int64 {
	return r._purchaseOrderId
}
