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
    interactDescription   string
    // 互动实例名称
    instanceName   string
    // 互动开始时间
    startTime   string
    // 互动结束时间
    endTime   string
    // 奖池ID
    lotteryCode   string
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
func (r *AlibabaInteractIsvadminBindRequest) SetInteractDescription(interactDescription string) error {
    r.interactDescription = interactDescription
    r.Set("interact_description", interactDescription)
    return nil
}

// InteractDescription Getter
func (r AlibabaInteractIsvadminBindRequest) GetInteractDescription() string {
    return r.interactDescription
}
// InstanceName Setter
// 互动实例名称
func (r *AlibabaInteractIsvadminBindRequest) SetInstanceName(instanceName string) error {
    r.instanceName = instanceName
    r.Set("instance_name", instanceName)
    return nil
}

// InstanceName Getter
func (r AlibabaInteractIsvadminBindRequest) GetInstanceName() string {
    return r.instanceName
}
// StartTime Setter
// 互动开始时间
func (r *AlibabaInteractIsvadminBindRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r AlibabaInteractIsvadminBindRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 互动结束时间
func (r *AlibabaInteractIsvadminBindRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r AlibabaInteractIsvadminBindRequest) GetEndTime() string {
    return r.endTime
}
// LotteryCode Setter
// 奖池ID
func (r *AlibabaInteractIsvadminBindRequest) SetLotteryCode(lotteryCode string) error {
    r.lotteryCode = lotteryCode
    r.Set("lottery_code", lotteryCode)
    return nil
}

// LotteryCode Getter
func (r AlibabaInteractIsvadminBindRequest) GetLotteryCode() string {
    return r.lotteryCode
}
