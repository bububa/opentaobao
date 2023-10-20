package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabafundplatformcardorderfetchcardasyncAPIRequest 异步批量生成储值卡 API请求
// alibaba.fundplatform.cardorder.fetch.card.async
//
// 外部业务方异步批量生成储值卡的接口。同步只返回接受成功，异步会通知制卡成功
type AlibabafundplatformcardorderfetchcardasyncAPIRequest struct {
	model.Params
	// 入参复杂对象
	_paramCardFetchAsyncRequest *CardFetchAsyncRequest
}

// NewAlibabafundplatformcardorderfetchcardasyncRequest 初始化AlibabafundplatformcardorderfetchcardasyncAPIRequest对象
func NewAlibabafundplatformcardorderfetchcardasyncRequest() *AlibabafundplatformcardorderfetchcardasyncAPIRequest {
	return &AlibabafundplatformcardorderfetchcardasyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabafundplatformcardorderfetchcardasyncAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.cardorder.fetch.card.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabafundplatformcardorderfetchcardasyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabafundplatformcardorderfetchcardasyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCardFetchAsyncRequest is ParamCardFetchAsyncRequest Setter
// 入参复杂对象
func (r *AlibabafundplatformcardorderfetchcardasyncAPIRequest) SetParamCardFetchAsyncRequest(_paramCardFetchAsyncRequest *CardFetchAsyncRequest) error {
	r._paramCardFetchAsyncRequest = _paramCardFetchAsyncRequest
	r.Set("param_card_fetch_async_request", _paramCardFetchAsyncRequest)
	return nil
}

// GetParamCardFetchAsyncRequest ParamCardFetchAsyncRequest Getter
func (r AlibabafundplatformcardorderfetchcardasyncAPIRequest) GetParamCardFetchAsyncRequest() *CardFetchAsyncRequest {
	return r._paramCardFetchAsyncRequest
}
