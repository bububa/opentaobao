package nropen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询服务支持地区列表 APIRequest
alibaba.ascp.industry.disivison.query

商家获取服务支持地区
*/
type AlibabaAscpIndustryDisivisonQueryRequest struct {
    model.Params

    // 服务编码
    serviceCode   string 

}

func NewAlibabaAscpIndustryDisivisonQueryRequest() *AlibabaAscpIndustryDisivisonQueryRequest{
    return &AlibabaAscpIndustryDisivisonQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAscpIndustryDisivisonQueryRequest) GetApiMethodName() string {
    return "alibaba.ascp.industry.disivison.query"
}

func (r AlibabaAscpIndustryDisivisonQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAscpIndustryDisivisonQueryRequest) SetServiceCode(serviceCode string) error {
    r.serviceCode = serviceCode
    r.Set("service_code", serviceCode)
    return nil
}

func (r AlibabaAscpIndustryDisivisonQueryRequest) GetServiceCode() string {
    return r.serviceCode
}

