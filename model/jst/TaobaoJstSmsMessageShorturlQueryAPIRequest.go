package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmsmessageshorturlqueryAPIRequest 聚石塔短链信息查询 API请求
// taobao.jst.sms.message.shorturl.query
//
// 聚石塔短链信息查询
type TaobaojstsmsmessageshorturlqueryAPIRequest struct {
	model.Params
	// 短链名，即域名后的字符串
	_shortName string
}

// NewTaobaojstsmsmessageshorturlqueryRequest 初始化TaobaojstsmsmessageshorturlqueryAPIRequest对象
func NewTaobaojstsmsmessageshorturlqueryRequest() *TaobaojstsmsmessageshorturlqueryAPIRequest {
	return &TaobaojstsmsmessageshorturlqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstsmsmessageshorturlqueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.message.shorturl.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstsmsmessageshorturlqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstsmsmessageshorturlqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShortName is ShortName Setter
// 短链名，即域名后的字符串
func (r *TaobaojstsmsmessageshorturlqueryAPIRequest) SetShortName(_shortName string) error {
	r._shortName = _shortName
	r.Set("short_name", _shortName)
	return nil
}

// GetShortName ShortName Getter
func (r TaobaojstsmsmessageshorturlqueryAPIRequest) GetShortName() string {
	return r._shortName
}
