package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改采购单备注 API请求
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

// 初始化TaobaoFenxiaoOrderRemarkUpdateRequest对象
func NewTaobaoFenxiaoOrderRemarkUpdateRequest() *TaobaoFenxiaoOrderRemarkUpdateRequest{
    return &TaobaoFenxiaoOrderRemarkUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoOrderRemarkUpdateRequest) GetApiMethodName() string {
    return "taobao.fenxiao.order.remark.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoOrderRemarkUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PurchaseOrderId Setter
// 采购单编号
func (r *TaobaoFenxiaoOrderRemarkUpdateRequest) SetPurchaseOrderId(purchaseOrderId int64) error {
    r.purchaseOrderId = purchaseOrderId
    r.Set("purchase_order_id", purchaseOrderId)
    return nil
}

// PurchaseOrderId Getter
func (r TaobaoFenxiaoOrderRemarkUpdateRequest) GetPurchaseOrderId() int64 {
    return r.purchaseOrderId
}
// SupplierMemo Setter
// 备注内容(供应商操作)
func (r *TaobaoFenxiaoOrderRemarkUpdateRequest) SetSupplierMemo(supplierMemo string) error {
    r.supplierMemo = supplierMemo
    r.Set("supplier_memo", supplierMemo)
    return nil
}

// SupplierMemo Getter
func (r TaobaoFenxiaoOrderRemarkUpdateRequest) GetSupplierMemo() string {
    return r.supplierMemo
}
// SupplierMemoFlag Setter
// 旗子的标记，1-5之间的数字。非1-5之间，都采用1作为默认。<br/>1:红色<br/>2:黄色<br/>3:绿色<br/>4:蓝色<br/>5:粉红色
func (r *TaobaoFenxiaoOrderRemarkUpdateRequest) SetSupplierMemoFlag(supplierMemoFlag int64) error {
    r.supplierMemoFlag = supplierMemoFlag
    r.Set("supplier_memo_flag", supplierMemoFlag)
    return nil
}

// SupplierMemoFlag Getter
func (r TaobaoFenxiaoOrderRemarkUpdateRequest) GetSupplierMemoFlag() int64 {
    return r.supplierMemoFlag
}
