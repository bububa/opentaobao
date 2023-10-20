package tbrefund

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSpecialRefundGetAPIRequest 特殊部分退纠纷单查询 API请求
// taobao.special.refund.get
//
// 获取单笔特殊部分退的纠纷单查询
type TaobaoSpecialRefundGetAPIRequest struct {
	model.Params
	// 需要返回的字段。目前支持有：refund_id, alipay_no, tid, oid, buyer_nick, seller_nick, total_fee, status, created, refund_fee, good_status, has_good_return, payment, reason, desc, num_iid, title, price, num, good_return_time, company_name, sid, address, shipping_type, refund_remind_timeout, refund_phase, refund_version, operation_contraint, attribute, outer_id, sku
	_fields []string
	// 退款单号
	_refundId int64
}

// NewTaobaoSpecialRefundGetRequest 初始化TaobaoSpecialRefundGetAPIRequest对象
func NewTaobaoSpecialRefundGetRequest() *TaobaoSpecialRefundGetAPIRequest {
	return &TaobaoSpecialRefundGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSpecialRefundGetAPIRequest) Reset() {
	r._fields = r._fields[:0]
	r._refundId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSpecialRefundGetAPIRequest) GetApiMethodName() string {
	return "taobao.special.refund.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSpecialRefundGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSpecialRefundGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段。目前支持有：refund_id, alipay_no, tid, oid, buyer_nick, seller_nick, total_fee, status, created, refund_fee, good_status, has_good_return, payment, reason, desc, num_iid, title, price, num, good_return_time, company_name, sid, address, shipping_type, refund_remind_timeout, refund_phase, refund_version, operation_contraint, attribute, outer_id, sku
func (r *TaobaoSpecialRefundGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoSpecialRefundGetAPIRequest) GetFields() []string {
	return r._fields
}

// SetRefundId is RefundId Setter
// 退款单号
func (r *TaobaoSpecialRefundGetAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoSpecialRefundGetAPIRequest) GetRefundId() int64 {
	return r._refundId
}

var poolTaobaoSpecialRefundGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSpecialRefundGetRequest()
	},
}

// GetTaobaoSpecialRefundGetRequest 从 sync.Pool 获取 TaobaoSpecialRefundGetAPIRequest
func GetTaobaoSpecialRefundGetAPIRequest() *TaobaoSpecialRefundGetAPIRequest {
	return poolTaobaoSpecialRefundGetAPIRequest.Get().(*TaobaoSpecialRefundGetAPIRequest)
}

// ReleaseTaobaoSpecialRefundGetAPIRequest 将 TaobaoSpecialRefundGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSpecialRefundGetAPIRequest(v *TaobaoSpecialRefundGetAPIRequest) {
	v.Reset()
	poolTaobaoSpecialRefundGetAPIRequest.Put(v)
}
