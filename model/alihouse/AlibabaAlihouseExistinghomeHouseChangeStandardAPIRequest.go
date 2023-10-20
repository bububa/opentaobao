package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomehousechangestandardAPIRequest 委托房源变更标准房源 API请求
// alibaba.alihouse.existinghome.house.change.standard
//
// 委托房源变更标准房源
type AlibabaalihouseexistinghomehousechangestandardAPIRequest struct {
	model.Params
	// 房源信息
	_house *EntrustChangeStandardDto
}

// NewAlibabaalihouseexistinghomehousechangestandardRequest 初始化AlibabaalihouseexistinghomehousechangestandardAPIRequest对象
func NewAlibabaalihouseexistinghomehousechangestandardRequest() *AlibabaalihouseexistinghomehousechangestandardAPIRequest {
	return &AlibabaalihouseexistinghomehousechangestandardAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihouseexistinghomehousechangestandardAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.existinghome.house.change.standard"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihouseexistinghomehousechangestandardAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihouseexistinghomehousechangestandardAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHouse is House Setter
// 房源信息
func (r *AlibabaalihouseexistinghomehousechangestandardAPIRequest) SetHouse(_house *EntrustChangeStandardDto) error {
	r._house = _house
	r.Set("house", _house)
	return nil
}

// GetHouse House Getter
func (r AlibabaalihouseexistinghomehousechangestandardAPIRequest) GetHouse() *EntrustChangeStandardDto {
	return r._house
}
