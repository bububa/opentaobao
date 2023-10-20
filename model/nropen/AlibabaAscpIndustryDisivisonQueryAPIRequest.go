package nropen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpindustrydisivisonqueryAPIRequest 查询服务支持地区列表 API请求
// alibaba.ascp.industry.disivison.query
//
// 商家获取服务支持地区
type AlibabaascpindustrydisivisonqueryAPIRequest struct {
	model.Params
	// 服务编码
	_serviceCode string
}

// NewAlibabaascpindustrydisivisonqueryRequest 初始化AlibabaascpindustrydisivisonqueryAPIRequest对象
func NewAlibabaascpindustrydisivisonqueryRequest() *AlibabaascpindustrydisivisonqueryAPIRequest {
	return &AlibabaascpindustrydisivisonqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaascpindustrydisivisonqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.disivison.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaascpindustrydisivisonqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaascpindustrydisivisonqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceCode is ServiceCode Setter
// 服务编码
func (r *AlibabaascpindustrydisivisonqueryAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// GetServiceCode ServiceCode Getter
func (r AlibabaascpindustrydisivisonqueryAPIRequest) GetServiceCode() string {
	return r._serviceCode
}
