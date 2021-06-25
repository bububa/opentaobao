package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改采购单备注 APIRequest
taobao.fenxiao.order.remark.update

供应商修改采购单备注
*/
type TaobaoFenxiaoOrderRemarkUpdateRequest struct {
    model.Params

    // 采购单编号
    purchaseOrderId   int64 

    // 备注内容(供应商操作)
    supplierMemo   string 

    // 旗子的标记，1-5之间的数字。非1-5之间，都采用1作为默认。<br/>1:红色<br/>2:黄色<br/>3:绿色<br/>4:蓝色<br/>5:粉红色
    supplierMemoFlag   int64 

}

func NewTaobaoFenxiaoOrderRemarkUpdateRequest() *TaobaoFenxiaoOrderRemarkUpdateRequest{
    return &TaobaoFenxiaoOrderRemarkUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFenxiaoOrderRemarkUpdateRequest) GetApiMethodName() string {
    return "taobao.fenxiao.order.remark.update"
}

func (r TaobaoFenxiaoOrderRemarkUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFenxiaoOrderRemarkUpdateRequest) SetPurchaseOrderId(purchaseOrderId int64) error {
    r.purchaseOrderId = purchaseOrderId
    r.Set("purchase_order_id", purchaseOrderId)
    return nil
}

func (r TaobaoFenxiaoOrderRemarkUpdateRequest) GetPurchaseOrderId() int64 {
    return r.purchaseOrderId
}

func (r *TaobaoFenxiaoOrderRemarkUpdateRequest) SetSupplierMemo(supplierMemo string) error {
    r.supplierMemo = supplierMemo
    r.Set("supplier_memo", supplierMemo)
    return nil
}

func (r TaobaoFenxiaoOrderRemarkUpdateRequest) GetSupplierMemo() string {
    return r.supplierMemo
}

func (r *TaobaoFenxiaoOrderRemarkUpdateRequest) SetSupplierMemoFlag(supplierMemoFlag int64) error {
    r.supplierMemoFlag = supplierMemoFlag
    r.Set("supplier_memo_flag", supplierMemoFlag)
    return nil
}

func (r TaobaoFenxiaoOrderRemarkUpdateRequest) GetSupplierMemoFlag() int64 {
    return r.supplierMemoFlag
}

