package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJdsRefundTracesGetAPIRequest 获取单条退款跟踪详情 API请求
// taobao.jds.refund.traces.get
//
// 获取聚石塔数据共享的交易全链路的退款信息
type TaobaoJdsRefundTracesGetAPIRequest struct {
	model.Params
	// 淘宝的退款编号
	_refundId int64
}

// NewTaobaoJdsRefundTracesGetRequest 初始化TaobaoJdsRefundTracesGetAPIRequest对象
func NewTaobaoJdsRefundTracesGetRequest() *TaobaoJdsRefundTracesGetAPIRequest {
	return &TaobaoJdsRefundTracesGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJdsRefundTracesGetAPIRequest) GetApiMethodName() string {
	return "taobao.jds.refund.traces.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJdsRefundTracesGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJdsRefundTracesGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundId is RefundId Setter
// 淘宝的退款编号
func (r *TaobaoJdsRefundTracesGetAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoJdsRefundTracesGetAPIRequest) GetRefundId() int64 {
	return r._refundId
}
