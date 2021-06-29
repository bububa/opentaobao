package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销量明星关键词删除 API请求
taobao.simba.salestar.keywords.delete

（新）关键词删除相关接口
*/
type TaobaoSimbaSalestarKeywordsDeleteRequest struct {
    model.Params
    // 关键词ids
    _bidwordIds   []int64
}

// 初始化TaobaoSimbaSalestarKeywordsDeleteRequest对象
func NewTaobaoSimbaSalestarKeywordsDeleteRequest() *TaobaoSimbaSalestarKeywordsDeleteRequest{
    return &TaobaoSimbaSalestarKeywordsDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarKeywordsDeleteRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.keywords.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarKeywordsDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BidwordIds Setter
// 关键词ids
func (r *TaobaoSimbaSalestarKeywordsDeleteRequest) SetBidwordIds(_bidwordIds []int64) error {
    r._bidwordIds = _bidwordIds
    r.Set("bidword_ids", _bidwordIds)
    return nil
}

// BidwordIds Getter
func (r TaobaoSimbaSalestarKeywordsDeleteRequest) GetBidwordIds() []int64 {
    return r._bidwordIds
}
