package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttPayOrderQueryauthstateAPIRequest 查询连包签约状态 API请求
// youku.ott.pay.order.queryauthstate
//
// 查询CP用户连包商品签约状态
type YoukuOttPayOrderQueryauthstateAPIRequest struct {
	model.Params
	// 原始签约订单号
	_originalCpOrderNo string
}

// NewYoukuOttPayOrderQueryauthstateRequest 初始化YoukuOttPayOrderQueryauthstateAPIRequest对象
func NewYoukuOttPayOrderQueryauthstateRequest() *YoukuOttPayOrderQueryauthstateAPIRequest {
	return &YoukuOttPayOrderQueryauthstateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttPayOrderQueryauthstateAPIRequest) GetApiMethodName() string {
	return "youku.ott.pay.order.queryauthstate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttPayOrderQueryauthstateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OriginalCpOrderNo Setter
// 原始签约订单号
func (r *YoukuOttPayOrderQueryauthstateAPIRequest) SetOriginalCpOrderNo(_originalCpOrderNo string) error {
	r._originalCpOrderNo = _originalCpOrderNo
	r.Set("original_cp_order_no", _originalCpOrderNo)
	return nil
}

// Get OriginalCpOrderNo Getter
func (r YoukuOttPayOrderQueryauthstateAPIRequest) GetOriginalCpOrderNo() string {
	return r._originalCpOrderNo
}
