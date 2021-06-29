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
    bidwords   []SiriusBidwordDto
    // 推广单元id
    adgroupId   int64
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
func (r *TaobaoSimbaKeywordAddRequest) SetBidwords(bidwords []SiriusBidwordDto) error {
    r.bidwords = bidwords
    r.Set("bidwords", bidwords)
    return nil
}

// Bidwords Getter
func (r TaobaoSimbaKeywordAddRequest) GetBidwords() []SiriusBidwordDto {
    return r.bidwords
}
// AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaKeywordAddRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaKeywordAddRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
