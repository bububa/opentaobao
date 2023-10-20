package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghousehousebaseAPIRequest （租房）同步房屋信息 API请求
// alibaba.alihouse.existinghouse.house.base
//
// 房屋信息上翻
type AlibabaalihouseexistinghousehousebaseAPIRequest struct {
	model.Params
	// 基础房屋信息
	_houseBase *HouseBaseDto
}

// NewAlibabaalihouseexistinghousehousebaseRequest 初始化AlibabaalihouseexistinghousehousebaseAPIRequest对象
func NewAlibabaalihouseexistinghousehousebaseRequest() *AlibabaalihouseexistinghousehousebaseAPIRequest {
	return &AlibabaalihouseexistinghousehousebaseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghousehousebaseAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghouse.house.base"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghousehousebaseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghousehousebaseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHouseBase is HouseBase Setter
// 基础房屋信息
func (r *AlibabaalihouseexistinghousehousebaseAPIRequest) SetHouseBase(_houseBase *HouseBaseDto) error {
	r._houseBase = _houseBase
	r.Set("house_base", _houseBase)
	return nil
}

// GetHouseBase HouseBase Getter
func (r AlibabaalihouseexistinghousehousebaseAPIRequest) GetHouseBase() *HouseBaseDto {
	return r._houseBase
}
