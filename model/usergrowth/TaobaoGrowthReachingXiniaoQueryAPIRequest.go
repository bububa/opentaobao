package usergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaogrowthreachingxiniaoqueryAPIRequest 查询溪鸟推荐信息数据 API请求
// taobao.growth.reaching.xiniao.query
//
// 查询溪鸟推荐信息数据
type TaobaogrowthreachingxiniaoqueryAPIRequest struct {
	model.Params
	// 请求参数
	_suggestionContext *XiNiaoSuggestionContextParam
	// 需要的数据量
	_wantedSize int64
}

// NewTaobaogrowthreachingxiniaoqueryRequest 初始化TaobaogrowthreachingxiniaoqueryAPIRequest对象
func NewTaobaogrowthreachingxiniaoqueryRequest() *TaobaogrowthreachingxiniaoqueryAPIRequest {
	return &TaobaogrowthreachingxiniaoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaogrowthreachingxiniaoqueryAPIRequest) GetApiMethodName() string {
	return "taobao.growth.reaching.xiniao.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaogrowthreachingxiniaoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaogrowthreachingxiniaoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSuggestionContext is SuggestionContext Setter
// 请求参数
func (r *TaobaogrowthreachingxiniaoqueryAPIRequest) SetSuggestionContext(_suggestionContext *XiNiaoSuggestionContextParam) error {
	r._suggestionContext = _suggestionContext
	r.Set("suggestion_context", _suggestionContext)
	return nil
}

// GetSuggestionContext SuggestionContext Getter
func (r TaobaogrowthreachingxiniaoqueryAPIRequest) GetSuggestionContext() *XiNiaoSuggestionContextParam {
	return r._suggestionContext
}

// SetWantedSize is WantedSize Setter
// 需要的数据量
func (r *TaobaogrowthreachingxiniaoqueryAPIRequest) SetWantedSize(_wantedSize int64) error {
	r._wantedSize = _wantedSize
	r.Set("wanted_size", _wantedSize)
	return nil
}

// GetWantedSize WantedSize Getter
func (r TaobaogrowthreachingxiniaoqueryAPIRequest) GetWantedSize() int64 {
	return r._wantedSize
}
