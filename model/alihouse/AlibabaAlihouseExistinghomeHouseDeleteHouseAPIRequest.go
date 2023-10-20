package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomehousedeletehouseAPIRequest 删除房源 API请求
// alibaba.alihouse.existinghome.house.delete.house
//
// 删除房源
type AlibabaalihouseexistinghomehousedeletehouseAPIRequest struct {
	model.Params
	// 房源信息
	_house *DeleteHouseDto
}

// NewAlibabaalihouseexistinghomehousedeletehouseRequest 初始化AlibabaalihouseexistinghomehousedeletehouseAPIRequest对象
func NewAlibabaalihouseexistinghomehousedeletehouseRequest() *AlibabaalihouseexistinghomehousedeletehouseAPIRequest {
	return &AlibabaalihouseexistinghomehousedeletehouseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomehousedeletehouseAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.house.delete.house"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomehousedeletehouseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomehousedeletehouseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHouse is House Setter
// 房源信息
func (r *AlibabaalihouseexistinghomehousedeletehouseAPIRequest) SetHouse(_house *DeleteHouseDto) error {
	r._house = _house
	r.Set("house", _house)
	return nil
}

// GetHouse House Getter
func (r AlibabaalihouseexistinghomehousedeletehouseAPIRequest) GetHouse() *DeleteHouseDto {
	return r._house
}
