package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里通信运营商信息回传 API请求
alibaba.base.order.supplier.notify

接收阿里通信流量运营商信息回传
*/
type AlibabaBaseOrderSupplierNotifyRequest struct {
    model.Params
    // 入参对象
    paramFlowSuppllierNotifyModel   *FlowSuppllierNotifyModel
}

// 初始化AlibabaBaseOrderSupplierNotifyRequest对象
func NewAlibabaBaseOrderSupplierNotifyRequest() *AlibabaBaseOrderSupplierNotifyRequest{
    return &AlibabaBaseOrderSupplierNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaBaseOrderSupplierNotifyRequest) GetApiMethodName() string {
    return "alibaba.base.order.supplier.notify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaBaseOrderSupplierNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamFlowSuppllierNotifyModel Setter
// 入参对象
func (r *AlibabaBaseOrderSupplierNotifyRequest) SetParamFlowSuppllierNotifyModel(paramFlowSuppllierNotifyModel *FlowSuppllierNotifyModel) error {
    r.paramFlowSuppllierNotifyModel = paramFlowSuppllierNotifyModel
    r.Set("param_flow_suppllier_notify_model", paramFlowSuppllierNotifyModel)
    return nil
}

// ParamFlowSuppllierNotifyModel Getter
func (r AlibabaBaseOrderSupplierNotifyRequest) GetParamFlowSuppllierNotifyModel() *FlowSuppllierNotifyModel {
    return r.paramFlowSuppllierNotifyModel
}
