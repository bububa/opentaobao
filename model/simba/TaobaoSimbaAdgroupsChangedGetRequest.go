package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页获取修改的推广组ID和修改时间 APIRequest
taobao.simba.adgroups.changed.get

分页获取修改的推广组ID和修改时间
*/
type TaobaoSimbaAdgroupsChangedGetRequest struct {
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

func NewTaobaoSimbaAdgroupsChangedGetRequest() *TaobaoSimbaAdgroupsChangedGetRequest{
    return &TaobaoSimbaAdgroupsChangedGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaAdgroupsChangedGetRequest) GetApiMethodName() string {
    return "taobao.simba.adgroups.changed.get"
}

func (r TaobaoSimbaAdgroupsChangedGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaAdgroupsChangedGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaAdgroupsChangedGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaAdgroupsChangedGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSimbaAdgroupsChangedGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoSimbaAdgroupsChangedGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaAdgroupsChangedGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSimbaAdgroupsChangedGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoSimbaAdgroupsChangedGetRequest) GetPageNo() int64 {
    return r.pageNo
}

