package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商15分钟获取数据 API请求
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

// 初始化TmallServiceSettleadjustmentSearchRequest对象
func NewTmallServiceSettleadjustmentSearchRequest() *TmallServiceSettleadjustmentSearchRequest{
    return &TmallServiceSettleadjustmentSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServiceSettleadjustmentSearchRequest) GetApiMethodName() string {
    return "tmall.service.settleadjustment.search"
}

// IRequest interface 方法, 获取API参数
func (r TmallServiceSettleadjustmentSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EndTime Setter
// 结束时间
func (r *TmallServiceSettleadjustmentSearchRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r TmallServiceSettleadjustmentSearchRequest) GetEndTime() string {
    return r.endTime
}
// StartTime Setter
// 开始时间
func (r *TmallServiceSettleadjustmentSearchRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r TmallServiceSettleadjustmentSearchRequest) GetStartTime() string {
    return r.startTime
}
