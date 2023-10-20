package fenxiao

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoOrderRemarkUpdateAPIRequest) Reset() {
	r._supplierMemo = ""
	r._purchaseOrderId = 0
	r._supplierMemoFlag = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoOrderRemarkUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.order.remark.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoOrderRemarkUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoOrderRemarkUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
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
// 旗子的标记，1-5之间的数字。非1-5之间，都采用1作为默认。&lt;br/&gt;1:红色&lt;br/&gt;2:黄色&lt;br/&gt;3:绿色&lt;br/&gt;4:蓝色&lt;br/&gt;5:粉红色
func (r *TaobaoFenxiaoOrderRemarkUpdateAPIRequest) SetSupplierMemoFlag(_supplierMemoFlag int64) error {
	r._supplierMemoFlag = _supplierMemoFlag
	r.Set("supplier_memo_flag", _supplierMemoFlag)
	return nil
}

// GetSupplierMemoFlag SupplierMemoFlag Getter
func (r TaobaoFenxiaoOrderRemarkUpdateAPIRequest) GetSupplierMemoFlag() int64 {
	return r._supplierMemoFlag
}

var poolTaobaoFenxiaoOrderRemarkUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoOrderRemarkUpdateRequest()
	},
}

// GetTaobaoFenxiaoOrderRemarkUpdateRequest 从 sync.Pool 获取 TaobaoFenxiaoOrderRemarkUpdateAPIRequest
func GetTaobaoFenxiaoOrderRemarkUpdateAPIRequest() *TaobaoFenxiaoOrderRemarkUpdateAPIRequest {
	return poolTaobaoFenxiaoOrderRemarkUpdateAPIRequest.Get().(*TaobaoFenxiaoOrderRemarkUpdateAPIRequest)
}

// ReleaseTaobaoFenxiaoOrderRemarkUpdateAPIRequest 将 TaobaoFenxiaoOrderRemarkUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoOrderRemarkUpdateAPIRequest(v *TaobaoFenxiaoOrderRemarkUpdateAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoOrderRemarkUpdateAPIRequest.Put(v)
}
