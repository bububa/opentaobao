package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
共享库存投保业务售后逆向订单数据获取 API请求
alibaba.wdkorder.sharestock.insurance.refundget

共享库存投保业务售后逆向订单数据获取
*/
type AlibabaWdkorderSharestockInsuranceRefundgetRequest struct {
    model.Params
    // 淘宝子订单ID
    tbSubOrderId   string
    // 退货单ID
    refundId   string
}

// 初始化AlibabaWdkorderSharestockInsuranceRefundgetRequest对象
func NewAlibabaWdkorderSharestockInsuranceRefundgetRequest() *AlibabaWdkorderSharestockInsuranceRefundgetRequest{
    return &AlibabaWdkorderSharestockInsuranceRefundgetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkorderSharestockInsuranceRefundgetRequest) GetApiMethodName() string {
    return "alibaba.wdkorder.sharestock.insurance.refundget"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkorderSharestockInsuranceRefundgetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TbSubOrderId Setter
// 淘宝子订单ID
func (r *AlibabaWdkorderSharestockInsuranceRefundgetRequest) SetTbSubOrderId(tbSubOrderId string) error {
    r.tbSubOrderId = tbSubOrderId
    r.Set("tb_sub_order_id", tbSubOrderId)
    return nil
}

// TbSubOrderId Getter
func (r AlibabaWdkorderSharestockInsuranceRefundgetRequest) GetTbSubOrderId() string {
    return r.tbSubOrderId
}
// RefundId Setter
// 退货单ID
func (r *AlibabaWdkorderSharestockInsuranceRefundgetRequest) SetRefundId(refundId string) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

// RefundId Getter
func (r AlibabaWdkorderSharestockInsuranceRefundgetRequest) GetRefundId() string {
    return r.refundId
}
