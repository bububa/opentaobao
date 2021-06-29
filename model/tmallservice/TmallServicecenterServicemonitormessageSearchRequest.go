package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据时间段查询服务商的服务预警消息列表(15分钟内) API请求
tmall.servicecenter.servicemonitormessage.search

根据时间段查询服务商的服务预警消息列表(15分钟内)
*/
type TmallServicecenterServicemonitormessageSearchRequest struct {
    model.Params
    // 开始时间long值
    _start   int64
    // 结束时间long值，与start相差15分钟
    _end   int64
}

// 初始化TmallServicecenterServicemonitormessageSearchRequest对象
func NewTmallServicecenterServicemonitormessageSearchRequest() *TmallServicecenterServicemonitormessageSearchRequest{
    return &TmallServicecenterServicemonitormessageSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicemonitormessageSearchRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicemonitormessage.search"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicemonitormessageSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Start Setter
// 开始时间long值
func (r *TmallServicecenterServicemonitormessageSearchRequest) SetStart(_start int64) error {
    r._start = _start
    r.Set("start", _start)
    return nil
}

// Start Getter
func (r TmallServicecenterServicemonitormessageSearchRequest) GetStart() int64 {
    return r._start
}
// End Setter
// 结束时间long值，与start相差15分钟
func (r *TmallServicecenterServicemonitormessageSearchRequest) SetEnd(_end int64) error {
    r._end = _end
    r.Set("end", _end)
    return nil
}

// End Getter
func (r TmallServicecenterServicemonitormessageSearchRequest) GetEnd() int64 {
    return r._end
}
