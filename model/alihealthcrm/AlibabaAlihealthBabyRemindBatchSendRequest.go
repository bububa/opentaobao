package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量发送母婴提醒 APIRequest
alibaba.alihealth.baby.remind.batch.send

批量发送母婴提醒
*/
type AlibabaAlihealthBabyRemindBatchSendRequest struct {
    model.Params

    // 批量发送提醒
    batchRemindRequests   []BatchRemindRequestDto 

    // 提醒类型：1-每日任务，2-疫苗提醒，3-产检提醒
    remindType   int64 

}

func NewAlibabaAlihealthBabyRemindBatchSendRequest() *AlibabaAlihealthBabyRemindBatchSendRequest{
    return &AlibabaAlihealthBabyRemindBatchSendRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthBabyRemindBatchSendRequest) GetApiMethodName() string {
    return "alibaba.alihealth.baby.remind.batch.send"
}

func (r AlibabaAlihealthBabyRemindBatchSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthBabyRemindBatchSendRequest) SetBatchRemindRequests(batchRemindRequests []BatchRemindRequestDto) error {
    r.batchRemindRequests = batchRemindRequests
    r.Set("batch_remind_requests", batchRemindRequests)
    return nil
}

func (r AlibabaAlihealthBabyRemindBatchSendRequest) GetBatchRemindRequests() []BatchRemindRequestDto {
    return r.batchRemindRequests
}

func (r *AlibabaAlihealthBabyRemindBatchSendRequest) SetRemindType(remindType int64) error {
    r.remindType = remindType
    r.Set("remind_type", remindType)
    return nil
}

func (r AlibabaAlihealthBabyRemindBatchSendRequest) GetRemindType() int64 {
    return r.remindType
}

