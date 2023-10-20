package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoYphRefundGetAPIRequest 一盘货商家单个查询退款单信息 API请求
// taobao.fenxiao.yph.refund.get
//
// 一盘货商家单个查询退款单信息
type TaobaoFenxiaoYphRefundGetAPIRequest struct {
	model.Params
	// 分销退款单id，此参数必填
	_refundId int64
	// 是否查询前台消费者退款
	_queryB2cRefund bool
}

// NewTaobaoFenxiaoYphRefundGetRequest 初始化TaobaoFenxiaoYphRefundGetAPIRequest对象
func NewTaobaoFenxiaoYphRefundGetRequest() *TaobaoFenxiaoYphRefundGetAPIRequest {
	return &TaobaoFenxiaoYphRefundGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoYphRefundGetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.yph.refund.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoYphRefundGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoYphRefundGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundId is RefundId Setter
// 分销退款单id，此参数必填
func (r *TaobaoFenxiaoYphRefundGetAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoFenxiaoYphRefundGetAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// SetQueryB2cRefund is QueryB2cRefund Setter
// 是否查询前台消费者退款
func (r *TaobaoFenxiaoYphRefundGetAPIRequest) SetQueryB2cRefund(_queryB2cRefund bool) error {
	r._queryB2cRefund = _queryB2cRefund
	r.Set("query_b2c_refund", _queryB2cRefund)
	return nil
}

// GetQueryB2cRefund QueryB2cRefund Getter
func (r TaobaoFenxiaoYphRefundGetAPIRequest) GetQueryB2cRefund() bool {
	return r._queryB2cRefund
}
