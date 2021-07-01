package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认收款 API请求
taobao.fenxiao.order.confirm.paid

供应商确认收款（非支付宝交易）。
*/
type TaobaoFenxiaoOrderConfirmPaidAPIRequest struct {
    model.Params
    // 采购单编号。
    _purchaseOrderId   int64
    // 确认支付信息（字数小于100）
    _confirmRemark   string
}

// 初始化TaobaoFenxiaoOrderConfirmPaidAPIRequest对象
func NewTaobaoFenxiaoOrderConfirmPaidRequest() *TaobaoFenxiaoOrderConfirmPaidAPIRequest{
    return &TaobaoFenxiaoOrderConfirmPaidAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoOrderConfirmPaidAPIRequest) GetApiMethodName() string {
    return "taobao.fenxiao.order.confirm.paid"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoOrderConfirmPaidAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PurchaseOrderId Setter
// 采购单编号。
func (r *TaobaoFenxiaoOrderConfirmPaidAPIRequest) SetPurchaseOrderId(_purchaseOrderId int64) error {
    r._purchaseOrderId = _purchaseOrderId
    r.Set("purchase_order_id", _purchaseOrderId)
    return nil
}

// PurchaseOrderId Getter
func (r TaobaoFenxiaoOrderConfirmPaidAPIRequest) GetPurchaseOrderId() int64 {
    return r._purchaseOrderId
}
// ConfirmRemark Setter
// 确认支付信息（字数小于100）
func (r *TaobaoFenxiaoOrderConfirmPaidAPIRequest) SetConfirmRemark(_confirmRemark string) error {
    r._confirmRemark = _confirmRemark
    r.Set("confirm_remark", _confirmRemark)
    return nil
}

// ConfirmRemark Getter
func (r TaobaoFenxiaoOrderConfirmPaidAPIRequest) GetConfirmRemark() string {
    return r._confirmRemark
}
