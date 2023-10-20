package tbrefund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaospecialrefundgetAPIRequest 特殊部分退纠纷单查询 API请求
// taobao.special.refund.get
//
// 获取单笔特殊部分退的纠纷单查询
type TaobaospecialrefundgetAPIRequest struct {
	model.Params
	// 需要返回的字段。目前支持有：refund_id, alipay_no, tid, oid, buyer_nick, seller_nick, total_fee, status, created, refund_fee, good_status, has_good_return, payment, reason, desc, num_iid, title, price, num, good_return_time, company_name, sid, address, shipping_type, refund_remind_timeout, refund_phase, refund_version, operation_contraint, attribute, outer_id, sku
	_fields []string
	// 退款单号
	_refundId int64
}

// NewTaobaospecialrefundgetRequest 初始化TaobaospecialrefundgetAPIRequest对象
func NewTaobaospecialrefundgetRequest() *TaobaospecialrefundgetAPIRequest {
	return &TaobaospecialrefundgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaospecialrefundgetAPIRequest) GetApiMethodName() string {
	return "taobao.special.refund.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaospecialrefundgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaospecialrefundgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段。目前支持有：refund_id, alipay_no, tid, oid, buyer_nick, seller_nick, total_fee, status, created, refund_fee, good_status, has_good_return, payment, reason, desc, num_iid, title, price, num, good_return_time, company_name, sid, address, shipping_type, refund_remind_timeout, refund_phase, refund_version, operation_contraint, attribute, outer_id, sku
func (r *TaobaospecialrefundgetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaospecialrefundgetAPIRequest) GetFields() []string {
	return r._fields
}

// SetRefundId is RefundId Setter
// 退款单号
func (r *TaobaospecialrefundgetAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaospecialrefundgetAPIRequest) GetRefundId() int64 {
	return r._refundId
}
