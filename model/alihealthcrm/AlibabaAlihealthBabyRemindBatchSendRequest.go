package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量发送母婴提醒 API请求
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

// 初始化AlibabaAlihealthBabyRemindBatchSendRequest对象
func NewAlibabaAlihealthBabyRemindBatchSendRequest() *AlibabaAlihealthBabyRemindBatchSendRequest{
    return &AlibabaAlihealthBabyRemindBatchSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBabyRemindBatchSendRequest) GetApiMethodName() string {
    return "alibaba.alihealth.baby.remind.batch.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBabyRemindBatchSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BatchRemindRequests Setter
// 批量发送提醒
func (r *AlibabaAlihealthBabyRemindBatchSendRequest) SetBatchRemindRequests(batchRemindRequests []BatchRemindRequestDto) error {
    r.batchRemindRequests = batchRemindRequests
    r.Set("batch_remind_requests", batchRemindRequests)
    return nil
}

// BatchRemindRequests Getter
func (r AlibabaAlihealthBabyRemindBatchSendRequest) GetBatchRemindRequests() []BatchRemindRequestDto {
    return r.batchRemindRequests
}
// RemindType Setter
// 提醒类型：1-每日任务，2-疫苗提醒，3-产检提醒
func (r *AlibabaAlihealthBabyRemindBatchSendRequest) SetRemindType(remindType int64) error {
    r.remindType = remindType
    r.Set("remind_type", remindType)
    return nil
}

// RemindType Getter
func (r AlibabaAlihealthBabyRemindBatchSendRequest) GetRemindType() int64 {
    return r.remindType
}
