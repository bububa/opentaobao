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
    _storeId   string
    // 外部订单ID
    _outOrderId   string
    // 退款单ID
    _refundId   string
    // 审核说明
    _auditMemo   string
    // 外部子订单列表
    _subRefundList   []Subrefundlist
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
func (r *AlibabaTclsAelophyRefundAgreeRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaTclsAelophyRefundAgreeRequest) GetStoreId() string {
    return r._storeId
}
// OutOrderId Setter
// 外部订单ID
func (r *AlibabaTclsAelophyRefundAgreeRequest) SetOutOrderId(_outOrderId string) error {
    r._outOrderId = _outOrderId
    r.Set("out_order_id", _outOrderId)
    return nil
}

// OutOrderId Getter
func (r AlibabaTclsAelophyRefundAgreeRequest) GetOutOrderId() string {
    return r._outOrderId
}
// RefundId Setter
// 退款单ID
func (r *AlibabaTclsAelophyRefundAgreeRequest) SetRefundId(_refundId string) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r AlibabaTclsAelophyRefundAgreeRequest) GetRefundId() string {
    return r._refundId
}
// AuditMemo Setter
// 审核说明
func (r *AlibabaTclsAelophyRefundAgreeRequest) SetAuditMemo(_auditMemo string) error {
    r._auditMemo = _auditMemo
    r.Set("audit_memo", _auditMemo)
    return nil
}

// AuditMemo Getter
func (r AlibabaTclsAelophyRefundAgreeRequest) GetAuditMemo() string {
    return r._auditMemo
}
// SubRefundList Setter
// 外部子订单列表
func (r *AlibabaTclsAelophyRefundAgreeRequest) SetSubRefundList(_subRefundList []Subrefundlist) error {
    r._subRefundList = _subRefundList
    r.Set("sub_refund_list", _subRefundList)
    return nil
}

// SubRefundList Getter
func (r AlibabaTclsAelophyRefundAgreeRequest) GetSubRefundList() []Subrefundlist {
    return r._subRefundList
}
