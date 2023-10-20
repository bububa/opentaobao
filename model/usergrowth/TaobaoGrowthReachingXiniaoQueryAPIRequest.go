package usergrowth

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGrowthReachingXiniaoQueryAPIRequest 查询溪鸟推荐信息数据 API请求
// taobao.growth.reaching.xiniao.query
//
// 查询溪鸟推荐信息数据
type TaobaoGrowthReachingXiniaoQueryAPIRequest struct {
	model.Params
	// 请求参数
	_suggestionContext *XiNiaoSuggestionContextParam
	// 需要的数据量
	_wantedSize int64
}

// NewTaobaoGrowthReachingXiniaoQueryRequest 初始化TaobaoGrowthReachingXiniaoQueryAPIRequest对象
func NewTaobaoGrowthReachingXiniaoQueryRequest() *TaobaoGrowthReachingXiniaoQueryAPIRequest {
	return &TaobaoGrowthReachingXiniaoQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoGrowthReachingXiniaoQueryAPIRequest) Reset() {
	r._suggestionContext = nil
	r._wantedSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoGrowthReachingXiniaoQueryAPIRequest) GetApiMethodName() string {
	return "taobao.growth.reaching.xiniao.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoGrowthReachingXiniaoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoGrowthReachingXiniaoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSuggestionContext is SuggestionContext Setter
// 请求参数
func (r *TaobaoGrowthReachingXiniaoQueryAPIRequest) SetSuggestionContext(_suggestionContext *XiNiaoSuggestionContextParam) error {
	r._suggestionContext = _suggestionContext
	r.Set("suggestion_context", _suggestionContext)
	return nil
}

// GetSuggestionContext SuggestionContext Getter
func (r TaobaoGrowthReachingXiniaoQueryAPIRequest) GetSuggestionContext() *XiNiaoSuggestionContextParam {
	return r._suggestionContext
}

// SetWantedSize is WantedSize Setter
// 需要的数据量
func (r *TaobaoGrowthReachingXiniaoQueryAPIRequest) SetWantedSize(_wantedSize int64) error {
	r._wantedSize = _wantedSize
	r.Set("wanted_size", _wantedSize)
	return nil
}

// GetWantedSize WantedSize Getter
func (r TaobaoGrowthReachingXiniaoQueryAPIRequest) GetWantedSize() int64 {
	return r._wantedSize
}

var poolTaobaoGrowthReachingXiniaoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoGrowthReachingXiniaoQueryRequest()
	},
}

// GetTaobaoGrowthReachingXiniaoQueryRequest 从 sync.Pool 获取 TaobaoGrowthReachingXiniaoQueryAPIRequest
func GetTaobaoGrowthReachingXiniaoQueryAPIRequest() *TaobaoGrowthReachingXiniaoQueryAPIRequest {
	return poolTaobaoGrowthReachingXiniaoQueryAPIRequest.Get().(*TaobaoGrowthReachingXiniaoQueryAPIRequest)
}

// ReleaseTaobaoGrowthReachingXiniaoQueryAPIRequest 将 TaobaoGrowthReachingXiniaoQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoGrowthReachingXiniaoQueryAPIRequest(v *TaobaoGrowthReachingXiniaoQueryAPIRequest) {
	v.Reset()
	poolTaobaoGrowthReachingXiniaoQueryAPIRequest.Put(v)
}
