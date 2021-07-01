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
type TaobaoSimbaKeywordAddAPIRequest struct {
    model.Params
    // 关键词相关信息
    _bidwords   []SiriusBidwordDto
    // 推广单元id
    _adgroupId   int64
}

// 初始化TaobaoSimbaKeywordAddAPIRequest对象
func NewTaobaoSimbaKeywordAddRequest() *TaobaoSimbaKeywordAddAPIRequest{
    return &TaobaoSimbaKeywordAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordAddAPIRequest) GetApiMethodName() string {
    return "taobao.simba.keyword.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Bidwords Setter
// 关键词相关信息
func (r *TaobaoSimbaKeywordAddAPIRequest) SetBidwords(_bidwords []SiriusBidwordDto) error {
    r._bidwords = _bidwords
    r.Set("bidwords", _bidwords)
    return nil
}

// Bidwords Getter
func (r TaobaoSimbaKeywordAddAPIRequest) GetBidwords() []SiriusBidwordDto {
    return r._bidwords
}
// AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaKeywordAddAPIRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaKeywordAddAPIRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
