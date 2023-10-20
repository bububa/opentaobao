package usergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaogrowthreachingsuggestionqueryAPIRequest 厂商生态推荐信息查询 API请求
// taobao.growth.reaching.suggestion.query
//
// 厂商生态推荐信息查询
type TaobaogrowthreachingsuggestionqueryAPIRequest struct {
	model.Params
	// 请求参数
	_suggestionContext *SuggestionContextParam
	// 需要的数据量
	_wantedSize int64
}

// NewTaobaogrowthreachingsuggestionqueryRequest 初始化TaobaogrowthreachingsuggestionqueryAPIRequest对象
func NewTaobaogrowthreachingsuggestionqueryRequest() *TaobaogrowthreachingsuggestionqueryAPIRequest {
	return &TaobaogrowthreachingsuggestionqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaogrowthreachingsuggestionqueryAPIRequest) GetApiMethodName() string {
	return "taobao.growth.reaching.suggestion.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaogrowthreachingsuggestionqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaogrowthreachingsuggestionqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSuggestionContext is SuggestionContext Setter
// 请求参数
func (r *TaobaogrowthreachingsuggestionqueryAPIRequest) SetSuggestionContext(_suggestionContext *SuggestionContextParam) error {
	r._suggestionContext = _suggestionContext
	r.Set("suggestion_context", _suggestionContext)
	return nil
}

// GetSuggestionContext SuggestionContext Getter
func (r TaobaogrowthreachingsuggestionqueryAPIRequest) GetSuggestionContext() *SuggestionContextParam {
	return r._suggestionContext
}

// SetWantedSize is WantedSize Setter
// 需要的数据量
func (r *TaobaogrowthreachingsuggestionqueryAPIRequest) SetWantedSize(_wantedSize int64) error {
	r._wantedSize = _wantedSize
	r.Set("wanted_size", _wantedSize)
	return nil
}

// GetWantedSize WantedSize Getter
func (r TaobaogrowthreachingsuggestionqueryAPIRequest) GetWantedSize() int64 {
	return r._wantedSize
}
