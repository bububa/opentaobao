package pur

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
po反馈创建 API请求
alibaba.pur.supplier.porespcreate

PO反馈接口
*/
type AlibabaPurSupplierPorespcreateRequest struct {
    model.Params
    // PO反馈信息
    _poResponse   []SupplierPoResponseDO
}

// 初始化AlibabaPurSupplierPorespcreateRequest对象
func NewAlibabaPurSupplierPorespcreateRequest() *AlibabaPurSupplierPorespcreateRequest{
    return &AlibabaPurSupplierPorespcreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPurSupplierPorespcreateRequest) GetApiMethodName() string {
    return "alibaba.pur.supplier.porespcreate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPurSupplierPorespcreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PoResponse Setter
// PO反馈信息
func (r *AlibabaPurSupplierPorespcreateRequest) SetPoResponse(_poResponse []SupplierPoResponseDO) error {
    r._poResponse = _poResponse
    r.Set("po_response", _poResponse)
    return nil
}

// PoResponse Getter
func (r AlibabaPurSupplierPorespcreateRequest) GetPoResponse() []SupplierPoResponseDO {
    return r._poResponse
}
