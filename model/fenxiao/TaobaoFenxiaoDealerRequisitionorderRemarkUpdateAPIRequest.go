package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIRequest 修改经销采购单备注 API请求
// taobao.fenxiao.dealer.requisitionorder.remark.update
//
// 供应商修改经销采购单备注
type TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIRequest struct {
	model.Params
	// 经销采购单ID
	_dealerOrderId int64
	// 备注留言，可为空
	_supplierMemo string
	// 旗子的标记，必选。<br/>1-5之间的数字。<br/>非1-5之间，都采用1作为默认。<br/>1:红色<br/>2:黄色<br/>3:绿色<br/>4:蓝色<br/>5:粉红色
	_supplierMemoFlag int64
}

// NewTaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest 初始化TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIRequest对象
func NewTaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest() *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIRequest {
	return &TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.remark.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DealerOrderId Setter
// 经销采购单ID
func (r *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIRequest) SetDealerOrderId(_dealerOrderId int64) error {
	r._dealerOrderId = _dealerOrderId
	r.Set("dealer_order_id", _dealerOrderId)
	return nil
}

// Get DealerOrderId Getter
func (r TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIRequest) GetDealerOrderId() int64 {
	return r._dealerOrderId
}

// Set is SupplierMemo Setter
// 备注留言，可为空
func (r *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIRequest) SetSupplierMemo(_supplierMemo string) error {
	r._supplierMemo = _supplierMemo
	r.Set("supplier_memo", _supplierMemo)
	return nil
}

// Get SupplierMemo Getter
func (r TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIRequest) GetSupplierMemo() string {
	return r._supplierMemo
}

// Set is SupplierMemoFlag Setter
// 旗子的标记，必选。<br/>1-5之间的数字。<br/>非1-5之间，都采用1作为默认。<br/>1:红色<br/>2:黄色<br/>3:绿色<br/>4:蓝色<br/>5:粉红色
func (r *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIRequest) SetSupplierMemoFlag(_supplierMemoFlag int64) error {
	r._supplierMemoFlag = _supplierMemoFlag
	r.Set("supplier_memo_flag", _supplierMemoFlag)
	return nil
}

// Get SupplierMemoFlag Getter
func (r TaobaoFenxiaoDealerRequisitionorderRemarkUpdateAPIRequest) GetSupplierMemoFlag() int64 {
	return r._supplierMemoFlag
}
