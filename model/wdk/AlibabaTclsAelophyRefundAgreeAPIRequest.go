package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophyrefundagreeAPIRequest saas 售后逆向 商户同意用户逆向申请 API请求
// alibaba.tcls.aelophy.refund.agree
//
// saas 售后逆向 商户同意用户逆向申请
type AlibabatclsaelophyrefundagreeAPIRequest struct {
	model.Params
	// 门店ID
	_storeId string
	// 外部订单ID
	_outOrderId string
	// 退款单ID
	_refundId string
	// 审核说明
	_auditMemo string
	// 外部子订单列表
	_subRefundList *Subrefundlist
	// 渠道
	_orderFrom int64
}

// NewAlibabatclsaelophyrefundagreeRequest 初始化AlibabatclsaelophyrefundagreeAPIRequest对象
func NewAlibabatclsaelophyrefundagreeRequest() *AlibabatclsaelophyrefundagreeAPIRequest {
	return &AlibabatclsaelophyrefundagreeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophyrefundagreeAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.refund.agree"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophyrefundagreeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophyrefundagreeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *AlibabatclsaelophyrefundagreeAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabatclsaelophyrefundagreeAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetOutOrderId is OutOrderId Setter
// 外部订单ID
func (r *AlibabatclsaelophyrefundagreeAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlibabatclsaelophyrefundagreeAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetRefundId is RefundId Setter
// 退款单ID
func (r *AlibabatclsaelophyrefundagreeAPIRequest) SetRefundId(_refundId string) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r AlibabatclsaelophyrefundagreeAPIRequest) GetRefundId() string {
	return r._refundId
}

// SetAuditMemo is AuditMemo Setter
// 审核说明
func (r *AlibabatclsaelophyrefundagreeAPIRequest) SetAuditMemo(_auditMemo string) error {
	r._auditMemo = _auditMemo
	r.Set("audit_memo", _auditMemo)
	return nil
}

// GetAuditMemo AuditMemo Getter
func (r AlibabatclsaelophyrefundagreeAPIRequest) GetAuditMemo() string {
	return r._auditMemo
}

// SetSubRefundList is SubRefundList Setter
// 外部子订单列表
func (r *AlibabatclsaelophyrefundagreeAPIRequest) SetSubRefundList(_subRefundList *Subrefundlist) error {
	r._subRefundList = _subRefundList
	r.Set("sub_refund_list", _subRefundList)
	return nil
}

// GetSubRefundList SubRefundList Getter
func (r AlibabatclsaelophyrefundagreeAPIRequest) GetSubRefundList() *Subrefundlist {
	return r._subRefundList
}

// SetOrderFrom is OrderFrom Setter
// 渠道
func (r *AlibabatclsaelophyrefundagreeAPIRequest) SetOrderFrom(_orderFrom int64) error {
	r._orderFrom = _orderFrom
	r.Set("order_from", _orderFrom)
	return nil
}

// GetOrderFrom OrderFrom Getter
func (r AlibabatclsaelophyrefundagreeAPIRequest) GetOrderFrom() int64 {
	return r._orderFrom
}
