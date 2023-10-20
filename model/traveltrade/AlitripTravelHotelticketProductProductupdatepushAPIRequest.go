package traveltrade

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelHotelticketProductProductupdatepushAPIRequest 产品批量变更推送通知 API请求
// alitrip.travel.hotelticket.product.productupdatepush
//
// 产品批量变更推送通知
type AlitripTravelHotelticketProductProductupdatepushAPIRequest struct {
	model.Params
	// 系统商分配给飞猪卖家的账号
	_accessKey string
	// 变更的产品信息
	_productUpdates *ProductUpdatePushDto
}

// NewAlitripTravelHotelticketProductProductupdatepushRequest 初始化AlitripTravelHotelticketProductProductupdatepushAPIRequest对象
func NewAlitripTravelHotelticketProductProductupdatepushRequest() *AlitripTravelHotelticketProductProductupdatepushAPIRequest {
	return &AlitripTravelHotelticketProductProductupdatepushAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTravelHotelticketProductProductupdatepushAPIRequest) Reset() {
	r._accessKey = ""
	r._productUpdates = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTravelHotelticketProductProductupdatepushAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.hotelticket.product.productupdatepush"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTravelHotelticketProductProductupdatepushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTravelHotelticketProductProductupdatepushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccessKey is AccessKey Setter
// 系统商分配给飞猪卖家的账号
func (r *AlitripTravelHotelticketProductProductupdatepushAPIRequest) SetAccessKey(_accessKey string) error {
	r._accessKey = _accessKey
	r.Set("access_key", _accessKey)
	return nil
}

// GetAccessKey AccessKey Getter
func (r AlitripTravelHotelticketProductProductupdatepushAPIRequest) GetAccessKey() string {
	return r._accessKey
}

// SetProductUpdates is ProductUpdates Setter
// 变更的产品信息
func (r *AlitripTravelHotelticketProductProductupdatepushAPIRequest) SetProductUpdates(_productUpdates *ProductUpdatePushDto) error {
	r._productUpdates = _productUpdates
	r.Set("product_updates", _productUpdates)
	return nil
}

// GetProductUpdates ProductUpdates Getter
func (r AlitripTravelHotelticketProductProductupdatepushAPIRequest) GetProductUpdates() *ProductUpdatePushDto {
	return r._productUpdates
}

var poolAlitripTravelHotelticketProductProductupdatepushAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTravelHotelticketProductProductupdatepushRequest()
	},
}

// GetAlitripTravelHotelticketProductProductupdatepushRequest 从 sync.Pool 获取 AlitripTravelHotelticketProductProductupdatepushAPIRequest
func GetAlitripTravelHotelticketProductProductupdatepushAPIRequest() *AlitripTravelHotelticketProductProductupdatepushAPIRequest {
	return poolAlitripTravelHotelticketProductProductupdatepushAPIRequest.Get().(*AlitripTravelHotelticketProductProductupdatepushAPIRequest)
}

// ReleaseAlitripTravelHotelticketProductProductupdatepushAPIRequest 将 AlitripTravelHotelticketProductProductupdatepushAPIRequest 放入 sync.Pool
func ReleaseAlitripTravelHotelticketProductProductupdatepushAPIRequest(v *AlitripTravelHotelticketProductProductupdatepushAPIRequest) {
	v.Reset()
	poolAlitripTravelHotelticketProductProductupdatepushAPIRequest.Put(v)
}
