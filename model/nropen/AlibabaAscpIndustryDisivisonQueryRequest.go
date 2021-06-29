package nropen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询服务支持地区列表 API请求
alibaba.ascp.industry.disivison.query

商家获取服务支持地区
*/
type AlibabaAscpIndustryDisivisonQueryRequest struct {
    model.Params
    // 服务编码
    _serviceCode   string
}

// 初始化AlibabaAscpIndustryDisivisonQueryRequest对象
func NewAlibabaAscpIndustryDisivisonQueryRequest() *AlibabaAscpIndustryDisivisonQueryRequest{
    return &AlibabaAscpIndustryDisivisonQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpIndustryDisivisonQueryRequest) GetApiMethodName() string {
    return "alibaba.ascp.industry.disivison.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpIndustryDisivisonQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServiceCode Setter
// 服务编码
func (r *AlibabaAscpIndustryDisivisonQueryRequest) SetServiceCode(_serviceCode string) error {
    r._serviceCode = _serviceCode
    r.Set("service_code", _serviceCode)
    return nil
}

// ServiceCode Getter
func (r AlibabaAscpIndustryDisivisonQueryRequest) GetServiceCode() string {
    return r._serviceCode
}
