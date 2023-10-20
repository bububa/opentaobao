package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsMessageShorturlQueryAPIRequest 聚石塔短链信息查询 API请求
// taobao.jst.sms.message.shorturl.query
//
// 聚石塔短链信息查询
type TaobaoJstSmsMessageShorturlQueryAPIRequest struct {
	model.Params
	// 短链名，即域名后的字符串
	_shortName string
}

// NewTaobaoJstSmsMessageShorturlQueryRequest 初始化TaobaoJstSmsMessageShorturlQueryAPIRequest对象
func NewTaobaoJstSmsMessageShorturlQueryRequest() *TaobaoJstSmsMessageShorturlQueryAPIRequest {
	return &TaobaoJstSmsMessageShorturlQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstSmsMessageShorturlQueryAPIRequest) Reset() {
	r._shortName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsMessageShorturlQueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.message.shorturl.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsMessageShorturlQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstSmsMessageShorturlQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShortName is ShortName Setter
// 短链名，即域名后的字符串
func (r *TaobaoJstSmsMessageShorturlQueryAPIRequest) SetShortName(_shortName string) error {
	r._shortName = _shortName
	r.Set("short_name", _shortName)
	return nil
}

// GetShortName ShortName Getter
func (r TaobaoJstSmsMessageShorturlQueryAPIRequest) GetShortName() string {
	return r._shortName
}

var poolTaobaoJstSmsMessageShorturlQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstSmsMessageShorturlQueryRequest()
	},
}

// GetTaobaoJstSmsMessageShorturlQueryRequest 从 sync.Pool 获取 TaobaoJstSmsMessageShorturlQueryAPIRequest
func GetTaobaoJstSmsMessageShorturlQueryAPIRequest() *TaobaoJstSmsMessageShorturlQueryAPIRequest {
	return poolTaobaoJstSmsMessageShorturlQueryAPIRequest.Get().(*TaobaoJstSmsMessageShorturlQueryAPIRequest)
}

// ReleaseTaobaoJstSmsMessageShorturlQueryAPIRequest 将 TaobaoJstSmsMessageShorturlQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstSmsMessageShorturlQueryAPIRequest(v *TaobaoJstSmsMessageShorturlQueryAPIRequest) {
	v.Reset()
	poolTaobaoJstSmsMessageShorturlQueryAPIRequest.Put(v)
}
