package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaonetouchlogisticsexpressaddressprovincelistAPIRequest 四级地址库-省 API请求
// alibaba.onetouch.logistics.express.address.province.list
//
// 四级地址库-省
type AlibabaonetouchlogisticsexpressaddressprovincelistAPIRequest struct {
	model.Params
	// 请求参数
	_paramQuery *AddressQueryDto
}

// NewAlibabaonetouchlogisticsexpressaddressprovincelistRequest 初始化AlibabaonetouchlogisticsexpressaddressprovincelistAPIRequest对象
func NewAlibabaonetouchlogisticsexpressaddressprovincelistRequest() *AlibabaonetouchlogisticsexpressaddressprovincelistAPIRequest {
	return &AlibabaonetouchlogisticsexpressaddressprovincelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaonetouchlogisticsexpressaddressprovincelistAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.address.province.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaonetouchlogisticsexpressaddressprovincelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaonetouchlogisticsexpressaddressprovincelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQuery is ParamQuery Setter
// 请求参数
func (r *AlibabaonetouchlogisticsexpressaddressprovincelistAPIRequest) SetParamQuery(_paramQuery *AddressQueryDto) error {
	r._paramQuery = _paramQuery
	r.Set("param_query", _paramQuery)
	return nil
}

// GetParamQuery ParamQuery Getter
func (r AlibabaonetouchlogisticsexpressaddressprovincelistAPIRequest) GetParamQuery() *AddressQueryDto {
	return r._paramQuery
}
