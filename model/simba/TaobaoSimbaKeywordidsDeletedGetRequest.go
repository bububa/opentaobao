package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取删除的词ID APIRequest
taobao.simba.keywordids.deleted.get

获取删除的词ID
*/
type TaobaoSimbaKeywordidsDeletedGetRequest struct {
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

func NewTaobaoSimbaKeywordidsDeletedGetRequest() *TaobaoSimbaKeywordidsDeletedGetRequest{
    return &TaobaoSimbaKeywordidsDeletedGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaKeywordidsDeletedGetRequest) GetApiMethodName() string {
    return "taobao.simba.keywordids.deleted.get"
}

func (r TaobaoSimbaKeywordidsDeletedGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaKeywordidsDeletedGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaKeywordidsDeletedGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaKeywordidsDeletedGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSimbaKeywordidsDeletedGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoSimbaKeywordidsDeletedGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaKeywordidsDeletedGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSimbaKeywordidsDeletedGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoSimbaKeywordidsDeletedGetRequest) GetPageNo() int64 {
    return r.pageNo
}

