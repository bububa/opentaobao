package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyRefundAgreeAPIRequest saas 售后逆向 商户同意用户逆向申请 API请求
// alibaba.tcls.aelophy.refund.agree
//
// saas 售后逆向 商户同意用户逆向申请
type AlibabaTclsAelophyRefundAgreeAPIRequest struct {
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

// NewAlibabaTclsAelophyRefundAgreeRequest 初始化AlibabaTclsAelophyRefundAgreeAPIRequest对象
func NewAlibabaTclsAelophyRefundAgreeRequest() *AlibabaTclsAelophyRefundAgreeAPIRequest {
	return &AlibabaTclsAelophyRefundAgreeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyRefundAgreeAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.refund.agree"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyRefundAgreeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyRefundAgreeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 门店ID
func (r *AlibabaTclsAelophyRefundAgreeAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaTclsAelophyRefundAgreeAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetOutOrderId is OutOrderId Setter
// 外部订单ID
func (r *AlibabaTclsAelophyRefundAgreeAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlibabaTclsAelophyRefundAgreeAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetRefundId is RefundId Setter
// 退款单ID
func (r *AlibabaTclsAelophyRefundAgreeAPIRequest) SetRefundId(_refundId string) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r AlibabaTclsAelophyRefundAgreeAPIRequest) GetRefundId() string {
	return r._refundId
}

// SetAuditMemo is AuditMemo Setter
// 审核说明
func (r *AlibabaTclsAelophyRefundAgreeAPIRequest) SetAuditMemo(_auditMemo string) error {
	r._auditMemo = _auditMemo
	r.Set("audit_memo", _auditMemo)
	return nil
}

// GetAuditMemo AuditMemo Getter
func (r AlibabaTclsAelophyRefundAgreeAPIRequest) GetAuditMemo() string {
	return r._auditMemo
}

// SetSubRefundList is SubRefundList Setter
// 外部子订单列表
func (r *AlibabaTclsAelophyRefundAgreeAPIRequest) SetSubRefundList(_subRefundList *Subrefundlist) error {
	r._subRefundList = _subRefundList
	r.Set("sub_refund_list", _subRefundList)
	return nil
}

// GetSubRefundList SubRefundList Getter
func (r AlibabaTclsAelophyRefundAgreeAPIRequest) GetSubRefundList() *Subrefundlist {
	return r._subRefundList
}

// SetOrderFrom is OrderFrom Setter
// 渠道
func (r *AlibabaTclsAelophyRefundAgreeAPIRequest) SetOrderFrom(_orderFrom int64) error {
	r._orderFrom = _orderFrom
	r.Set("order_from", _orderFrom)
	return nil
}

// GetOrderFrom OrderFrom Getter
func (r AlibabaTclsAelophyRefundAgreeAPIRequest) GetOrderFrom() int64 {
	return r._orderFrom
}
