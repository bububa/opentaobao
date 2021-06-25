package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商15分钟获取数据 APIRequest
tmall.service.settleadjustment.search

天猫服务平台，按修改时间，时间间隔在15中内（包含15分钟），获取调整单数据
*/
type TmallServiceSettleadjustmentSearchRequest struct {
    model.Params

    // 结束时间
    endTime   string 

    // 开始时间
    startTime   string 

}

func NewTmallServiceSettleadjustmentSearchRequest() *TmallServiceSettleadjustmentSearchRequest{
    return &TmallServiceSettleadjustmentSearchRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServiceSettleadjustmentSearchRequest) GetApiMethodName() string {
    return "tmall.service.settleadjustment.search"
}

func (r TmallServiceSettleadjustmentSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServiceSettleadjustmentSearchRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r TmallServiceSettleadjustmentSearchRequest) GetEndTime() string {
    return r.endTime
}

func (r *TmallServiceSettleadjustmentSearchRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r TmallServiceSettleadjustmentSearchRequest) GetStartTime() string {
    return r.startTime
}

