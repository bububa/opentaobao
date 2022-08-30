package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelOrderPayAPIRequest 阿信酒店分销订单支付 API请求
// taobao.alitrip.travel.axin.hotel.order.pay
//
// 阿信酒店分销订单支付
type TaobaoAlitripTravelAxinHotelOrderPayAPIRequest struct {
	model.Params
	// 分销商id
	_distributorTid int64
	// 采购单号
	_purchaseOrderId int64
}

// NewTaobaoAlitripTravelAxinHotelOrderPayRequest 初始化TaobaoAlitripTravelAxinHotelOrderPayAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelOrderPayRequest() *TaobaoAlitripTravelAxinHotelOrderPayAPIRequest {
	return &TaobaoAlitripTravelAxinHotelOrderPayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelOrderPayAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.order.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelOrderPayAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoAlitripTravelAxinHotelOrderPayAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelAxinHotelOrderPayAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetPurchaseOrderId is PurchaseOrderId Setter
// 采购单号
func (r *TaobaoAlitripTravelAxinHotelOrderPayAPIRequest) SetPurchaseOrderId(_purchaseOrderId int64) error {
	r._purchaseOrderId = _purchaseOrderId
	r.Set("purchase_order_id", _purchaseOrderId)
	return nil
}

// GetPurchaseOrderId PurchaseOrderId Getter
func (r TaobaoAlitripTravelAxinHotelOrderPayAPIRequest) GetPurchaseOrderId() int64 {
	return r._purchaseOrderId
}
