package aliospay

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunAliosPayTradeQueryAPIRequest 查询支付结果接口 API请求
// aliyun.alios.pay.trade.query
//
// 商户用来查询支付结果接口
type AliyunAliosPayTradeQueryAPIRequest struct {
	model.Params
	// 请求参数
	_queryTradeRequest *QueryTradeRequest
}

// NewAliyunAliosPayTradeQueryRequest 初始化AliyunAliosPayTradeQueryAPIRequest对象
func NewAliyunAliosPayTradeQueryRequest() *AliyunAliosPayTradeQueryAPIRequest {
	return &AliyunAliosPayTradeQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliyunAliosPayTradeQueryAPIRequest) Reset() {
	r._queryTradeRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunAliosPayTradeQueryAPIRequest) GetApiMethodName() string {
	return "aliyun.alios.pay.trade.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunAliosPayTradeQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunAliosPayTradeQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryTradeRequest is QueryTradeRequest Setter
// 请求参数
func (r *AliyunAliosPayTradeQueryAPIRequest) SetQueryTradeRequest(_queryTradeRequest *QueryTradeRequest) error {
	r._queryTradeRequest = _queryTradeRequest
	r.Set("query_trade_request", _queryTradeRequest)
	return nil
}

// GetQueryTradeRequest QueryTradeRequest Getter
func (r AliyunAliosPayTradeQueryAPIRequest) GetQueryTradeRequest() *QueryTradeRequest {
	return r._queryTradeRequest
}

var poolAliyunAliosPayTradeQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAliyunAliosPayTradeQueryRequest()
	},
}

// GetAliyunAliosPayTradeQueryRequest 从 sync.Pool 获取 AliyunAliosPayTradeQueryAPIRequest
func GetAliyunAliosPayTradeQueryAPIRequest() *AliyunAliosPayTradeQueryAPIRequest {
	return poolAliyunAliosPayTradeQueryAPIRequest.Get().(*AliyunAliosPayTradeQueryAPIRequest)
}

// ReleaseAliyunAliosPayTradeQueryAPIRequest 将 AliyunAliosPayTradeQueryAPIRequest 放入 sync.Pool
func ReleaseAliyunAliosPayTradeQueryAPIRequest(v *AliyunAliosPayTradeQueryAPIRequest) {
	v.Reset()
	poolAliyunAliosPayTradeQueryAPIRequest.Put(v)
}
