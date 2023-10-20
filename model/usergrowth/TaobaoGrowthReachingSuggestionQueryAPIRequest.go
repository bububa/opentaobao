package usergrowth

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGrowthReachingSuggestionQueryAPIRequest 厂商生态推荐信息查询 API请求
// taobao.growth.reaching.suggestion.query
//
// 厂商生态推荐信息查询
type TaobaoGrowthReachingSuggestionQueryAPIRequest struct {
	model.Params
	// 请求参数
	_suggestionContext *SuggestionContextParam
	// 需要的数据量
	_wantedSize int64
}

// NewTaobaoGrowthReachingSuggestionQueryRequest 初始化TaobaoGrowthReachingSuggestionQueryAPIRequest对象
func NewTaobaoGrowthReachingSuggestionQueryRequest() *TaobaoGrowthReachingSuggestionQueryAPIRequest {
	return &TaobaoGrowthReachingSuggestionQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoGrowthReachingSuggestionQueryAPIRequest) Reset() {
	r._suggestionContext = nil
	r._wantedSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoGrowthReachingSuggestionQueryAPIRequest) GetApiMethodName() string {
	return "taobao.growth.reaching.suggestion.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoGrowthReachingSuggestionQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoGrowthReachingSuggestionQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSuggestionContext is SuggestionContext Setter
// 请求参数
func (r *TaobaoGrowthReachingSuggestionQueryAPIRequest) SetSuggestionContext(_suggestionContext *SuggestionContextParam) error {
	r._suggestionContext = _suggestionContext
	r.Set("suggestion_context", _suggestionContext)
	return nil
}

// GetSuggestionContext SuggestionContext Getter
func (r TaobaoGrowthReachingSuggestionQueryAPIRequest) GetSuggestionContext() *SuggestionContextParam {
	return r._suggestionContext
}

// SetWantedSize is WantedSize Setter
// 需要的数据量
func (r *TaobaoGrowthReachingSuggestionQueryAPIRequest) SetWantedSize(_wantedSize int64) error {
	r._wantedSize = _wantedSize
	r.Set("wanted_size", _wantedSize)
	return nil
}

// GetWantedSize WantedSize Getter
func (r TaobaoGrowthReachingSuggestionQueryAPIRequest) GetWantedSize() int64 {
	return r._wantedSize
}

var poolTaobaoGrowthReachingSuggestionQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoGrowthReachingSuggestionQueryRequest()
	},
}

// GetTaobaoGrowthReachingSuggestionQueryRequest 从 sync.Pool 获取 TaobaoGrowthReachingSuggestionQueryAPIRequest
func GetTaobaoGrowthReachingSuggestionQueryAPIRequest() *TaobaoGrowthReachingSuggestionQueryAPIRequest {
	return poolTaobaoGrowthReachingSuggestionQueryAPIRequest.Get().(*TaobaoGrowthReachingSuggestionQueryAPIRequest)
}

// ReleaseTaobaoGrowthReachingSuggestionQueryAPIRequest 将 TaobaoGrowthReachingSuggestionQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoGrowthReachingSuggestionQueryAPIRequest(v *TaobaoGrowthReachingSuggestionQueryAPIRequest) {
	v.Reset()
	poolTaobaoGrowthReachingSuggestionQueryAPIRequest.Put(v)
}
