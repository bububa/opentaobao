package lstlogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstshipordercancelAPIRequest 零售通发货单取消 API请求
// alibaba.lst.shiporder.cancel
//
// 通过该接口可以取消零售通运保保发货单，并处理相关业务流程。
type AlibabalstshipordercancelAPIRequest struct {
	model.Params
	// 需要退款的明细ID
	_detailOrderIds []string
	// 取消原因
	_reason string
	// 订单号
	_outOrderId string
}

// NewAlibabalstshipordercancelRequest 初始化AlibabalstshipordercancelAPIRequest对象
func NewAlibabalstshipordercancelRequest() *AlibabalstshipordercancelAPIRequest {
	return &AlibabalstshipordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstshipordercancelAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.shiporder.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstshipordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstshipordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDetailOrderIds is DetailOrderIds Setter
// 需要退款的明细ID
func (r *AlibabalstshipordercancelAPIRequest) SetDetailOrderIds(_detailOrderIds []string) error {
	r._detailOrderIds = _detailOrderIds
	r.Set("detail_order_ids", _detailOrderIds)
	return nil
}

// GetDetailOrderIds DetailOrderIds Getter
func (r AlibabalstshipordercancelAPIRequest) GetDetailOrderIds() []string {
	return r._detailOrderIds
}

// SetReason is Reason Setter
// 取消原因
func (r *AlibabalstshipordercancelAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r AlibabalstshipordercancelAPIRequest) GetReason() string {
	return r._reason
}

// SetOutOrderId is OutOrderId Setter
// 订单号
func (r *AlibabalstshipordercancelAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r AlibabalstshipordercancelAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}
