package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取删除的推广组ID APIRequest
taobao.simba.adgroupids.deleted.get

获取删除的推广组ID
*/
type TaobaoSimbaAdgroupidsDeletedGetRequest struct {
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

func NewTaobaoSimbaAdgroupidsDeletedGetRequest() *TaobaoSimbaAdgroupidsDeletedGetRequest{
    return &TaobaoSimbaAdgroupidsDeletedGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaAdgroupidsDeletedGetRequest) GetApiMethodName() string {
    return "taobao.simba.adgroupids.deleted.get"
}

func (r TaobaoSimbaAdgroupidsDeletedGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaAdgroupidsDeletedGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaAdgroupidsDeletedGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaAdgroupidsDeletedGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSimbaAdgroupidsDeletedGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoSimbaAdgroupidsDeletedGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaAdgroupidsDeletedGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSimbaAdgroupidsDeletedGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoSimbaAdgroupidsDeletedGetRequest) GetPageNo() int64 {
    return r.pageNo
}

