package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV绑定外部门店id和外部商品id API请求
alibaba.alihealth.dental.item.bind

ISV绑定外部门店id和外部商品id
*/
type AlibabaAlihealthDentalItemBindAPIRequest struct {
    model.Params
    // bind_list
    _bindList   []StoreItemRelRequest
    // 类型 1 天猫门店 2 支付宝门店
    _type   int64
}

// 初始化AlibabaAlihealthDentalItemBindAPIRequest对象
func NewAlibabaAlihealthDentalItemBindRequest() *AlibabaAlihealthDentalItemBindAPIRequest{
    return &AlibabaAlihealthDentalItemBindAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalItemBindAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dental.item.bind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalItemBindAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BindList Setter
// bind_list
func (r *AlibabaAlihealthDentalItemBindAPIRequest) SetBindList(_bindList []StoreItemRelRequest) error {
    r._bindList = _bindList
    r.Set("bind_list", _bindList)
    return nil
}

// BindList Getter
func (r AlibabaAlihealthDentalItemBindAPIRequest) GetBindList() []StoreItemRelRequest {
    return r._bindList
}
// Type Setter
// 类型 1 天猫门店 2 支付宝门店
func (r *AlibabaAlihealthDentalItemBindAPIRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaAlihealthDentalItemBindAPIRequest) GetType() int64 {
    return r._type
}
