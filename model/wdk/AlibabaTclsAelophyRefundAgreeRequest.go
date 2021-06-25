package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
saas 售后逆向 商户同意用户逆向申请 APIRequest
alibaba.tcls.aelophy.refund.agree

saas 售后逆向 商户同意用户逆向申请
*/
type AlibabaTclsAelophyRefundAgreeRequest struct {
    model.Params

    // 门店ID
    storeId   string 

    // 外部订单ID
    outOrderId   string 

    // 退款单ID
    refundId   string 

    // 审核说明
    auditMemo   string 

    // 外部子订单列表
    subRefundList   []Subrefundlist 

}

func NewAlibabaTclsAelophyRefundAgreeRequest() *AlibabaTclsAelophyRefundAgreeRequest{
    return &AlibabaTclsAelophyRefundAgreeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTclsAelophyRefundAgreeRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.refund.agree"
}

func (r AlibabaTclsAelophyRefundAgreeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTclsAelophyRefundAgreeRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaTclsAelophyRefundAgreeRequest) GetStoreId() string {
    return r.storeId
}

func (r *AlibabaTclsAelophyRefundAgreeRequest) SetOutOrderId(outOrderId string) error {
    r.outOrderId = outOrderId
    r.Set("out_order_id", outOrderId)
    return nil
}

func (r AlibabaTclsAelophyRefundAgreeRequest) GetOutOrderId() string {
    return r.outOrderId
}

func (r *AlibabaTclsAelophyRefundAgreeRequest) SetRefundId(refundId string) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r AlibabaTclsAelophyRefundAgreeRequest) GetRefundId() string {
    return r.refundId
}

func (r *AlibabaTclsAelophyRefundAgreeRequest) SetAuditMemo(auditMemo string) error {
    r.auditMemo = auditMemo
    r.Set("audit_memo", auditMemo)
    return nil
}

func (r AlibabaTclsAelophyRefundAgreeRequest) GetAuditMemo() string {
    return r.auditMemo
}

func (r *AlibabaTclsAelophyRefundAgreeRequest) SetSubRefundList(subRefundList []Subrefundlist) error {
    r.subRefundList = subRefundList
    r.Set("sub_refund_list", subRefundList)
    return nil
}

func (r AlibabaTclsAelophyRefundAgreeRequest) GetSubRefundList() []Subrefundlist {
    return r.subRefundList
}

