package fundplatform

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformCardorderFetchCardAsyncAPIRequest 异步批量生成储值卡 API请求
// alibaba.fundplatform.cardorder.fetch.card.async
//
// 外部业务方异步批量生成储值卡的接口。同步只返回接受成功，异步会通知制卡成功
type AlibabaFundplatformCardorderFetchCardAsyncAPIRequest struct {
	model.Params
	// 入参复杂对象
	_paramCardFetchAsyncRequest *CardFetchAsyncRequest
}

// NewAlibabaFundplatformCardorderFetchCardAsyncRequest 初始化AlibabaFundplatformCardorderFetchCardAsyncAPIRequest对象
func NewAlibabaFundplatformCardorderFetchCardAsyncRequest() *AlibabaFundplatformCardorderFetchCardAsyncAPIRequest {
	return &AlibabaFundplatformCardorderFetchCardAsyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaFundplatformCardorderFetchCardAsyncAPIRequest) Reset() {
	r._paramCardFetchAsyncRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardorderFetchCardAsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.cardorder.fetch.card.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardorderFetchCardAsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaFundplatformCardorderFetchCardAsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamCardFetchAsyncRequest is ParamCardFetchAsyncRequest Setter
// 入参复杂对象
func (r *AlibabaFundplatformCardorderFetchCardAsyncAPIRequest) SetParamCardFetchAsyncRequest(_paramCardFetchAsyncRequest *CardFetchAsyncRequest) error {
	r._paramCardFetchAsyncRequest = _paramCardFetchAsyncRequest
	r.Set("param_card_fetch_async_request", _paramCardFetchAsyncRequest)
	return nil
}

// GetParamCardFetchAsyncRequest ParamCardFetchAsyncRequest Getter
func (r AlibabaFundplatformCardorderFetchCardAsyncAPIRequest) GetParamCardFetchAsyncRequest() *CardFetchAsyncRequest {
	return r._paramCardFetchAsyncRequest
}

var poolAlibabaFundplatformCardorderFetchCardAsyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaFundplatformCardorderFetchCardAsyncRequest()
	},
}

// GetAlibabaFundplatformCardorderFetchCardAsyncRequest 从 sync.Pool 获取 AlibabaFundplatformCardorderFetchCardAsyncAPIRequest
func GetAlibabaFundplatformCardorderFetchCardAsyncAPIRequest() *AlibabaFundplatformCardorderFetchCardAsyncAPIRequest {
	return poolAlibabaFundplatformCardorderFetchCardAsyncAPIRequest.Get().(*AlibabaFundplatformCardorderFetchCardAsyncAPIRequest)
}

// ReleaseAlibabaFundplatformCardorderFetchCardAsyncAPIRequest 将 AlibabaFundplatformCardorderFetchCardAsyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaFundplatformCardorderFetchCardAsyncAPIRequest(v *AlibabaFundplatformCardorderFetchCardAsyncAPIRequest) {
	v.Reset()
	poolAlibabaFundplatformCardorderFetchCardAsyncAPIRequest.Put(v)
}
