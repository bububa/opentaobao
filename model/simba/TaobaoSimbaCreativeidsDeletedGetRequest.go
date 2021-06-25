package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取删除的创意ID APIRequest
taobao.simba.creativeids.deleted.get

获取删除的创意ID
*/
type TaobaoSimbaCreativeidsDeletedGetRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 得到这个时间点之后的数据，不能大于一个月
    startTime   string 

    // 返回的每页数据量大小,默认200最大1000
    pageSize   int64 

    // 返回的第几页数据，默认为1
    pageNo   int64 

}

func NewTaobaoSimbaCreativeidsDeletedGetRequest() *TaobaoSimbaCreativeidsDeletedGetRequest{
    return &TaobaoSimbaCreativeidsDeletedGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaCreativeidsDeletedGetRequest) GetApiMethodName() string {
    return "taobao.simba.creativeids.deleted.get"
}

func (r TaobaoSimbaCreativeidsDeletedGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaCreativeidsDeletedGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaCreativeidsDeletedGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaCreativeidsDeletedGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TaobaoSimbaCreativeidsDeletedGetRequest) GetStartTime() string {
    return r.startTime
}

func (r *TaobaoSimbaCreativeidsDeletedGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r TaobaoSimbaCreativeidsDeletedGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *TaobaoSimbaCreativeidsDeletedGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r TaobaoSimbaCreativeidsDeletedGetRequest) GetPageNo() int64 {
    return r.pageNo
}

