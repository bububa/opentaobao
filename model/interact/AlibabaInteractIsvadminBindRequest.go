package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建及绑定互动实例 API请求
alibaba.interact.isvadmin.bind

创建互动实例，并绑定奖池
*/
type AlibabaInteractIsvadminBindRequest struct {
    model.Params
    // 描述信息
    _interactDescription   string
    // 互动实例名称
    _instanceName   string
    // 互动开始时间
    _startTime   string
    // 互动结束时间
    _endTime   string
    // 奖池ID
    _lotteryCode   string
}

// 初始化AlibabaInteractIsvadminBindRequest对象
func NewAlibabaInteractIsvadminBindRequest() *AlibabaInteractIsvadminBindRequest{
    return &AlibabaInteractIsvadminBindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractIsvadminBindRequest) GetApiMethodName() string {
    return "alibaba.interact.isvadmin.bind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractIsvadminBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InteractDescription Setter
// 描述信息
func (r *AlibabaInteractIsvadminBindRequest) SetInteractDescription(_interactDescription string) error {
    r._interactDescription = _interactDescription
    r.Set("interact_description", _interactDescription)
    return nil
}

// InteractDescription Getter
func (r AlibabaInteractIsvadminBindRequest) GetInteractDescription() string {
    return r._interactDescription
}
// InstanceName Setter
// 互动实例名称
func (r *AlibabaInteractIsvadminBindRequest) SetInstanceName(_instanceName string) error {
    r._instanceName = _instanceName
    r.Set("instance_name", _instanceName)
    return nil
}

// InstanceName Getter
func (r AlibabaInteractIsvadminBindRequest) GetInstanceName() string {
    return r._instanceName
}
// StartTime Setter
// 互动开始时间
func (r *AlibabaInteractIsvadminBindRequest) SetStartTime(_startTime string) error {
    r._startTime = _startTime
    r.Set("start_time", _startTime)
    return nil
}

// StartTime Getter
func (r AlibabaInteractIsvadminBindRequest) GetStartTime() string {
    return r._startTime
}
// EndTime Setter
// 互动结束时间
func (r *AlibabaInteractIsvadminBindRequest) SetEndTime(_endTime string) error {
    r._endTime = _endTime
    r.Set("end_time", _endTime)
    return nil
}

// EndTime Getter
func (r AlibabaInteractIsvadminBindRequest) GetEndTime() string {
    return r._endTime
}
// LotteryCode Setter
// 奖池ID
func (r *AlibabaInteractIsvadminBindRequest) SetLotteryCode(_lotteryCode string) error {
    r._lotteryCode = _lotteryCode
    r.Set("lottery_code", _lotteryCode)
    return nil
}

// LotteryCode Getter
func (r AlibabaInteractIsvadminBindRequest) GetLotteryCode() string {
    return r._lotteryCode
}
