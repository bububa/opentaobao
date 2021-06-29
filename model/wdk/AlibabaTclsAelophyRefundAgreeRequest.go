package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
saas 售后逆向 商户同意用户逆向申请 API请求
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

// 初始化AlibabaTclsAelophyRefundAgreeRequest对象
func NewAlibabaTclsAelophyRefundAgreeRequest() *AlibabaTclsAelophyRefundAgreeRequest{
    return &AlibabaTclsAelophyRefundAgreeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyRefundAgreeRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.refund.agree"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyRefundAgreeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreId Setter
// 门店ID
func (r *AlibabaTclsAelophyRefundAgreeRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaTclsAelophyRefundAgreeRequest) GetStoreId() string {
    return r.storeId
}
// OutOrderId Setter
// 外部订单ID
func (r *AlibabaTclsAelophyRefundAgreeRequest) SetOutOrderId(outOrderId string) error {
    r.outOrderId = outOrderId
    r.Set("out_order_id", outOrderId)
    return nil
}

// OutOrderId Getter
func (r AlibabaTclsAelophyRefundAgreeRequest) GetOutOrderId() string {
    return r.outOrderId
}
// RefundId Setter
// 退款单ID
func (r *AlibabaTclsAelophyRefundAgreeRequest) SetRefundId(refundId string) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

// RefundId Getter
func (r AlibabaTclsAelophyRefundAgreeRequest) GetRefundId() string {
    return r.refundId
}
// AuditMemo Setter
// 审核说明
func (r *AlibabaTclsAelophyRefundAgreeRequest) SetAuditMemo(auditMemo string) error {
    r.auditMemo = auditMemo
    r.Set("audit_memo", auditMemo)
    return nil
}

// AuditMemo Getter
func (r AlibabaTclsAelophyRefundAgreeRequest) GetAuditMemo() string {
    return r.auditMemo
}
// SubRefundList Setter
// 外部子订单列表
func (r *AlibabaTclsAelophyRefundAgreeRequest) SetSubRefundList(subRefundList []Subrefundlist) error {
    r.subRefundList = subRefundList
    r.Set("sub_refund_list", subRefundList)
    return nil
}

// SubRefundList Getter
func (r AlibabaTclsAelophyRefundAgreeRequest) GetSubRefundList() []Subrefundlist {
    return r.subRefundList
}
