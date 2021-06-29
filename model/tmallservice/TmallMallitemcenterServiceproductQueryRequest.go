package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务商服务产品查询 API请求
tmall.mallitemcenter.serviceproduct.query

查询天猫服务的服务商发布的服务产品
*/
type TmallMallitemcenterServiceproductQueryRequest struct {
    model.Params
    // 服务产品id
    id   int64
    // 产品状态
    status   int64
    // 服务名称
    serviceCode   string
    // 产品类型
    serviceProductType   int64
}

// 初始化TmallMallitemcenterServiceproductQueryRequest对象
func NewTmallMallitemcenterServiceproductQueryRequest() *TmallMallitemcenterServiceproductQueryRequest{
    return &TmallMallitemcenterServiceproductQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallMallitemcenterServiceproductQueryRequest) GetApiMethodName() string {
    return "tmall.mallitemcenter.serviceproduct.query"
}

// IRequest interface 方法, 获取API参数
func (r TmallMallitemcenterServiceproductQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 服务产品id
func (r *TmallMallitemcenterServiceproductQueryRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r TmallMallitemcenterServiceproductQueryRequest) GetId() int64 {
    return r.id
}
// Status Setter
// 产品状态
func (r *TmallMallitemcenterServiceproductQueryRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TmallMallitemcenterServiceproductQueryRequest) GetStatus() int64 {
    return r.status
}
// ServiceCode Setter
// 服务名称
func (r *TmallMallitemcenterServiceproductQueryRequest) SetServiceCode(serviceCode string) error {
    r.serviceCode = serviceCode
    r.Set("service_code", serviceCode)
    return nil
}

// ServiceCode Getter
func (r TmallMallitemcenterServiceproductQueryRequest) GetServiceCode() string {
    return r.serviceCode
}
// ServiceProductType Setter
// 产品类型
func (r *TmallMallitemcenterServiceproductQueryRequest) SetServiceProductType(serviceProductType int64) error {
    r.serviceProductType = serviceProductType
    r.Set("service_product_type", serviceProductType)
    return nil
}

// ServiceProductType Getter
func (r TmallMallitemcenterServiceproductQueryRequest) GetServiceProductType() int64 {
    return r.serviceProductType
}
