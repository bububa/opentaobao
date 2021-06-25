package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页获取修改过的关键词ID、宝贝id、修改时间 APIRequest
taobao.simba.keywords.changed.get

分页获取修改过的关键词ID、宝贝id、修改时间
*/
type TaobaoSimbaKeywordsChangedGetRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 得到此时间点之后的数据，不能大于一个月
    startTime   string 

    // 返回的每页数据量大小,默认200最大1000
    pageSize   int64 

    // 返回的第几页数据，默认为1
    pageNo   int64 

}

func NewTaobaoSimbaKeywordsChangedGetRequest() *TaobaoSimbaKeywordsChangedGetRequest{
    return &TaobaoSimbaKeywordsChangedGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaKeywordsChangedGetRequest) GetApiMethodName() string {
    return "taobao.simba.keywords.changed.get"
}

func (r TaobaoSimbaKeywordsChangedGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaKeywordsChangedGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaKeywordsChangedGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaKeywordsChangedGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSimbaKeywordsChangedGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoSimbaKeywordsChangedGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaKeywordsChangedGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSimbaKeywordsChangedGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoSimbaKeywordsChangedGetRequest) GetPageNo() int64 {
    return r.pageNo
}

