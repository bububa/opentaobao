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
type TaobaoSimbaKeywordUpdateAPIRequest struct {
    model.Params
    // 关键词相关信息
    _bidwords   []SiriusBidwordDTO
}

// 初始化TaobaoSimbaKeywordUpdateAPIRequest对象
func NewTaobaoSimbaKeywordUpdateRequest() *TaobaoSimbaKeywordUpdateAPIRequest{
    return &TaobaoSimbaKeywordUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.simba.keyword.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Bidwords Setter
// 关键词相关信息
func (r *TaobaoSimbaKeywordUpdateAPIRequest) SetBidwords(_bidwords []SiriusBidwordDTO) error {
    r._bidwords = _bidwords
    r.Set("bidwords", _bidwords)
    return nil
}

// Bidwords Getter
func (r TaobaoSimbaKeywordUpdateAPIRequest) GetBidwords() []SiriusBidwordDTO {
    return r._bidwords
}
