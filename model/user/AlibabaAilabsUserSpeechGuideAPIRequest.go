package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsUserSpeechGuideAPIRequest 引导语推荐接口 API请求
// alibaba.ailabs.user.speech.guide
//
// 根据用户的语音query，返回相应的引导语推荐
type AlibabaAilabsUserSpeechGuideAPIRequest struct {
	model.Params
	// 用户query
	_query string
}

// NewAlibabaAilabsUserSpeechGuideRequest 初始化AlibabaAilabsUserSpeechGuideAPIRequest对象
func NewAlibabaAilabsUserSpeechGuideRequest() *AlibabaAilabsUserSpeechGuideAPIRequest {
	return &AlibabaAilabsUserSpeechGuideAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsUserSpeechGuideAPIRequest) Reset() {
	r._query = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsUserSpeechGuideAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.user.speech.guide"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsUserSpeechGuideAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsUserSpeechGuideAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 用户query
func (r *AlibabaAilabsUserSpeechGuideAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabaAilabsUserSpeechGuideAPIRequest) GetQuery() string {
	return r._query
}

var poolAlibabaAilabsUserSpeechGuideAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsUserSpeechGuideRequest()
	},
}

// GetAlibabaAilabsUserSpeechGuideRequest 从 sync.Pool 获取 AlibabaAilabsUserSpeechGuideAPIRequest
func GetAlibabaAilabsUserSpeechGuideAPIRequest() *AlibabaAilabsUserSpeechGuideAPIRequest {
	return poolAlibabaAilabsUserSpeechGuideAPIRequest.Get().(*AlibabaAilabsUserSpeechGuideAPIRequest)
}

// ReleaseAlibabaAilabsUserSpeechGuideAPIRequest 将 AlibabaAilabsUserSpeechGuideAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsUserSpeechGuideAPIRequest(v *AlibabaAilabsUserSpeechGuideAPIRequest) {
	v.Reset()
	poolAlibabaAilabsUserSpeechGuideAPIRequest.Put(v)
}
