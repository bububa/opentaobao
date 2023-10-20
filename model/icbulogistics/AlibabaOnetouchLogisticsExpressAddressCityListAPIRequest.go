package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaonetouchlogisticsexpressaddresscitylistAPIRequest 四级地址库-市 API请求
// alibaba.onetouch.logistics.express.address.city.list
//
// 四级地址库-市
type AlibabaonetouchlogisticsexpressaddresscitylistAPIRequest struct {
	model.Params
	// 请求参数
	_paramQuery *AddressQueryDto
}

// NewAlibabaonetouchlogisticsexpressaddresscitylistRequest 初始化AlibabaonetouchlogisticsexpressaddresscitylistAPIRequest对象
func NewAlibabaonetouchlogisticsexpressaddresscitylistRequest() *AlibabaonetouchlogisticsexpressaddresscitylistAPIRequest {
	return &AlibabaonetouchlogisticsexpressaddresscitylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaonetouchlogisticsexpressaddresscitylistAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.address.city.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaonetouchlogisticsexpressaddresscitylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaonetouchlogisticsexpressaddresscitylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQuery is ParamQuery Setter
// 请求参数
func (r *AlibabaonetouchlogisticsexpressaddresscitylistAPIRequest) SetParamQuery(_paramQuery *AddressQueryDto) error {
	r._paramQuery = _paramQuery
	r.Set("param_query", _paramQuery)
	return nil
}

// GetParamQuery ParamQuery Getter
func (r AlibabaonetouchlogisticsexpressaddresscitylistAPIRequest) GetParamQuery() *AddressQueryDto {
	return r._paramQuery
}
