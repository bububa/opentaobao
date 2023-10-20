package nropen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryDisivisonQueryAPIRequest 查询服务支持地区列表 API请求
// alibaba.ascp.industry.disivison.query
//
// 商家获取服务支持地区
type AlibabaAscpIndustryDisivisonQueryAPIRequest struct {
	model.Params
	// 服务编码
	_serviceCode string
}

// NewAlibabaAscpIndustryDisivisonQueryRequest 初始化AlibabaAscpIndustryDisivisonQueryAPIRequest对象
func NewAlibabaAscpIndustryDisivisonQueryRequest() *AlibabaAscpIndustryDisivisonQueryAPIRequest {
	return &AlibabaAscpIndustryDisivisonQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpIndustryDisivisonQueryAPIRequest) Reset() {
	r._serviceCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpIndustryDisivisonQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.industry.disivison.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpIndustryDisivisonQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpIndustryDisivisonQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceCode is ServiceCode Setter
// 服务编码
func (r *AlibabaAscpIndustryDisivisonQueryAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// GetServiceCode ServiceCode Getter
func (r AlibabaAscpIndustryDisivisonQueryAPIRequest) GetServiceCode() string {
	return r._serviceCode
}

var poolAlibabaAscpIndustryDisivisonQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpIndustryDisivisonQueryRequest()
	},
}

// GetAlibabaAscpIndustryDisivisonQueryRequest 从 sync.Pool 获取 AlibabaAscpIndustryDisivisonQueryAPIRequest
func GetAlibabaAscpIndustryDisivisonQueryAPIRequest() *AlibabaAscpIndustryDisivisonQueryAPIRequest {
	return poolAlibabaAscpIndustryDisivisonQueryAPIRequest.Get().(*AlibabaAscpIndustryDisivisonQueryAPIRequest)
}

// ReleaseAlibabaAscpIndustryDisivisonQueryAPIRequest 将 AlibabaAscpIndustryDisivisonQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpIndustryDisivisonQueryAPIRequest(v *AlibabaAscpIndustryDisivisonQueryAPIRequest) {
	v.Reset()
	poolAlibabaAscpIndustryDisivisonQueryAPIRequest.Put(v)
}
