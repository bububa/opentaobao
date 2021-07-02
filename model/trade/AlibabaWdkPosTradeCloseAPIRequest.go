package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkPosTradeCloseAPIRequest 轻pos品牌营销关单接口 API请求
// alibaba.wdk.pos.trade.close
//
// 轻pos品牌营销场景，提供关单接口给外部商家
type AlibabaWdkPosTradeCloseAPIRequest struct {
	model.Params
	// 关单请求
	_closeRequest *FastBuyPosCloseRequest
}

// NewAlibabaWdkPosTradeCloseRequest 初始化AlibabaWdkPosTradeCloseAPIRequest对象
func NewAlibabaWdkPosTradeCloseRequest() *AlibabaWdkPosTradeCloseAPIRequest {
	return &AlibabaWdkPosTradeCloseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkPosTradeCloseAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.pos.trade.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkPosTradeCloseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCloseRequest is CloseRequest Setter
// 关单请求
func (r *AlibabaWdkPosTradeCloseAPIRequest) SetCloseRequest(_closeRequest *FastBuyPosCloseRequest) error {
	r._closeRequest = _closeRequest
	r.Set("close_request", _closeRequest)
	return nil
}

// GetCloseRequest CloseRequest Getter
func (r AlibabaWdkPosTradeCloseAPIRequest) GetCloseRequest() *FastBuyPosCloseRequest {
	return r._closeRequest
}
