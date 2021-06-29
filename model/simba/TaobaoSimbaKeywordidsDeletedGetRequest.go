package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取删除的词ID API请求
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

// 初始化TaobaoSimbaKeywordidsDeletedGetRequest对象
func NewTaobaoSimbaKeywordidsDeletedGetRequest() *TaobaoSimbaKeywordidsDeletedGetRequest{
    return &TaobaoSimbaKeywordidsDeletedGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordidsDeletedGetRequest) GetApiMethodName() string {
    return "taobao.simba.keywordids.deleted.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordidsDeletedGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaKeywordidsDeletedGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaKeywordidsDeletedGetRequest) GetNick() string {
    return r.nick
}
// StartTime Setter
// 得到此时间点之后的数据，不能大于一个月
func (r *TaobaoSimbaKeywordidsDeletedGetRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaKeywordidsDeletedGetRequest) GetStartTime() string {
    return r.startTime
}
// PageSize Setter
// 返回的每页数据量大小,默认200最大1000
func (r *TaobaoSimbaKeywordidsDeletedGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaKeywordidsDeletedGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// PageNo Setter
// 返回的第几页数据，默认为1
func (r *TaobaoSimbaKeywordidsDeletedGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaKeywordidsDeletedGetRequest) GetPageNo() int64 {
    return r.pageNo
}
