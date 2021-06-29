package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商处理订单接口（订购成功/失败、发货） API请求
alibaba.tianji.supplier.order.result

供应商处理订单接口（订购成功/失败、发货）
*/
type AlibabaTianjiSupplierOrderResultRequest struct {
    model.Params
    // 供应商处理订单结果反馈参数
    _supplierOrderResultModel   *SupplierOrderResultModel
}

// 初始化AlibabaTianjiSupplierOrderResultRequest对象
func NewAlibabaTianjiSupplierOrderResultRequest() *AlibabaTianjiSupplierOrderResultRequest{
    return &AlibabaTianjiSupplierOrderResultRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTianjiSupplierOrderResultRequest) GetApiMethodName() string {
    return "alibaba.tianji.supplier.order.result"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTianjiSupplierOrderResultRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SupplierOrderResultModel Setter
// 供应商处理订单结果反馈参数
func (r *AlibabaTianjiSupplierOrderResultRequest) SetSupplierOrderResultModel(_supplierOrderResultModel *SupplierOrderResultModel) error {
    r._supplierOrderResultModel = _supplierOrderResultModel
    r.Set("supplier_order_result_model", _supplierOrderResultModel)
    return nil
}

// SupplierOrderResultModel Getter
func (r AlibabaTianjiSupplierOrderResultRequest) GetSupplierOrderResultModel() *SupplierOrderResultModel {
    return r._supplierOrderResultModel
}
