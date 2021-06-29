package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
共享库存逆向订单理赔单回传 API请求
alibaba.wdkorder.sharestock.insurance.refundcallback

共享库存逆向订单理赔单回传
*/
type AlibabaWdkorderSharestockInsuranceRefundcallbackRequest struct {
    model.Params
    // 淘宝交易子单ID
    tbSubOrderId   int64
    // 退款单ID
    refundId   string
    // 理赔单ID
    claimId   string
}

// 初始化AlibabaWdkorderSharestockInsuranceRefundcallbackRequest对象
func NewAlibabaWdkorderSharestockInsuranceRefundcallbackRequest() *AlibabaWdkorderSharestockInsuranceRefundcallbackRequest{
    return &AlibabaWdkorderSharestockInsuranceRefundcallbackRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) GetApiMethodName() string {
    return "alibaba.wdkorder.sharestock.insurance.refundcallback"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TbSubOrderId Setter
// 淘宝交易子单ID
func (r *AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) SetTbSubOrderId(tbSubOrderId int64) error {
    r.tbSubOrderId = tbSubOrderId
    r.Set("tb_sub_order_id", tbSubOrderId)
    return nil
}

// TbSubOrderId Getter
func (r AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) GetTbSubOrderId() int64 {
    return r.tbSubOrderId
}
// RefundId Setter
// 退款单ID
func (r *AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) SetRefundId(refundId string) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

// RefundId Getter
func (r AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) GetRefundId() string {
    return r.refundId
}
// ClaimId Setter
// 理赔单ID
func (r *AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) SetClaimId(claimId string) error {
    r.claimId = claimId
    r.Set("claim_id", claimId)
    return nil
}

// ClaimId Getter
func (r AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) GetClaimId() string {
    return r.claimId
}
