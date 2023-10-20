package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeregionsyncAPIRequest 城区数据同步 API请求
// alibaba.alihouse.newhome.region.sync
//
// 城区数据同步
type AlibabaalihousenewhomeregionsyncAPIRequest struct {
	model.Params
	// 城区数据
	_baseRegionDto *BaseRegionDto
}

// NewAlibabaalihousenewhomeregionsyncRequest 初始化AlibabaalihousenewhomeregionsyncAPIRequest对象
func NewAlibabaalihousenewhomeregionsyncRequest() *AlibabaalihousenewhomeregionsyncAPIRequest {
	return &AlibabaalihousenewhomeregionsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeregionsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.region.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeregionsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeregionsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBaseRegionDto is BaseRegionDto Setter
// 城区数据
func (r *AlibabaalihousenewhomeregionsyncAPIRequest) SetBaseRegionDto(_baseRegionDto *BaseRegionDto) error {
	r._baseRegionDto = _baseRegionDto
	r.Set("base_region_dto", _baseRegionDto)
	return nil
}

// GetBaseRegionDto BaseRegionDto Getter
func (r AlibabaalihousenewhomeregionsyncAPIRequest) GetBaseRegionDto() *BaseRegionDto {
	return r._baseRegionDto
}
