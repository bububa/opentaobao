package pur

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
asn创建 API请求
alibaba.pur.supplier.asncreate

asn创建
*/
type AlibabaPurSupplierAsncreateRequest struct {
    model.Params
    // asn头信息
    asn   *SupplierAsnInfoVO
}

// 初始化AlibabaPurSupplierAsncreateRequest对象
func NewAlibabaPurSupplierAsncreateRequest() *AlibabaPurSupplierAsncreateRequest{
    return &AlibabaPurSupplierAsncreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPurSupplierAsncreateRequest) GetApiMethodName() string {
    return "alibaba.pur.supplier.asncreate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPurSupplierAsncreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Asn Setter
// asn头信息
func (r *AlibabaPurSupplierAsncreateRequest) SetAsn(asn *SupplierAsnInfoVO) error {
    r.asn = asn
    r.Set("asn", asn)
    return nil
}

// Asn Getter
func (r AlibabaPurSupplierAsncreateRequest) GetAsn() *SupplierAsnInfoVO {
    return r.asn
}
