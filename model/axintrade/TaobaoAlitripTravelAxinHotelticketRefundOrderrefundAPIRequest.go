package axintrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest 阿信度假业务申请退款 API请求
// taobao.alitrip.travel.axin.hotelticket.refund.orderrefund
//
// 阿信度假业务申请退款
type TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest struct {
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

// NewTaobaoAlitripTravelAxinHotelticketRefundOrderrefundRequest 初始化TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelticketRefundOrderrefundRequest() *TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest {
	return &TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest) Reset() {
	r._remark = ""
	r._reason = ""
	r._distributorTid = 0
	r._purchaseSubOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotelticket.refund.orderrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRemark is Remark Setter
// 备注
func (r *TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest) GetRemark() string {
	return r._remark
}

// SetReason is Reason Setter
// 退款原因
func (r *TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest) GetReason() string {
	return r._reason
}

// SetDistributorTid is DistributorTid Setter
// 分销商ID
func (r *TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetPurchaseSubOrderId is PurchaseSubOrderId Setter
// 采购单号
func (r *TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest) SetPurchaseSubOrderId(_purchaseSubOrderId int64) error {
	r._purchaseSubOrderId = _purchaseSubOrderId
	r.Set("purchase_sub_order_id", _purchaseSubOrderId)
	return nil
}

// GetPurchaseSubOrderId PurchaseSubOrderId Getter
func (r TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest) GetPurchaseSubOrderId() int64 {
	return r._purchaseSubOrderId
}

var poolTaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelAxinHotelticketRefundOrderrefundRequest()
	},
}

// GetTaobaoAlitripTravelAxinHotelticketRefundOrderrefundRequest 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest
func GetTaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest() *TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest {
	return poolTaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest.Get().(*TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest)
}

// ReleaseTaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest 将 TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest(v *TaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelticketRefundOrderrefundAPIRequest.Put(v)
}
