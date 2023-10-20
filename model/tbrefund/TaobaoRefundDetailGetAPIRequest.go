package tbrefund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaorefunddetailgetAPIRequest 退款详情页渲染 API请求
// taobao.refund.detail.get
//
// 退款详情页渲染
type TaobaorefunddetailgetAPIRequest struct {
	model.Params
	// 需要返回的字段。目前支持有：allowedOperations,refund_version
	_fields string
	// 退款编号
	_refundId int64
}

// NewTaobaorefunddetailgetRequest 初始化TaobaorefunddetailgetAPIRequest对象
func NewTaobaorefunddetailgetRequest() *TaobaorefunddetailgetAPIRequest {
	return &TaobaorefunddetailgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaorefunddetailgetAPIRequest) GetApiMethodName() string {
	return "taobao.refund.detail.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaorefunddetailgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaorefunddetailgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段。目前支持有：allowedOperations,refund_version
func (r *TaobaorefunddetailgetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaorefunddetailgetAPIRequest) GetFields() string {
	return r._fields
}

// SetRefundId is RefundId Setter
// 退款编号
func (r *TaobaorefunddetailgetAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaorefunddetailgetAPIRequest) GetRefundId() int64 {
	return r._refundId
}
