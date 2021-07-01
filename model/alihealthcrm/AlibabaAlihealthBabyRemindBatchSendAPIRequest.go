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
type AlibabaAlihealthBabyRemindBatchSendAPIRequest struct {
    model.Params
    // 批量发送提醒
    _batchRemindRequests   []BatchRemindRequestDto
    // 提醒类型：1-每日任务，2-疫苗提醒，3-产检提醒
    _remindType   int64
}

// 初始化AlibabaAlihealthBabyRemindBatchSendAPIRequest对象
func NewAlibabaAlihealthBabyRemindBatchSendRequest() *AlibabaAlihealthBabyRemindBatchSendAPIRequest{
    return &AlibabaAlihealthBabyRemindBatchSendAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBabyRemindBatchSendAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.baby.remind.batch.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBabyRemindBatchSendAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BatchRemindRequests Setter
// 批量发送提醒
func (r *AlibabaAlihealthBabyRemindBatchSendAPIRequest) SetBatchRemindRequests(_batchRemindRequests []BatchRemindRequestDto) error {
    r._batchRemindRequests = _batchRemindRequests
    r.Set("batch_remind_requests", _batchRemindRequests)
    return nil
}

// BatchRemindRequests Getter
func (r AlibabaAlihealthBabyRemindBatchSendAPIRequest) GetBatchRemindRequests() []BatchRemindRequestDto {
    return r._batchRemindRequests
}
// RemindType Setter
// 提醒类型：1-每日任务，2-疫苗提醒，3-产检提醒
func (r *AlibabaAlihealthBabyRemindBatchSendAPIRequest) SetRemindType(_remindType int64) error {
    r._remindType = _remindType
    r.Set("remind_type", _remindType)
    return nil
}

// RemindType Getter
func (r AlibabaAlihealthBabyRemindBatchSendAPIRequest) GetRemindType() int64 {
    return r._remindType
}
