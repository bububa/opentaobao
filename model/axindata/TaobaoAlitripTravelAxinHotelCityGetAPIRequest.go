package axindata

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelCityGetAPIRequest 城市列表信息查询-阿信 API请求
// taobao.alitrip.travel.axin.hotel.city.get
//
// 阿信城市列表信息查询
type TaobaoAlitripTravelAxinHotelCityGetAPIRequest struct {
	model.Params
	// 分销商id，阿信分配
	_distributorTid int64
}

// NewTaobaoAlitripTravelAxinHotelCityGetRequest 初始化TaobaoAlitripTravelAxinHotelCityGetAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelCityGetRequest() *TaobaoAlitripTravelAxinHotelCityGetAPIRequest {
	return &TaobaoAlitripTravelAxinHotelCityGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelAxinHotelCityGetAPIRequest) Reset() {
	r._distributorTid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelCityGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.city.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelCityGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelAxinHotelCityGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributorTid is DistributorTid Setter
// 分销商id，阿信分配
func (r *TaobaoAlitripTravelAxinHotelCityGetAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelAxinHotelCityGetAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

var poolTaobaoAlitripTravelAxinHotelCityGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelAxinHotelCityGetRequest()
	},
}

// GetTaobaoAlitripTravelAxinHotelCityGetRequest 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelCityGetAPIRequest
func GetTaobaoAlitripTravelAxinHotelCityGetAPIRequest() *TaobaoAlitripTravelAxinHotelCityGetAPIRequest {
	return poolTaobaoAlitripTravelAxinHotelCityGetAPIRequest.Get().(*TaobaoAlitripTravelAxinHotelCityGetAPIRequest)
}

// ReleaseTaobaoAlitripTravelAxinHotelCityGetAPIRequest 将 TaobaoAlitripTravelAxinHotelCityGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelCityGetAPIRequest(v *TaobaoAlitripTravelAxinHotelCityGetAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelCityGetAPIRequest.Put(v)
}
