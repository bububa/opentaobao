package tbrefund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaorefundmessagesgetAPIRequest 查询退款留言/凭证列表 API请求
// taobao.refund.messages.get
//
// 查询退款留言/凭证列表
type TaobaorefundmessagesgetAPIRequest struct {
	model.Params
	// 需返回的字段列表。可选值：RefundMessage结构体中的所有字段，以半角逗号(,)分隔。
	_fields []string
	// 退款阶段，可选值：onsale（售中），aftersale（售后），天猫退款为必传。
	_refundPhase string
	// 页码
	_pageNo int64
	// 每页条数
	_pageSize int64
	// 退款单号
	_refundId int64
}

// NewTaobaorefundmessagesgetRequest 初始化TaobaorefundmessagesgetAPIRequest对象
func NewTaobaorefundmessagesgetRequest() *TaobaorefundmessagesgetAPIRequest {
	return &TaobaorefundmessagesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaorefundmessagesgetAPIRequest) GetApiMethodName() string {
	return "taobao.refund.messages.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaorefundmessagesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaorefundmessagesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需返回的字段列表。可选值：RefundMessage结构体中的所有字段，以半角逗号(,)分隔。
func (r *TaobaorefundmessagesgetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaorefundmessagesgetAPIRequest) GetFields() []string {
	return r._fields
}

// SetRefundPhase is RefundPhase Setter
// 退款阶段，可选值：onsale（售中），aftersale（售后），天猫退款为必传。
func (r *TaobaorefundmessagesgetAPIRequest) SetRefundPhase(_refundPhase string) error {
	r._refundPhase = _refundPhase
	r.Set("refund_phase", _refundPhase)
	return nil
}

// GetRefundPhase RefundPhase Getter
func (r TaobaorefundmessagesgetAPIRequest) GetRefundPhase() string {
	return r._refundPhase
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaorefundmessagesgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaorefundmessagesgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *TaobaorefundmessagesgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaorefundmessagesgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetRefundId is RefundId Setter
// 退款单号
func (r *TaobaorefundmessagesgetAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaorefundmessagesgetAPIRequest) GetRefundId() int64 {
	return r._refundId
}
