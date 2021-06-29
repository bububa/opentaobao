package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
四类数据同步 API请求
alibaba.alihealth.pregnancy.data.sync

经期调整；基础体温；排卵试纸；B超测排数据同步
*/
type AlibabaAlihealthPregnancyDataSyncRequest struct {
    model.Params
    // 用户id
    userId   int64
    // 三方id
    outerId   int64
    // 4经期调整；1基础体温；2排卵试纸；3 B超测排
    eventType   int64
    // 四类数据定制化详情
    data   string
    // 测量日期
    measureDate   int64
    // 0-新增 1-修改 2-删除
    operationType   int64
    // 经期数据json串
    periodMsg   string
}

// 初始化AlibabaAlihealthPregnancyDataSyncRequest对象
func NewAlibabaAlihealthPregnancyDataSyncRequest() *AlibabaAlihealthPregnancyDataSyncRequest{
    return &AlibabaAlihealthPregnancyDataSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPregnancyDataSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.pregnancy.data.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPregnancyDataSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户id
func (r *AlibabaAlihealthPregnancyDataSyncRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaAlihealthPregnancyDataSyncRequest) GetUserId() int64 {
    return r.userId
}
// OuterId Setter
// 三方id
func (r *AlibabaAlihealthPregnancyDataSyncRequest) SetOuterId(outerId int64) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r AlibabaAlihealthPregnancyDataSyncRequest) GetOuterId() int64 {
    return r.outerId
}
// EventType Setter
// 4经期调整；1基础体温；2排卵试纸；3 B超测排
func (r *AlibabaAlihealthPregnancyDataSyncRequest) SetEventType(eventType int64) error {
    r.eventType = eventType
    r.Set("event_type", eventType)
    return nil
}

// EventType Getter
func (r AlibabaAlihealthPregnancyDataSyncRequest) GetEventType() int64 {
    return r.eventType
}
// Data Setter
// 四类数据定制化详情
func (r *AlibabaAlihealthPregnancyDataSyncRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

// Data Getter
func (r AlibabaAlihealthPregnancyDataSyncRequest) GetData() string {
    return r.data
}
// MeasureDate Setter
// 测量日期
func (r *AlibabaAlihealthPregnancyDataSyncRequest) SetMeasureDate(measureDate int64) error {
    r.measureDate = measureDate
    r.Set("measure_date", measureDate)
    return nil
}

// MeasureDate Getter
func (r AlibabaAlihealthPregnancyDataSyncRequest) GetMeasureDate() int64 {
    return r.measureDate
}
// OperationType Setter
// 0-新增 1-修改 2-删除
func (r *AlibabaAlihealthPregnancyDataSyncRequest) SetOperationType(operationType int64) error {
    r.operationType = operationType
    r.Set("operation_type", operationType)
    return nil
}

// OperationType Getter
func (r AlibabaAlihealthPregnancyDataSyncRequest) GetOperationType() int64 {
    return r.operationType
}
// PeriodMsg Setter
// 经期数据json串
func (r *AlibabaAlihealthPregnancyDataSyncRequest) SetPeriodMsg(periodMsg string) error {
    r.periodMsg = periodMsg
    r.Set("period_msg", periodMsg)
    return nil
}

// PeriodMsg Getter
func (r AlibabaAlihealthPregnancyDataSyncRequest) GetPeriodMsg() string {
    return r.periodMsg
}
