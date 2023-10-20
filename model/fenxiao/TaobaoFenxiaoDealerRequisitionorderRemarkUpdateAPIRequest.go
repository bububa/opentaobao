package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaodealerrequisitionorderremarkupdateAPIRequest 修改经销采购单备注 API请求
// taobao.fenxiao.dealer.requisitionorder.remark.update
//
// 供应商修改经销采购单备注
type TaobaofenxiaodealerrequisitionorderremarkupdateAPIRequest struct {
	model.Params
	// 备注留言，可为空
	_supplierMemo string
	// 经销采购单ID
	_dealerOrderId int64
	// 旗子的标记，必选。<br/>1-5之间的数字。<br/>非1-5之间，都采用1作为默认。<br/>1:红色<br/>2:黄色<br/>3:绿色<br/>4:蓝色<br/>5:粉红色
	_supplierMemoFlag int64
}

// NewTaobaofenxiaodealerrequisitionorderremarkupdateRequest 初始化TaobaofenxiaodealerrequisitionorderremarkupdateAPIRequest对象
func NewTaobaofenxiaodealerrequisitionorderremarkupdateRequest() *TaobaofenxiaodealerrequisitionorderremarkupdateAPIRequest {
	return &TaobaofenxiaodealerrequisitionorderremarkupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaodealerrequisitionorderremarkupdateAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.remark.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaodealerrequisitionorderremarkupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaodealerrequisitionorderremarkupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSupplierMemo is SupplierMemo Setter
// 备注留言，可为空
func (r *TaobaofenxiaodealerrequisitionorderremarkupdateAPIRequest) SetSupplierMemo(_supplierMemo string) error {
	r._supplierMemo = _supplierMemo
	r.Set("supplier_memo", _supplierMemo)
	return nil
}

// GetSupplierMemo SupplierMemo Getter
func (r TaobaofenxiaodealerrequisitionorderremarkupdateAPIRequest) GetSupplierMemo() string {
	return r._supplierMemo
}

// SetDealerOrderId is DealerOrderId Setter
// 经销采购单ID
func (r *TaobaofenxiaodealerrequisitionorderremarkupdateAPIRequest) SetDealerOrderId(_dealerOrderId int64) error {
	r._dealerOrderId = _dealerOrderId
	r.Set("dealer_order_id", _dealerOrderId)
	return nil
}

// GetDealerOrderId DealerOrderId Getter
func (r TaobaofenxiaodealerrequisitionorderremarkupdateAPIRequest) GetDealerOrderId() int64 {
	return r._dealerOrderId
}

// SetSupplierMemoFlag is SupplierMemoFlag Setter
// 旗子的标记，必选。&lt;br/&gt;1-5之间的数字。&lt;br/&gt;非1-5之间，都采用1作为默认。&lt;br/&gt;1:红色&lt;br/&gt;2:黄色&lt;br/&gt;3:绿色&lt;br/&gt;4:蓝色&lt;br/&gt;5:粉红色
func (r *TaobaofenxiaodealerrequisitionorderremarkupdateAPIRequest) SetSupplierMemoFlag(_supplierMemoFlag int64) error {
	r._supplierMemoFlag = _supplierMemoFlag
	r.Set("supplier_memo_flag", _supplierMemoFlag)
	return nil
}

// GetSupplierMemoFlag SupplierMemoFlag Getter
func (r TaobaofenxiaodealerrequisitionorderremarkupdateAPIRequest) GetSupplierMemoFlag() int64 {
	return r._supplierMemoFlag
}
