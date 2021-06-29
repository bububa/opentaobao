package pur

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
采购供应商订单明细查询接口 API请求
alibaba.ceres.supplier.po.querydetail

采购供应商订单明细查询接口
*/
type AlibabaCeresSupplierPoQuerydetailRequest struct {
    model.Params
    // 订单编号
    poNo   string
}

// 初始化AlibabaCeresSupplierPoQuerydetailRequest对象
func NewAlibabaCeresSupplierPoQuerydetailRequest() *AlibabaCeresSupplierPoQuerydetailRequest{
    return &AlibabaCeresSupplierPoQuerydetailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCeresSupplierPoQuerydetailRequest) GetApiMethodName() string {
    return "alibaba.ceres.supplier.po.querydetail"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCeresSupplierPoQuerydetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PoNo Setter
// 订单编号
func (r *AlibabaCeresSupplierPoQuerydetailRequest) SetPoNo(poNo string) error {
    r.poNo = poNo
    r.Set("po_no", poNo)
    return nil
}

// PoNo Getter
func (r AlibabaCeresSupplierPoQuerydetailRequest) GetPoNo() string {
    return r.poNo
}
