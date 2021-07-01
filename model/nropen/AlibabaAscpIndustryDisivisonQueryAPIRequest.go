package nropen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAscpIndustryDisivisonQueryAPIRequest
查询服务支持地区列表 API请求
alibaba.ascp.industry.disivison.query

商家获取服务支持地区 */
type AlibabaAscpIndustryDisivisonQueryAPIRequest struct {
	model.Params
	// 服务编码
	_serviceCode string
}

// NewAlibabaAscpIndustryDisivisonQueryRequest 初始化AlibabaAscpIndustryDisivisonQueryAPIRequest对象
func NewAlibabaAscpIndustryDisivisonQueryRequest() *AlibabaAscpIndustryDisivisonQueryAPIRequest {
	return &AlibabaAscpIndustryDisivisonQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpIndustryDisivisonQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.disivison.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpIndustryDisivisonQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ServiceCode Setter
// 服务编码
func (r *AlibabaAscpIndustryDisivisonQueryAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// Get ServiceCode Getter
func (r AlibabaAscpIndustryDisivisonQueryAPIRequest) GetServiceCode() string {
	return r._serviceCode
}
