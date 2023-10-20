package trade

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkPosTradeCloseAPIRequest) Reset() {
	r._closeRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkPosTradeCloseAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.pos.trade.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkPosTradeCloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkPosTradeCloseAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaWdkPosTradeCloseAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkPosTradeCloseRequest()
	},
}

// GetAlibabaWdkPosTradeCloseRequest 从 sync.Pool 获取 AlibabaWdkPosTradeCloseAPIRequest
func GetAlibabaWdkPosTradeCloseAPIRequest() *AlibabaWdkPosTradeCloseAPIRequest {
	return poolAlibabaWdkPosTradeCloseAPIRequest.Get().(*AlibabaWdkPosTradeCloseAPIRequest)
}

// ReleaseAlibabaWdkPosTradeCloseAPIRequest 将 AlibabaWdkPosTradeCloseAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkPosTradeCloseAPIRequest(v *AlibabaWdkPosTradeCloseAPIRequest) {
	v.Reset()
	poolAlibabaWdkPosTradeCloseAPIRequest.Put(v)
}
