package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoOrderRemarkUpdateAPIRequest 修改采购单备注 API请求
// taobao.fenxiao.order.remark.update
//
// 供应商修改采购单备注
type TaobaoFenxiaoOrderRemarkUpdateAPIRequest struct {
	model.Params
	// 备注内容(供应商操作)
	_supplierMemo string
	// 采购单编号
	_purchaseOrderId int64
	// 旗子的标记，1-5之间的数字。非1-5之间，都采用1作为默认。<br/>1:红色<br/>2:黄色<br/>3:绿色<br/>4:蓝色<br/>5:粉红色
	_supplierMemoFlag int64
}

// NewTaobaoFenxiaoOrderRemarkUpdateRequest 初始化TaobaoFenxiaoOrderRemarkUpdateAPIRequest对象
func NewTaobaoFenxiaoOrderRemarkUpdateRequest() *TaobaoFenxiaoOrderRemarkUpdateAPIRequest {
	return &TaobaoFenxiaoOrderRemarkUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoOrderRemarkUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.order.remark.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoOrderRemarkUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSupplierMemo is SupplierMemo Setter
// 备注内容(供应商操作)
func (r *TaobaoFenxiaoOrderRemarkUpdateAPIRequest) SetSupplierMemo(_supplierMemo string) error {
	r._supplierMemo = _supplierMemo
	r.Set("supplier_memo", _supplierMemo)
	return nil
}

// GetSupplierMemo SupplierMemo Getter
func (r TaobaoFenxiaoOrderRemarkUpdateAPIRequest) GetSupplierMemo() string {
	return r._supplierMemo
}

// SetPurchaseOrderId is PurchaseOrderId Setter
// 采购单编号
func (r *TaobaoFenxiaoOrderRemarkUpdateAPIRequest) SetPurchaseOrderId(_purchaseOrderId int64) error {
	r._purchaseOrderId = _purchaseOrderId
	r.Set("purchase_order_id", _purchaseOrderId)
	return nil
}

// GetPurchaseOrderId PurchaseOrderId Getter
func (r TaobaoFenxiaoOrderRemarkUpdateAPIRequest) GetPurchaseOrderId() int64 {
	return r._purchaseOrderId
}

// SetSupplierMemoFlag is SupplierMemoFlag Setter
// 旗子的标记，1-5之间的数字。非1-5之间，都采用1作为默认。<br/>1:红色<br/>2:黄色<br/>3:绿色<br/>4:蓝色<br/>5:粉红色
func (r *TaobaoFenxiaoOrderRemarkUpdateAPIRequest) SetSupplierMemoFlag(_supplierMemoFlag int64) error {
	r._supplierMemoFlag = _supplierMemoFlag
	r.Set("supplier_memo_flag", _supplierMemoFlag)
	return nil
}

// GetSupplierMemoFlag SupplierMemoFlag Getter
func (r TaobaoFenxiaoOrderRemarkUpdateAPIRequest) GetSupplierMemoFlag() int64 {
	return r._supplierMemoFlag
}
