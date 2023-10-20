package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuottpayorderqueryauthstateAPIRequest 查询连包签约状态 API请求
// youku.ott.pay.order.queryauthstate
//
// 查询CP用户连包商品签约状态
type YoukuottpayorderqueryauthstateAPIRequest struct {
	model.Params
	// 原始签约订单号
	_originalCpOrderNo string
}

// NewYoukuottpayorderqueryauthstateRequest 初始化YoukuottpayorderqueryauthstateAPIRequest对象
func NewYoukuottpayorderqueryauthstateRequest() *YoukuottpayorderqueryauthstateAPIRequest {
	return &YoukuottpayorderqueryauthstateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuottpayorderqueryauthstateAPIRequest) GetApiMethodName() string {
	return "youku.ott.pay.order.queryauthstate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuottpayorderqueryauthstateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuottpayorderqueryauthstateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOriginalCpOrderNo is OriginalCpOrderNo Setter
// 原始签约订单号
func (r *YoukuottpayorderqueryauthstateAPIRequest) SetOriginalCpOrderNo(_originalCpOrderNo string) error {
	r._originalCpOrderNo = _originalCpOrderNo
	r.Set("original_cp_order_no", _originalCpOrderNo)
	return nil
}

// GetOriginalCpOrderNo OriginalCpOrderNo Getter
func (r YoukuottpayorderqueryauthstateAPIRequest) GetOriginalCpOrderNo() string {
	return r._originalCpOrderNo
}
