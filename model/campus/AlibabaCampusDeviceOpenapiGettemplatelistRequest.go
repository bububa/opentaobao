package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询设备模板 APIRequest
alibaba.campus.device.openapi.gettemplatelist

查询设备模板信息
*/
type AlibabaCampusDeviceOpenapiGettemplatelistRequest struct {
    model.Params

    // 设备模板查询对象
    query   *TemplateApiQuery 

}

func NewAlibabaCampusDeviceOpenapiGettemplatelistRequest() *AlibabaCampusDeviceOpenapiGettemplatelistRequest{
    return &AlibabaCampusDeviceOpenapiGettemplatelistRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCampusDeviceOpenapiGettemplatelistRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.gettemplatelist"
}

func (r AlibabaCampusDeviceOpenapiGettemplatelistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCampusDeviceOpenapiGettemplatelistRequest) SetQuery(query *TemplateApiQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaCampusDeviceOpenapiGettemplatelistRequest) GetQuery() *TemplateApiQuery {
    return r.query
}

