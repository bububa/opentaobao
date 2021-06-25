package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据一个关键词Id列表取得一组关键词 APIRequest
taobao.simba.keywordsbykeywordids.get

根据一个关键词Id列表取得一组关键词
*/
type TaobaoSimbaKeywordsbykeywordidsGetRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 关键词Id数组，最多200个；
    keywordIds   []Number 

}

func NewTaobaoSimbaKeywordsbykeywordidsGetRequest() *TaobaoSimbaKeywordsbykeywordidsGetRequest{
    return &TaobaoSimbaKeywordsbykeywordidsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaKeywordsbykeywordidsGetRequest) GetApiMethodName() string {
    return "taobao.simba.keywordsbykeywordids.get"
}

func (r TaobaoSimbaKeywordsbykeywordidsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaKeywordsbykeywordidsGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaKeywordsbykeywordidsGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaKeywordsbykeywordidsGetRequest) SetKeywordIds(keywordIds []Number) error {
    r.keywordIds = keywordIds
    r.Set("keyword_ids", keywordIds)
    return nil
}

func (r TaobaoSimbaKeywordsbykeywordidsGetRequest) GetKeywordIds() []Number {
    return r.keywordIds
}

