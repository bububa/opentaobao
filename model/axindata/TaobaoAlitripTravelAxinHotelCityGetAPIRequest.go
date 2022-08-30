package axindata

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelCityGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.city.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelCityGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
