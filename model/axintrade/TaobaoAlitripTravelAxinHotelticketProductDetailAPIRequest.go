package axintrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest 阿信酒景套餐产品详情查询 API请求
// taobao.alitrip.travel.axin.hotelticket.product.detail
//
// 阿信酒景套餐产品详情查询
type TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest struct {
	model.Params
	// 分销商id
	_distributorTid int64
	// 产品id
	_productId int64
}

// NewTaobaoAlitripTravelAxinHotelticketProductDetailRequest 初始化TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelticketProductDetailRequest() *TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest {
	return &TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest) Reset() {
	r._distributorTid = 0
	r._productId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotelticket.product.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetProductId is ProductId Setter
// 产品id
func (r *TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest) GetProductId() int64 {
	return r._productId
}

var poolTaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelAxinHotelticketProductDetailRequest()
	},
}

// GetTaobaoAlitripTravelAxinHotelticketProductDetailRequest 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest
func GetTaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest() *TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest {
	return poolTaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest.Get().(*TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest)
}

// ReleaseTaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest 将 TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest(v *TaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelticketProductDetailAPIRequest.Put(v)
}
