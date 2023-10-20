package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkpostradecloseAPIRequest 轻pos品牌营销关单接口 API请求
// alibaba.wdk.pos.trade.close
//
// 轻pos品牌营销场景，提供关单接口给外部商家
type AlibabawdkpostradecloseAPIRequest struct {
	model.Params
	// 关单请求
	_closeRequest *FastBuyPosCloseRequest
}

// NewAlibabawdkpostradecloseRequest 初始化AlibabawdkpostradecloseAPIRequest对象
func NewAlibabawdkpostradecloseRequest() *AlibabawdkpostradecloseAPIRequest {
	return &AlibabawdkpostradecloseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkpostradecloseAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.pos.trade.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkpostradecloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkpostradecloseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCloseRequest is CloseRequest Setter
// 关单请求
func (r *AlibabawdkpostradecloseAPIRequest) SetCloseRequest(_closeRequest *FastBuyPosCloseRequest) error {
	r._closeRequest = _closeRequest
	r.Set("close_request", _closeRequest)
	return nil
}

// GetCloseRequest CloseRequest Getter
func (r AlibabawdkpostradecloseAPIRequest) GetCloseRequest() *FastBuyPosCloseRequest {
	return r._closeRequest
}
