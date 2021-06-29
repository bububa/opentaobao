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
type TaobaoFenxiaoOrderConfirmPaidRequest struct {
    model.Params
    // 采购单编号。
    purchaseOrderId   int64
    // 确认支付信息（字数小于100）
    confirmRemark   string
}

// 初始化TaobaoFenxiaoOrderConfirmPaidRequest对象
func NewTaobaoFenxiaoOrderConfirmPaidRequest() *TaobaoFenxiaoOrderConfirmPaidRequest{
    return &TaobaoFenxiaoOrderConfirmPaidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoOrderConfirmPaidRequest) GetApiMethodName() string {
    return "taobao.fenxiao.order.confirm.paid"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoOrderConfirmPaidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PurchaseOrderId Setter
// 采购单编号。
func (r *TaobaoFenxiaoOrderConfirmPaidRequest) SetPurchaseOrderId(purchaseOrderId int64) error {
    r.purchaseOrderId = purchaseOrderId
    r.Set("purchase_order_id", purchaseOrderId)
    return nil
}

// PurchaseOrderId Getter
func (r TaobaoFenxiaoOrderConfirmPaidRequest) GetPurchaseOrderId() int64 {
    return r.purchaseOrderId
}
// ConfirmRemark Setter
// 确认支付信息（字数小于100）
func (r *TaobaoFenxiaoOrderConfirmPaidRequest) SetConfirmRemark(confirmRemark string) error {
    r.confirmRemark = confirmRemark
    r.Set("confirm_remark", confirmRemark)
    return nil
}

// ConfirmRemark Getter
func (r TaobaoFenxiaoOrderConfirmPaidRequest) GetConfirmRemark() string {
    return r.confirmRemark
}
