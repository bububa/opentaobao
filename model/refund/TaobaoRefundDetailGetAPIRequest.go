package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRefundDetailGetAPIRequest 退款详情页渲染 API请求
// taobao.refund.detail.get
//
// 退款详情页渲染
type TaobaoRefundDetailGetAPIRequest struct {
	model.Params
	// 需要返回的字段。目前支持有：allowedOperations,refund_version
	_fields string
	// 退款编号
	_refundId int64
}

// NewTaobaoRefundDetailGetRequest 初始化TaobaoRefundDetailGetAPIRequest对象
func NewTaobaoRefundDetailGetRequest() *TaobaoRefundDetailGetAPIRequest {
	return &TaobaoRefundDetailGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRefundDetailGetAPIRequest) GetApiMethodName() string {
	return "taobao.refund.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRefundDetailGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFields is Fields Setter
// 需要返回的字段。目前支持有：allowedOperations,refund_version
func (r *TaobaoRefundDetailGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoRefundDetailGetAPIRequest) GetFields() string {
	return r._fields
}

// SetRefundId is RefundId Setter
// 退款编号
func (r *TaobaoRefundDetailGetAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaoRefundDetailGetAPIRequest) GetRefundId() int64 {
	return r._refundId
}
