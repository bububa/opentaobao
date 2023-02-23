package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelDetailQueryAPIRequest 阿信酒店分销-标准酒店详情查询 API请求
// taobao.alitrip.travel.axin.hotel.detail.query
//
// 标准酒店详情查询
type TaobaoAlitripTravelAxinHotelDetailQueryAPIRequest struct {
	model.Params
	// 资源渠道
	_resourceChannel string
	// 标准酒店id
	_shid int64
	// 分销商id
	_distributorTid int64
}

// NewTaobaoAlitripTravelAxinHotelDetailQueryRequest 初始化TaobaoAlitripTravelAxinHotelDetailQueryAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelDetailQueryRequest() *TaobaoAlitripTravelAxinHotelDetailQueryAPIRequest {
	return &TaobaoAlitripTravelAxinHotelDetailQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelDetailQueryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelDetailQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelAxinHotelDetailQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResourceChannel is ResourceChannel Setter
// 资源渠道
func (r *TaobaoAlitripTravelAxinHotelDetailQueryAPIRequest) SetResourceChannel(_resourceChannel string) error {
	r._resourceChannel = _resourceChannel
	r.Set("resource_channel", _resourceChannel)
	return nil
}

// GetResourceChannel ResourceChannel Getter
func (r TaobaoAlitripTravelAxinHotelDetailQueryAPIRequest) GetResourceChannel() string {
	return r._resourceChannel
}

// SetShid is Shid Setter
// 标准酒店id
func (r *TaobaoAlitripTravelAxinHotelDetailQueryAPIRequest) SetShid(_shid int64) error {
	r._shid = _shid
	r.Set("shid", _shid)
	return nil
}

// GetShid Shid Getter
func (r TaobaoAlitripTravelAxinHotelDetailQueryAPIRequest) GetShid() int64 {
	return r._shid
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoAlitripTravelAxinHotelDetailQueryAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelAxinHotelDetailQueryAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}
