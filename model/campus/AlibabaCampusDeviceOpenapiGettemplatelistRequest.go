package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询设备模板 API请求
alibaba.campus.device.openapi.gettemplatelist

查询设备模板信息
*/
type AlibabaCampusDeviceOpenapiGettemplatelistRequest struct {
    model.Params
    // 设备模板查询对象
    query   *TemplateApiQuery
}

// 初始化AlibabaCampusDeviceOpenapiGettemplatelistRequest对象
func NewAlibabaCampusDeviceOpenapiGettemplatelistRequest() *AlibabaCampusDeviceOpenapiGettemplatelistRequest{
    return &AlibabaCampusDeviceOpenapiGettemplatelistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusDeviceOpenapiGettemplatelistRequest) GetApiMethodName() string {
    return "alibaba.campus.device.openapi.gettemplatelist"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusDeviceOpenapiGettemplatelistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 设备模板查询对象
func (r *AlibabaCampusDeviceOpenapiGettemplatelistRequest) SetQuery(query *TemplateApiQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r AlibabaCampusDeviceOpenapiGettemplatelistRequest) GetQuery() *TemplateApiQuery {
    return r.query
}
