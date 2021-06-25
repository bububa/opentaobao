package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建及绑定互动实例 APIRequest
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

func NewAlibabaInteractIsvadminBindRequest() *AlibabaInteractIsvadminBindRequest{
    return &AlibabaInteractIsvadminBindRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractIsvadminBindRequest) GetApiMethodName() string {
    return "alibaba.interact.isvadmin.bind"
}

func (r AlibabaInteractIsvadminBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractIsvadminBindRequest) SetInteractDescription(interactDescription string) error {
    r.interactDescription = interactDescription
    r.Set("interact_description", interactDescription)
    return nil
}

func (r AlibabaInteractIsvadminBindRequest) GetInteractDescription() string {
    return r.interactDescription
}

func (r *AlibabaInteractIsvadminBindRequest) SetInstanceName(instanceName string) error {
    r.instanceName = instanceName
    r.Set("instance_name", instanceName)
    return nil
}

func (r AlibabaInteractIsvadminBindRequest) GetInstanceName() string {
    return r.instanceName
}

func (r *AlibabaInteractIsvadminBindRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r AlibabaInteractIsvadminBindRequest) GetStartTime() string {
    return r.startTime
}

func (r *AlibabaInteractIsvadminBindRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r AlibabaInteractIsvadminBindRequest) GetEndTime() string {
    return r.endTime
}

func (r *AlibabaInteractIsvadminBindRequest) SetLotteryCode(lotteryCode string) error {
    r.lotteryCode = lotteryCode
    r.Set("lottery_code", lotteryCode)
    return nil
}

func (r AlibabaInteractIsvadminBindRequest) GetLotteryCode() string {
    return r.lotteryCode
}

