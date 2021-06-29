package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取删除的创意ID API请求
taobao.simba.creativeids.deleted.get

获取删除的创意ID
*/
type TaobaoSimbaCreativeidsDeletedGetRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 得到这个时间点之后的数据，不能大于一个月
    _startTime   string
    // 返回的每页数据量大小,默认200最大1000
    _pageSize   int64
    // 返回的第几页数据，默认为1
    _pageNo   int64
}

// 初始化TaobaoSimbaCreativeidsDeletedGetRequest对象
func NewTaobaoSimbaCreativeidsDeletedGetRequest() *TaobaoSimbaCreativeidsDeletedGetRequest{
    return &TaobaoSimbaCreativeidsDeletedGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCreativeidsDeletedGetRequest) GetApiMethodName() string {
    return "taobao.simba.creativeids.deleted.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCreativeidsDeletedGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCreativeidsDeletedGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCreativeidsDeletedGetRequest) GetNick() string {
    return r._nick
}
// StartTime Setter
// 得到这个时间点之后的数据，不能大于一个月
func (r *TaobaoSimbaCreativeidsDeletedGetRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaCreativeidsDeletedGetRequest) GetStartTime() string {
    return r._startTime
}
// PageSize Setter
// 返回的每页数据量大小,默认200最大1000
func (r *TaobaoSimbaCreativeidsDeletedGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaCreativeidsDeletedGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNo Setter
// 返回的第几页数据，默认为1
func (r *TaobaoSimbaCreativeidsDeletedGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaCreativeidsDeletedGetRequest) GetPageNo() int64 {
    return r._pageNo
}
