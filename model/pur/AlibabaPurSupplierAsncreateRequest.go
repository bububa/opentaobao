package pur

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
asn创建 APIRequest
alibaba.pur.supplier.asncreate

asn创建
*/
type AlibabaPurSupplierAsncreateRequest struct {
    model.Params

    // asn头信息
    asn   *SupplierAsnInfoVO 

}

func NewAlibabaPurSupplierAsncreateRequest() *AlibabaPurSupplierAsncreateRequest{
    return &AlibabaPurSupplierAsncreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaPurSupplierAsncreateRequest) GetApiMethodName() string {
    return "alibaba.pur.supplier.asncreate"
}

func (r AlibabaPurSupplierAsncreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaPurSupplierAsncreateRequest) SetAsn(asn *SupplierAsnInfoVO) error {
    r.asn = asn
    r.Set("asn", asn)
    return nil
}

func (r AlibabaPurSupplierAsncreateRequest) GetAsn() *SupplierAsnInfoVO {
    return r.asn
}

