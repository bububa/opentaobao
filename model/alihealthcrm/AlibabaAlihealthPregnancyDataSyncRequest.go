package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
四类数据同步 APIRequest
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

func NewAlibabaAlihealthPregnancyDataSyncRequest() *AlibabaAlihealthPregnancyDataSyncRequest{
    return &AlibabaAlihealthPregnancyDataSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthPregnancyDataSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.pregnancy.data.sync"
}

func (r AlibabaAlihealthPregnancyDataSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthPregnancyDataSyncRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaAlihealthPregnancyDataSyncRequest) GetUserId() int64 {
    return r.userId
}

func (r *AlibabaAlihealthPregnancyDataSyncRequest) SetOuterId(outerId int64) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r AlibabaAlihealthPregnancyDataSyncRequest) GetOuterId() int64 {
    return r.outerId
}

func (r *AlibabaAlihealthPregnancyDataSyncRequest) SetEventType(eventType int64) error {
    r.eventType = eventType
    r.Set("event_type", eventType)
    return nil
}

func (r AlibabaAlihealthPregnancyDataSyncRequest) GetEventType() int64 {
    return r.eventType
}

func (r *AlibabaAlihealthPregnancyDataSyncRequest) SetData(data string) error {
    r.data = data
    r.Set("data", data)
    return nil
}

func (r AlibabaAlihealthPregnancyDataSyncRequest) GetData() string {
    return r.data
}

func (r *AlibabaAlihealthPregnancyDataSyncRequest) SetMeasureDate(measureDate int64) error {
    r.measureDate = measureDate
    r.Set("measure_date", measureDate)
    return nil
}

func (r AlibabaAlihealthPregnancyDataSyncRequest) GetMeasureDate() int64 {
    return r.measureDate
}

func (r *AlibabaAlihealthPregnancyDataSyncRequest) SetOperationType(operationType int64) error {
    r.operationType = operationType
    r.Set("operation_type", operationType)
    return nil
}

func (r AlibabaAlihealthPregnancyDataSyncRequest) GetOperationType() int64 {
    return r.operationType
}

func (r *AlibabaAlihealthPregnancyDataSyncRequest) SetPeriodMsg(periodMsg string) error {
    r.periodMsg = periodMsg
    r.Set("period_msg", periodMsg)
    return nil
}

func (r AlibabaAlihealthPregnancyDataSyncRequest) GetPeriodMsg() string {
    return r.periodMsg
}

