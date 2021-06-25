package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据时间段查询服务商的服务预警消息列表(15分钟内) APIRequest
tmall.servicecenter.servicemonitormessage.search

根据时间段查询服务商的服务预警消息列表(15分钟内)
*/
type TmallServicecenterServicemonitormessageSearchRequest struct {
    model.Params

    // 开始时间long值
    start   int64 

    // 结束时间long值，与start相差15分钟
    end   int64 

}

func NewTmallServicecenterServicemonitormessageSearchRequest() *TmallServicecenterServicemonitormessageSearchRequest{
    return &TmallServicecenterServicemonitormessageSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterServicemonitormessageSearchRequest) GetApiMethodName() string {
    return "tmall.servicecenter.servicemonitormessage.search"
}

func (r TmallServicecenterServicemonitormessageSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterServicemonitormessageSearchRequest) SetStart(start int64) error {
    r.start = start
    r.Set("start", start)
    return nil
}

func (r TmallServicecenterServicemonitormessageSearchRequest) GetStart() int64 {
    return r.start
}

func (r *TmallServicecenterServicemonitormessageSearchRequest) SetEnd(end int64) error {
    r.end = end
    r.Set("end", end)
    return nil
}

func (r TmallServicecenterServicemonitormessageSearchRequest) GetEnd() int64 {
    return r.end
}

