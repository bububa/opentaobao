package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsuserspeechguideAPIRequest 引导语推荐接口 API请求
// alibaba.ailabs.user.speech.guide
//
// 根据用户的语音query，返回相应的引导语推荐
type AlibabaailabsuserspeechguideAPIRequest struct {
	model.Params
	// 用户query
	_query string
}

// NewAlibabaailabsuserspeechguideRequest 初始化AlibabaailabsuserspeechguideAPIRequest对象
func NewAlibabaailabsuserspeechguideRequest() *AlibabaailabsuserspeechguideAPIRequest {
	return &AlibabaailabsuserspeechguideAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabsuserspeechguideAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.user.speech.guide"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabsuserspeechguideAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabsuserspeechguideAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 用户query
func (r *AlibabaailabsuserspeechguideAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaailabsuserspeechguideAPIRequest) GetQuery() string {
	return r._query
}
