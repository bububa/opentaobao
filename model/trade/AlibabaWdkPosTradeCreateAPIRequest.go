package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkPosTradeCreateAPIRequest 轻pos品牌营销下单接口 API请求
// alibaba.wdk.pos.trade.create
//
// 提供给石基进行轻pos品牌营销下单
type AlibabaWdkPosTradeCreateAPIRequest struct {
	model.Params
	// 下单请求
	_createRequest *FastBuyPosCreateRequest
}

// NewAlibabaWdkPosTradeCreateRequest 初始化AlibabaWdkPosTradeCreateAPIRequest对象
func NewAlibabaWdkPosTradeCreateRequest() *AlibabaWdkPosTradeCreateAPIRequest {
	return &AlibabaWdkPosTradeCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkPosTradeCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.pos.trade.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkPosTradeCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CreateRequest Setter
// 下单请求
func (r *AlibabaWdkPosTradeCreateAPIRequest) SetCreateRequest(_createRequest *FastBuyPosCreateRequest) error {
	r._createRequest = _createRequest
	r.Set("create_request", _createRequest)
	return nil
}

// Get CreateRequest Getter
func (r AlibabaWdkPosTradeCreateAPIRequest) GetCreateRequest() *FastBuyPosCreateRequest {
	return r._createRequest
}
