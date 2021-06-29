package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）关键词更新相关接口 API请求
taobao.simba.keyword.update

（新）关键词更新相关接口
*/
type TaobaoSimbaKeywordUpdateRequest struct {
    model.Params
    // 关键词相关信息
    _bidwords   []SiriusBidwordDTO
}

// 初始化TaobaoSimbaKeywordUpdateRequest对象
func NewTaobaoSimbaKeywordUpdateRequest() *TaobaoSimbaKeywordUpdateRequest{
    return &TaobaoSimbaKeywordUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.keyword.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Bidwords Setter
// 关键词相关信息
func (r *TaobaoSimbaKeywordUpdateRequest) SetBidwords(_bidwords []SiriusBidwordDTO) error {
    r._bidwords = _bidwords
    r.Set("bidwords", _bidwords)
    return nil
}

// Bidwords Getter
func (r TaobaoSimbaKeywordUpdateRequest) GetBidwords() []SiriusBidwordDTO {
    return r._bidwords
}
