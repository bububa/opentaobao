package pur

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
po反馈创建 APIRequest
alibaba.pur.supplier.porespcreate

PO反馈接口
*/
type AlibabaPurSupplierPorespcreateRequest struct {
    model.Params

    // PO反馈信息
    poResponse   []SupplierPoResponseDO 

}

func NewAlibabaPurSupplierPorespcreateRequest() *AlibabaPurSupplierPorespcreateRequest{
    return &AlibabaPurSupplierPorespcreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaPurSupplierPorespcreateRequest) GetApiMethodName() string {
    return "alibaba.pur.supplier.porespcreate"
}

func (r AlibabaPurSupplierPorespcreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaPurSupplierPorespcreateRequest) SetPoResponse(poResponse []SupplierPoResponseDO) error {
    r.poResponse = poResponse
    r.Set("po_response", poResponse)
    return nil
}

func (r AlibabaPurSupplierPorespcreateRequest) GetPoResponse() []SupplierPoResponseDO {
    return r.poResponse
}

