package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认收款 APIRequest
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

func NewTaobaoFenxiaoOrderConfirmPaidRequest() *TaobaoFenxiaoOrderConfirmPaidRequest{
    return &TaobaoFenxiaoOrderConfirmPaidRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoOrderConfirmPaidRequest) GetApiMethodName() string {
    return "taobao.fenxiao.order.confirm.paid"
}

func (r TaobaoFenxiaoOrderConfirmPaidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoOrderConfirmPaidRequest) SetPurchaseOrderId(purchaseOrderId int64) error {
    r.purchaseOrderId = purchaseOrderId
    r.Set("purchase_order_id", purchaseOrderId)
    return nil
}

func (r TaobaoFenxiaoOrderConfirmPaidRequest) GetPurchaseOrderId() int64 {
    return r.purchaseOrderId
}

func (r *TaobaoFenxiaoOrderConfirmPaidRequest) SetConfirmRemark(confirmRemark string) error {
    r.confirmRemark = confirmRemark
    r.Set("confirm_remark", confirmRemark)
    return nil
}

func (r TaobaoFenxiaoOrderConfirmPaidRequest) GetConfirmRemark() string {
    return r.confirmRemark
}

