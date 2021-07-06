package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallHkOrderClearanceQueryAPIRequest 天猫国际订单清关信息 API请求
// tmall.hk.order.clearance.query
//
// 天猫国际订单清关信息查询
type TmallHkOrderClearanceQueryAPIRequest struct {
	model.Params
	// 调用方业务身份(由国际侧配置提供给调用方)
	_businessSymbol string
	// 交易主订单号
	_bizOrderId int64
}

// NewTmallHkOrderClearanceQueryRequest 初始化TmallHkOrderClearanceQueryAPIRequest对象
func NewTmallHkOrderClearanceQueryRequest() *TmallHkOrderClearanceQueryAPIRequest {
	return &TmallHkOrderClearanceQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallHkOrderClearanceQueryAPIRequest) GetApiMethodName() string {
	return "tmall.hk.order.clearance.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallHkOrderClearanceQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBusinessSymbol is BusinessSymbol Setter
// 调用方业务身份(由国际侧配置提供给调用方)
func (r *TmallHkOrderClearanceQueryAPIRequest) SetBusinessSymbol(_businessSymbol string) error {
	r._businessSymbol = _businessSymbol
	r.Set("business_symbol", _businessSymbol)
	return nil
}

// GetBusinessSymbol BusinessSymbol Getter
func (r TmallHkOrderClearanceQueryAPIRequest) GetBusinessSymbol() string {
	return r._businessSymbol
}

// SetBizOrderId is BizOrderId Setter
// 交易主订单号
func (r *TmallHkOrderClearanceQueryAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TmallHkOrderClearanceQueryAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
