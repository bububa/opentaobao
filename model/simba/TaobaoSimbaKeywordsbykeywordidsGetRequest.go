package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据一个关键词Id列表取得一组关键词 API请求
taobao.simba.keywordsbykeywordids.get

根据一个关键词Id列表取得一组关键词
*/
type TaobaoSimbaKeywordsbykeywordidsGetRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 关键词Id数组，最多200个；
    _keywordIds   []int64
}

// 初始化TaobaoSimbaKeywordsbykeywordidsGetRequest对象
func NewTaobaoSimbaKeywordsbykeywordidsGetRequest() *TaobaoSimbaKeywordsbykeywordidsGetRequest{
    return &TaobaoSimbaKeywordsbykeywordidsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsbykeywordidsGetRequest) GetApiMethodName() string {
    return "taobao.simba.keywordsbykeywordids.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsbykeywordidsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaKeywordsbykeywordidsGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaKeywordsbykeywordidsGetRequest) GetNick() string {
    return r._nick
}
// KeywordIds Setter
// 关键词Id数组，最多200个；
func (r *TaobaoSimbaKeywordsbykeywordidsGetRequest) SetKeywordIds(_keywordIds []int64) error {
    r._keywordIds = _keywordIds
    r.Set("keyword_ids", _keywordIds)
    return nil
}

// KeywordIds Getter
func (r TaobaoSimbaKeywordsbykeywordidsGetRequest) GetKeywordIds() []int64 {
    return r._keywordIds
}
