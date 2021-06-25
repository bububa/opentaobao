package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商处理订单接口（订购成功/失败、发货） APIRequest
alibaba.tianji.supplier.order.result

供应商处理订单接口（订购成功/失败、发货）
*/
type AlibabaTianjiSupplierOrderResultRequest struct {
    model.Params

    // 供应商处理订单结果反馈参数
    supplierOrderResultModel   *SupplierOrderResultModel 

}

func NewAlibabaTianjiSupplierOrderResultRequest() *AlibabaTianjiSupplierOrderResultRequest{
    return &AlibabaTianjiSupplierOrderResultRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTianjiSupplierOrderResultRequest) GetApiMethodName() string {
    return "alibaba.tianji.supplier.order.result"
}

func (r AlibabaTianjiSupplierOrderResultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTianjiSupplierOrderResultRequest) SetSupplierOrderResultModel(supplierOrderResultModel *SupplierOrderResultModel) error {
    r.supplierOrderResultModel = supplierOrderResultModel
    r.Set("supplier_order_result_model", supplierOrderResultModel)
    return nil
}

func (r AlibabaTianjiSupplierOrderResultRequest) GetSupplierOrderResultModel() *SupplierOrderResultModel {
    return r.supplierOrderResultModel
}

