package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取删除的推广组ID API请求
taobao.simba.adgroupids.deleted.get

获取删除的推广组ID
*/
type TaobaoSimbaAdgroupidsDeletedGetAPIRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 得到此时间点之后的数据，不能大于一个月
    _startTime   string
    // 返回的每页数据量大小,默认200最大1000
    _pageSize   int64
    // 返回的第几页数据，默认为1
    _pageNo   int64
}

// 初始化TaobaoSimbaAdgroupidsDeletedGetAPIRequest对象
func NewTaobaoSimbaAdgroupidsDeletedGetRequest() *TaobaoSimbaAdgroupidsDeletedGetAPIRequest{
    return &TaobaoSimbaAdgroupidsDeletedGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupidsDeletedGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.adgroupids.deleted.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupidsDeletedGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaAdgroupidsDeletedGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaAdgroupidsDeletedGetAPIRequest) GetNick() string {
    return r._nick
}
// StartTime Setter
// 得到此时间点之后的数据，不能大于一个月
func (r *TaobaoSimbaAdgroupidsDeletedGetAPIRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaAdgroupidsDeletedGetAPIRequest) GetStartTime() string {
    return r._startTime
}
// PageSize Setter
// 返回的每页数据量大小,默认200最大1000
func (r *TaobaoSimbaAdgroupidsDeletedGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaAdgroupidsDeletedGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNo Setter
// 返回的第几页数据，默认为1
func (r *TaobaoSimbaAdgroupidsDeletedGetAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaAdgroupidsDeletedGetAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
