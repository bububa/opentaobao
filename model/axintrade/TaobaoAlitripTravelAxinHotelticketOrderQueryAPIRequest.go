package axintrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest 阿信度假交易订单查询接口 API请求
// taobao.alitrip.travel.axin.hotelticket.order.query
//
// 阿信度假交易订单查询接口
type TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest struct {
	model.Params
	// 分销商id
	_distributorTid int64
	// 采购单id
	_purchaseSubOrderId int64
}

// NewTaobaoAlitripTravelAxinHotelticketOrderQueryRequest 初始化TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelticketOrderQueryRequest() *TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest {
	return &TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest) Reset() {
	r._distributorTid = 0
	r._purchaseSubOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotelticket.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetPurchaseSubOrderId is PurchaseSubOrderId Setter
// 采购单id
func (r *TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest) SetPurchaseSubOrderId(_purchaseSubOrderId int64) error {
	r._purchaseSubOrderId = _purchaseSubOrderId
	r.Set("purchase_sub_order_id", _purchaseSubOrderId)
	return nil
}

// GetPurchaseSubOrderId PurchaseSubOrderId Getter
func (r TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest) GetPurchaseSubOrderId() int64 {
	return r._purchaseSubOrderId
}

var poolTaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelAxinHotelticketOrderQueryRequest()
	},
}

// GetTaobaoAlitripTravelAxinHotelticketOrderQueryRequest 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest
func GetTaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest() *TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest {
	return poolTaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest.Get().(*TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest)
}

// ReleaseTaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest 将 TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest(v *TaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelticketOrderQueryAPIRequest.Put(v)
}
