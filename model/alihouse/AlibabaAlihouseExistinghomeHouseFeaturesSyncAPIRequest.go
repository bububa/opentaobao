package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomehousefeaturessyncAPIRequest 房源标准打标数据同步 API请求
// alibaba.alihouse.existinghome.house.features.sync
//
// 房源标准打标数据同步
type AlibabaalihouseexistinghomehousefeaturessyncAPIRequest struct {
	model.Params
	// 房源上翻标
	_houseFeatures *SyncHouseFeaturesDto
}

// NewAlibabaalihouseexistinghomehousefeaturessyncRequest 初始化AlibabaalihouseexistinghomehousefeaturessyncAPIRequest对象
func NewAlibabaalihouseexistinghomehousefeaturessyncRequest() *AlibabaalihouseexistinghomehousefeaturessyncAPIRequest {
	return &AlibabaalihouseexistinghomehousefeaturessyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomehousefeaturessyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.house.features.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomehousefeaturessyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomehousefeaturessyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHouseFeatures is HouseFeatures Setter
// 房源上翻标
func (r *AlibabaalihouseexistinghomehousefeaturessyncAPIRequest) SetHouseFeatures(_houseFeatures *SyncHouseFeaturesDto) error {
	r._houseFeatures = _houseFeatures
	r.Set("house_features", _houseFeatures)
	return nil
}

// GetHouseFeatures HouseFeatures Getter
func (r AlibabaalihouseexistinghomehousefeaturessyncAPIRequest) GetHouseFeatures() *SyncHouseFeaturesDto {
	return r._houseFeatures
}
