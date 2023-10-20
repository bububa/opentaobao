package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordUpdateAPIRequest （新）关键词更新相关接口 API请求
// taobao.simba.keyword.update
//
// （新）关键词更新相关接口
type TaobaoSimbaKeywordUpdateAPIRequest struct {
	model.Params
	// 关键词相关信息
	_bidwords []SiriusBidwordDto
}

// NewTaobaoSimbaKeywordUpdateRequest 初始化TaobaoSimbaKeywordUpdateAPIRequest对象
func NewTaobaoSimbaKeywordUpdateRequest() *TaobaoSimbaKeywordUpdateAPIRequest {
	return &TaobaoSimbaKeywordUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaKeywordUpdateAPIRequest) Reset() {
	r._bidwords = r._bidwords[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keyword.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaKeywordUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidwords is Bidwords Setter
// 关键词相关信息
func (r *TaobaoSimbaKeywordUpdateAPIRequest) SetBidwords(_bidwords []SiriusBidwordDto) error {
	r._bidwords = _bidwords
	r.Set("bidwords", _bidwords)
	return nil
}

// GetBidwords Bidwords Getter
func (r TaobaoSimbaKeywordUpdateAPIRequest) GetBidwords() []SiriusBidwordDto {
	return r._bidwords
}

var poolTaobaoSimbaKeywordUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaKeywordUpdateRequest()
	},
}

// GetTaobaoSimbaKeywordUpdateRequest 从 sync.Pool 获取 TaobaoSimbaKeywordUpdateAPIRequest
func GetTaobaoSimbaKeywordUpdateAPIRequest() *TaobaoSimbaKeywordUpdateAPIRequest {
	return poolTaobaoSimbaKeywordUpdateAPIRequest.Get().(*TaobaoSimbaKeywordUpdateAPIRequest)
}

// ReleaseTaobaoSimbaKeywordUpdateAPIRequest 将 TaobaoSimbaKeywordUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaKeywordUpdateAPIRequest(v *TaobaoSimbaKeywordUpdateAPIRequest) {
	v.Reset()
	poolTaobaoSimbaKeywordUpdateAPIRequest.Put(v)
}
