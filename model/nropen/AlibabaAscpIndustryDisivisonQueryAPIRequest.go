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
type AlibabaAscpIndustryDisivisonQueryAPIRequest struct {
    model.Params
    // 服务编码
    _serviceCode   string
}

// 初始化AlibabaAscpIndustryDisivisonQueryAPIRequest对象
func NewAlibabaAscpIndustryDisivisonQueryRequest() *AlibabaAscpIndustryDisivisonQueryAPIRequest{
    return &AlibabaAscpIndustryDisivisonQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAscpIndustryDisivisonQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.ascp.industry.disivison.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAscpIndustryDisivisonQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ServiceCode Setter
// 服务编码
func (r *AlibabaAscpIndustryDisivisonQueryAPIRequest) SetServiceCode(_serviceCode string) error {
    r._serviceCode = _serviceCode
    r.Set("service_code", _serviceCode)
    return nil
}

// ServiceCode Getter
func (r AlibabaAscpIndustryDisivisonQueryAPIRequest) GetServiceCode() string {
    return r._serviceCode
}
