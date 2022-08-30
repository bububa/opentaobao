package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelRoomListQueryAPIRequest 阿信酒店分销-标准酒店房型列表查询 API请求
// taobao.alitrip.travel.axin.hotel.room.list.query
//
// 标准酒店房型列表查询
type TaobaoAlitripTravelAxinHotelRoomListQueryAPIRequest struct {
	model.Params
	// 资源渠道
	_resourceChannel string
	// 标准酒店id
	_shid int64
	// 分销商id
	_distributorTid int64
}

// NewTaobaoAlitripTravelAxinHotelRoomListQueryRequest 初始化TaobaoAlitripTravelAxinHotelRoomListQueryAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelRoomListQueryRequest() *TaobaoAlitripTravelAxinHotelRoomListQueryAPIRequest {
	return &TaobaoAlitripTravelAxinHotelRoomListQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelRoomListQueryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.room.list.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelRoomListQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetResourceChannel is ResourceChannel Setter
// 资源渠道
func (r *TaobaoAlitripTravelAxinHotelRoomListQueryAPIRequest) SetResourceChannel(_resourceChannel string) error {
	r._resourceChannel = _resourceChannel
	r.Set("resource_channel", _resourceChannel)
	return nil
}

// GetResourceChannel ResourceChannel Getter
func (r TaobaoAlitripTravelAxinHotelRoomListQueryAPIRequest) GetResourceChannel() string {
	return r._resourceChannel
}

// SetShid is Shid Setter
// 标准酒店id
func (r *TaobaoAlitripTravelAxinHotelRoomListQueryAPIRequest) SetShid(_shid int64) error {
	r._shid = _shid
	r.Set("shid", _shid)
	return nil
}

// GetShid Shid Getter
func (r TaobaoAlitripTravelAxinHotelRoomListQueryAPIRequest) GetShid() int64 {
	return r._shid
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoAlitripTravelAxinHotelRoomListQueryAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelAxinHotelRoomListQueryAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}
