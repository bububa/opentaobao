package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoorderconfirmpaidAPIRequest 确认收款 API请求
// taobao.fenxiao.order.confirm.paid
//
// 供应商确认收款（非支付宝交易）。
type TaobaofenxiaoorderconfirmpaidAPIRequest struct {
	model.Params
	// 确认支付信息（字数小于100）
	_confirmRemark string
	// 采购单编号。
	_purchaseOrderId int64
}

// NewTaobaofenxiaoorderconfirmpaidRequest 初始化TaobaofenxiaoorderconfirmpaidAPIRequest对象
func NewTaobaofenxiaoorderconfirmpaidRequest() *TaobaofenxiaoorderconfirmpaidAPIRequest {
	return &TaobaofenxiaoorderconfirmpaidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaoorderconfirmpaidAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.order.confirm.paid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaoorderconfirmpaidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaoorderconfirmpaidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConfirmRemark is ConfirmRemark Setter
// 确认支付信息（字数小于100）
func (r *TaobaofenxiaoorderconfirmpaidAPIRequest) SetConfirmRemark(_confirmRemark string) error {
	r._confirmRemark = _confirmRemark
	r.Set("confirm_remark", _confirmRemark)
	return nil
}

// GetConfirmRemark ConfirmRemark Getter
func (r TaobaofenxiaoorderconfirmpaidAPIRequest) GetConfirmRemark() string {
	return r._confirmRemark
}

// SetPurchaseOrderId is PurchaseOrderId Setter
// 采购单编号。
func (r *TaobaofenxiaoorderconfirmpaidAPIRequest) SetPurchaseOrderId(_purchaseOrderId int64) error {
	r._purchaseOrderId = _purchaseOrderId
	r.Set("purchase_order_id", _purchaseOrderId)
	return nil
}

// GetPurchaseOrderId PurchaseOrderId Getter
func (r TaobaofenxiaoorderconfirmpaidAPIRequest) GetPurchaseOrderId() int64 {
	return r._purchaseOrderId
}
