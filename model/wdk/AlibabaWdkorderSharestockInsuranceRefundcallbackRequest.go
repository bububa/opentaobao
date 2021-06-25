package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
共享库存逆向订单理赔单回传 APIRequest
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

func NewAlibabaWdkorderSharestockInsuranceRefundcallbackRequest() *AlibabaWdkorderSharestockInsuranceRefundcallbackRequest{
    return &AlibabaWdkorderSharestockInsuranceRefundcallbackRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) GetApiMethodName() string {
    return "alibaba.wdkorder.sharestock.insurance.refundcallback"
}

func (r AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) SetTbSubOrderId(tbSubOrderId int64) error {
    r.tbSubOrderId = tbSubOrderId
    r.Set("tb_sub_order_id", tbSubOrderId)
    return nil
}

func (r AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) GetTbSubOrderId() int64 {
    return r.tbSubOrderId
}

func (r *AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) SetRefundId(refundId string) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) GetRefundId() string {
    return r.refundId
}

func (r *AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) SetClaimId(claimId string) error {
    r.claimId = claimId
    r.Set("claim_id", claimId)
    return nil
}

func (r AlibabaWdkorderSharestockInsuranceRefundcallbackRequest) GetClaimId() string {
    return r.claimId
}

