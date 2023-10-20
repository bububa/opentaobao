package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelroomlistqueryAPIRequest 阿信酒店分销-标准酒店房型列表查询 API请求
// taobao.alitrip.travel.axin.hotel.room.list.query
//
// 标准酒店房型列表查询
type TaobaoalitriptravelaxinhotelroomlistqueryAPIRequest struct {
	model.Params
	// 资源渠道
	_resourceChannel string
	// 标准酒店id
	_shid int64
	// 分销商id
	_distributorTid int64
}

// NewTaobaoalitriptravelaxinhotelroomlistqueryRequest 初始化TaobaoalitriptravelaxinhotelroomlistqueryAPIRequest对象
func NewTaobaoalitriptravelaxinhotelroomlistqueryRequest() *TaobaoalitriptravelaxinhotelroomlistqueryAPIRequest {
	return &TaobaoalitriptravelaxinhotelroomlistqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelaxinhotelroomlistqueryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.room.list.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelaxinhotelroomlistqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelaxinhotelroomlistqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResourceChannel is ResourceChannel Setter
// 资源渠道
func (r *TaobaoalitriptravelaxinhotelroomlistqueryAPIRequest) SetResourceChannel(_resourceChannel string) error {
	r._resourceChannel = _resourceChannel
	r.Set("resource_channel", _resourceChannel)
	return nil
}

// GetResourceChannel ResourceChannel Getter
func (r TaobaoalitriptravelaxinhotelroomlistqueryAPIRequest) GetResourceChannel() string {
	return r._resourceChannel
}

// SetShid is Shid Setter
// 标准酒店id
func (r *TaobaoalitriptravelaxinhotelroomlistqueryAPIRequest) SetShid(_shid int64) error {
	r._shid = _shid
	r.Set("shid", _shid)
	return nil
}

// GetShid Shid Getter
func (r TaobaoalitriptravelaxinhotelroomlistqueryAPIRequest) GetShid() int64 {
	return r._shid
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoalitriptravelaxinhotelroomlistqueryAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoalitriptravelaxinhotelroomlistqueryAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}
