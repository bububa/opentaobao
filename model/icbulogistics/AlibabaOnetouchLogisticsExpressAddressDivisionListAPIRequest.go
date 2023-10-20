package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaonetouchlogisticsexpressaddressdivisionlistAPIRequest 四级地址库-区域 API请求
// alibaba.onetouch.logistics.express.address.division.list
//
// 四级地址库-区
type AlibabaonetouchlogisticsexpressaddressdivisionlistAPIRequest struct {
	model.Params
	// 请求参数
	_paramQuery *AddressQueryDto
}

// NewAlibabaonetouchlogisticsexpressaddressdivisionlistRequest 初始化AlibabaonetouchlogisticsexpressaddressdivisionlistAPIRequest对象
func NewAlibabaonetouchlogisticsexpressaddressdivisionlistRequest() *AlibabaonetouchlogisticsexpressaddressdivisionlistAPIRequest {
	return &AlibabaonetouchlogisticsexpressaddressdivisionlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaonetouchlogisticsexpressaddressdivisionlistAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.address.division.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaonetouchlogisticsexpressaddressdivisionlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaonetouchlogisticsexpressaddressdivisionlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQuery is ParamQuery Setter
// 请求参数
func (r *AlibabaonetouchlogisticsexpressaddressdivisionlistAPIRequest) SetParamQuery(_paramQuery *AddressQueryDto) error {
	r._paramQuery = _paramQuery
	r.Set("param_query", _paramQuery)
	return nil
}

// GetParamQuery ParamQuery Getter
func (r AlibabaonetouchlogisticsexpressaddressdivisionlistAPIRequest) GetParamQuery() *AddressQueryDto {
	return r._paramQuery
}
