package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest 阿信酒店批量报价查询接口 API请求
// taobao.alitrip.travel.axin.hotel.price.batch.get
//
// 阿信酒店批量报价查询接口
type TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest struct {
	model.Params
	// 结束时间，不包含
	_endDate string
	// 开始时间，包含
	_startDate string
	// 资源渠道
	_resourceChannel string
	// 酒店列表
	_hotelList *HotelDto
	// 分销商id
	_distributorTid int64
}

// NewTaobaoAlitripTravelAxinHotelPriceBatchGetRequest 初始化TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelPriceBatchGetRequest() *TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest {
	return &TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.price.batch.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetEndDate is EndDate Setter
// 结束时间，不包含
func (r *TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetStartDate is StartDate Setter
// 开始时间，包含
func (r *TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetResourceChannel is ResourceChannel Setter
// 资源渠道
func (r *TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest) SetResourceChannel(_resourceChannel string) error {
	r._resourceChannel = _resourceChannel
	r.Set("resource_channel", _resourceChannel)
	return nil
}

// GetResourceChannel ResourceChannel Getter
func (r TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest) GetResourceChannel() string {
	return r._resourceChannel
}

// SetHotelList is HotelList Setter
// 酒店列表
func (r *TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest) SetHotelList(_hotelList *HotelDto) error {
	r._hotelList = _hotelList
	r.Set("hotel_list", _hotelList)
	return nil
}

// GetHotelList HotelList Getter
func (r TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest) GetHotelList() *HotelDto {
	return r._hotelList
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelAxinHotelPriceBatchGetAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}
