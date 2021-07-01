package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyRefundAgreeAPIRequest
saas 售后逆向 商户同意用户逆向申请 API请求
alibaba.tcls.aelophy.refund.agree

saas 售后逆向 商户同意用户逆向申请 */
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
	_subRefundList []Subrefundlist
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
func (r AlibabaTclsAelophyRefundAgreeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StoreId Setter
// 门店ID
func (r *AlibabaTclsAelophyRefundAgreeAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r AlibabaTclsAelophyRefundAgreeAPIRequest) GetStoreId() string {
	return r._storeId
}

// Set is OutOrderId Setter
// 外部订单ID
func (r *AlibabaTclsAelophyRefundAgreeAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// Get OutOrderId Getter
func (r AlibabaTclsAelophyRefundAgreeAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// Set is RefundId Setter
// 退款单ID
func (r *AlibabaTclsAelophyRefundAgreeAPIRequest) SetRefundId(_refundId string) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// Get RefundId Getter
func (r AlibabaTclsAelophyRefundAgreeAPIRequest) GetRefundId() string {
	return r._refundId
}

// Set is AuditMemo Setter
// 审核说明
func (r *AlibabaTclsAelophyRefundAgreeAPIRequest) SetAuditMemo(_auditMemo string) error {
	r._auditMemo = _auditMemo
	r.Set("audit_memo", _auditMemo)
	return nil
}

// Get AuditMemo Getter
func (r AlibabaTclsAelophyRefundAgreeAPIRequest) GetAuditMemo() string {
	return r._auditMemo
}

// Set is SubRefundList Setter
// 外部子订单列表
func (r *AlibabaTclsAelophyRefundAgreeAPIRequest) SetSubRefundList(_subRefundList []Subrefundlist) error {
	r._subRefundList = _subRefundList
	r.Set("sub_refund_list", _subRefundList)
	return nil
}

// Get SubRefundList Getter
func (r AlibabaTclsAelophyRefundAgreeAPIRequest) GetSubRefundList() []Subrefundlist {
	return r._subRefundList
}
