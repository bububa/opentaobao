package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）关键词新增接口 API请求
taobao.simba.keyword.add

（新）关键词更新相关接口
*/
type TaobaoSimbaKeywordAddRequest struct {
    model.Params
    // 关键词相关信息
    _bidwords   []SiriusBidwordDTO
    // 推广单元id
    _adgroupId   int64
}

// 初始化TaobaoSimbaKeywordAddRequest对象
func NewTaobaoSimbaKeywordAddRequest() *TaobaoSimbaKeywordAddRequest{
    return &TaobaoSimbaKeywordAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordAddRequest) GetApiMethodName() string {
    return "taobao.simba.keyword.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Bidwords Setter
// 关键词相关信息
func (r *TaobaoSimbaKeywordAddRequest) SetBidwords(_bidwords []SiriusBidwordDTO) error {
    r._bidwords = _bidwords
    r.Set("bidwords", _bidwords)
    return nil
}

// Bidwords Getter
func (r TaobaoSimbaKeywordAddRequest) GetBidwords() []SiriusBidwordDTO {
    return r._bidwords
}
// AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaKeywordAddRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaKeywordAddRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
