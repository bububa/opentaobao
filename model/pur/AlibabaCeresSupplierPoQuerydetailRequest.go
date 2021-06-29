package pur

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
采购供应商订单明细查询接口 APIRequest
alibaba.ceres.supplier.po.querydetail

采购供应商订单明细查询接口
*/
type AlibabaCeresSupplierPoQuerydetailRequest struct {
    model.Params

    // 订单编号
    poNo   string 

}

func NewAlibabaCeresSupplierPoQuerydetailRequest() *AlibabaCeresSupplierPoQuerydetailRequest{
    return &AlibabaCeresSupplierPoQuerydetailRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCeresSupplierPoQuerydetailRequest) GetApiMethodName() string {
    return "alibaba.ceres.supplier.po.querydetail"
}

func (r AlibabaCeresSupplierPoQuerydetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCeresSupplierPoQuerydetailRequest) SetPoNo(poNo string) error {
    r.poNo = poNo
    r.Set("po_no", poNo)
    return nil
}

func (r AlibabaCeresSupplierPoQuerydetailRequest) GetPoNo() string {
    return r.poNo
}

