package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaKeywordUpdateAPIRequest
（新）关键词更新相关接口 API请求
taobao.simba.keyword.update

（新）关键词更新相关接口 */
type TaobaoSimbaKeywordUpdateAPIRequest struct {
	model.Params
	// 关键词相关信息
	_bidwords []SiriusBidwordDto
}

// NewTaobaoSimbaKeywordUpdateRequest 初始化TaobaoSimbaKeywordUpdateAPIRequest对象
func NewTaobaoSimbaKeywordUpdateRequest() *TaobaoSimbaKeywordUpdateAPIRequest {
	return &TaobaoSimbaKeywordUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keyword.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Bidwords Setter
// 关键词相关信息
func (r *TaobaoSimbaKeywordUpdateAPIRequest) SetBidwords(_bidwords []SiriusBidwordDto) error {
	r._bidwords = _bidwords
	r.Set("bidwords", _bidwords)
	return nil
}

// Get Bidwords Getter
func (r TaobaoSimbaKeywordUpdateAPIRequest) GetBidwords() []SiriusBidwordDto {
	return r._bidwords
}
