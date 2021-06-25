package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取修改的推广组ID APIRequest
taobao.simba.adgroupids.changed.get

获取修改的推广组ID
*/
type TaobaoSimbaAdgroupidsChangedGetRequest struct {
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

func NewTaobaoSimbaAdgroupidsChangedGetRequest() *TaobaoSimbaAdgroupidsChangedGetRequest{
    return &TaobaoSimbaAdgroupidsChangedGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaAdgroupidsChangedGetRequest) GetApiMethodName() string {
    return "taobao.simba.adgroupids.changed.get"
}

func (r TaobaoSimbaAdgroupidsChangedGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaAdgroupidsChangedGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaAdgroupidsChangedGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaAdgroupidsChangedGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSimbaAdgroupidsChangedGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoSimbaAdgroupidsChangedGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaAdgroupidsChangedGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSimbaAdgroupidsChangedGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoSimbaAdgroupidsChangedGetRequest) GetPageNo() int64 {
    return r.pageNo
}

