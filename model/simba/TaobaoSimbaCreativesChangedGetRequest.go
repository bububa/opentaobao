package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页获取修改过的广告创意ID和修改时间 API请求
taobao.simba.creatives.changed.get

分页获取修改过的广告创意ID和修改时间
*/
type TaobaoSimbaCreativesChangedGetRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 返回的每页数据量大小,默认200最大1000
    _pageSize   int64
    // 返回的第几页数据，默认为1
    _pageNo   int64
    // 得到此时间点之后的数据，不能大于一个月
    _startTime   string
}

// 初始化TaobaoSimbaCreativesChangedGetRequest对象
func NewTaobaoSimbaCreativesChangedGetRequest() *TaobaoSimbaCreativesChangedGetRequest{
    return &TaobaoSimbaCreativesChangedGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCreativesChangedGetRequest) GetApiMethodName() string {
    return "taobao.simba.creatives.changed.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCreativesChangedGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCreativesChangedGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCreativesChangedGetRequest) GetNick() string {
    return r._nick
}
// PageSize Setter
// 返回的每页数据量大小,默认200最大1000
func (r *TaobaoSimbaCreativesChangedGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoSimbaCreativesChangedGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// PageNo Setter
// 返回的第几页数据，默认为1
func (r *TaobaoSimbaCreativesChangedGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoSimbaCreativesChangedGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// StartTime Setter
// 得到此时间点之后的数据，不能大于一个月
func (r *TaobaoSimbaCreativesChangedGetRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r TaobaoSimbaCreativesChangedGetRequest) GetStartTime() string {
    return r._startTime
}
