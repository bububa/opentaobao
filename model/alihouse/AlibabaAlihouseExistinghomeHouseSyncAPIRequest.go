package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomehousesyncAPIRequest 房源信息同步 API请求
// alibaba.alihouse.existinghome.house.sync
//
// 房源信息同步
type AlibabaalihouseexistinghomehousesyncAPIRequest struct {
	model.Params
	// 房源信息
	_house *SyncExistingHouseDto
}

// NewAlibabaalihouseexistinghomehousesyncRequest 初始化AlibabaalihouseexistinghomehousesyncAPIRequest对象
func NewAlibabaalihouseexistinghomehousesyncRequest() *AlibabaalihouseexistinghomehousesyncAPIRequest {
	return &AlibabaalihouseexistinghomehousesyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomehousesyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.house.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomehousesyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomehousesyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHouse is House Setter
// 房源信息
func (r *AlibabaalihouseexistinghomehousesyncAPIRequest) SetHouse(_house *SyncExistingHouseDto) error {
	r._house = _house
	r.Set("house", _house)
	return nil
}

// GetHouse House Getter
func (r AlibabaalihouseexistinghomehousesyncAPIRequest) GetHouse() *SyncExistingHouseDto {
	return r._house
}
