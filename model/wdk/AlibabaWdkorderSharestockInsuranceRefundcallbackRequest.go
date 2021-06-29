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
    _tbSubOrderId   int64
    // 退款单ID
    _refundId   string
    // 理赔单ID
    _claimId   string
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
func (r *AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) SetTbSubOrderId(_tbSubOrderId int64) error {
    r._tbSubOrderId = _tbSubOrderId
    r.Set("tb_sub_order_id", _tbSubOrderId)
    return nil
}

// TbSubOrderId Getter
func (r AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) GetTbSubOrderId() int64 {
    return r._tbSubOrderId
}
// RefundId Setter
// 退款单ID
func (r *AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) SetRefundId(_refundId string) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) GetRefundId() string {
    return r._refundId
}
// ClaimId Setter
// 理赔单ID
func (r *AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) SetClaimId(_claimId string) error {
    r._claimId = _claimId
    r.Set("claim_id", _claimId)
    return nil
}

// ClaimId Getter
func (r AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) GetClaimId() string {
    return r._claimId
}
