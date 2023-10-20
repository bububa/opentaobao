package icbulogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaonetouchlogisticsexpressaddressstreetlistAPIRequest 四级地址库-街道 API请求
// alibaba.onetouch.logistics.express.address.street.list
//
// 四级地址库-街道模糊查询
type AlibabaonetouchlogisticsexpressaddressstreetlistAPIRequest struct {
	model.Params
	// 请求参数
	_paramQuery *AddressQueryDto
}

// NewAlibabaonetouchlogisticsexpressaddressstreetlistRequest 初始化AlibabaonetouchlogisticsexpressaddressstreetlistAPIRequest对象
func NewAlibabaonetouchlogisticsexpressaddressstreetlistRequest() *AlibabaonetouchlogisticsexpressaddressstreetlistAPIRequest {
	return &AlibabaonetouchlogisticsexpressaddressstreetlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaonetouchlogisticsexpressaddressstreetlistAPIRequest) GetApiMethodName() string {
	return "alibaba.onetouch.logistics.express.address.street.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaonetouchlogisticsexpressaddressstreetlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaonetouchlogisticsexpressaddressstreetlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQuery is ParamQuery Setter
// 请求参数
func (r *AlibabaonetouchlogisticsexpressaddressstreetlistAPIRequest) SetParamQuery(_paramQuery *AddressQueryDto) error {
	r._paramQuery = _paramQuery
	r.Set("param_query", _paramQuery)
	return nil
}

// GetParamQuery ParamQuery Getter
func (r AlibabaonetouchlogisticsexpressaddressstreetlistAPIRequest) GetParamQuery() *AddressQueryDto {
	return r._paramQuery
}
