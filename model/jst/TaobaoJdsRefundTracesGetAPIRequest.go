package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojdsrefundtracesgetAPIRequest 获取单条退款跟踪详情 API请求
// taobao.jds.refund.traces.get
//
// 获取聚石塔数据共享的交易全链路的退款信息
type TaobaojdsrefundtracesgetAPIRequest struct {
	model.Params
	// 淘宝的退款编号
	_refundId int64
}

// NewTaobaojdsrefundtracesgetRequest 初始化TaobaojdsrefundtracesgetAPIRequest对象
func NewTaobaojdsrefundtracesgetRequest() *TaobaojdsrefundtracesgetAPIRequest {
	return &TaobaojdsrefundtracesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojdsrefundtracesgetAPIRequest) GetApiMethodName() string {
	return "taobao.jds.refund.traces.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojdsrefundtracesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojdsrefundtracesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundId is RefundId Setter
// 淘宝的退款编号
func (r *TaobaojdsrefundtracesgetAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaojdsrefundtracesgetAPIRequest) GetRefundId() int64 {
	return r._refundId
}
