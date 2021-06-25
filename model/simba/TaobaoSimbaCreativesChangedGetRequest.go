package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页获取修改过的广告创意ID和修改时间 APIRequest
taobao.simba.creatives.changed.get

分页获取修改过的广告创意ID和修改时间
*/
type TaobaoSimbaCreativesChangedGetRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 返回的每页数据量大小,默认200最大1000
    pageSize   int64 

    // 返回的第几页数据，默认为1
    pageNo   int64 

    // 得到此时间点之后的数据，不能大于一个月
    startTime   string 

}

func NewTaobaoSimbaCreativesChangedGetRequest() *TaobaoSimbaCreativesChangedGetRequest{
    return &TaobaoSimbaCreativesChangedGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaCreativesChangedGetRequest) GetApiMethodName() string {
    return "taobao.simba.creatives.changed.get"
}

func (r TaobaoSimbaCreativesChangedGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaCreativesChangedGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaCreativesChangedGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaCreativesChangedGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaCreativesChangedGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSimbaCreativesChangedGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoSimbaCreativesChangedGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *TaobaoSimbaCreativesChangedGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSimbaCreativesChangedGetRequest) GetStartTime() string {
    return r.startTime
}

