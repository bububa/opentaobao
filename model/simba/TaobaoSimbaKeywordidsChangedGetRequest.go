package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取修改的词ID APIRequest
taobao.simba.keywordids.changed.get

获取修改的词ID
*/
type TaobaoSimbaKeywordidsChangedGetRequest struct {
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

func NewTaobaoSimbaKeywordidsChangedGetRequest() *TaobaoSimbaKeywordidsChangedGetRequest{
    return &TaobaoSimbaKeywordidsChangedGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaKeywordidsChangedGetRequest) GetApiMethodName() string {
    return "taobao.simba.keywordids.changed.get"
}

func (r TaobaoSimbaKeywordidsChangedGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaKeywordidsChangedGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaKeywordidsChangedGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaKeywordidsChangedGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSimbaKeywordidsChangedGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoSimbaKeywordidsChangedGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaKeywordidsChangedGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSimbaKeywordidsChangedGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoSimbaKeywordidsChangedGetRequest) GetPageNo() int64 {
    return r.pageNo
}

