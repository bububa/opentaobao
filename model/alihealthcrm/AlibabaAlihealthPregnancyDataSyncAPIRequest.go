package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthPregnancyDataSyncAPIRequest
四类数据同步 API请求
alibaba.alihealth.pregnancy.data.sync

经期调整；基础体温；排卵试纸；B超测排数据同步 */
type AlibabaAlihealthPregnancyDataSyncAPIRequest struct {
	model.Params
	// 用户id
	_userId int64
	// 三方id
	_outerId int64
	// 4经期调整；1基础体温；2排卵试纸；3 B超测排
	_eventType int64
	// 四类数据定制化详情
	_data string
	// 测量日期
	_measureDate int64
	// 0-新增 1-修改 2-删除
	_operationType int64
	// 经期数据json串
	_periodMsg string
}

// NewAlibabaAlihealthPregnancyDataSyncRequest 初始化AlibabaAlihealthPregnancyDataSyncAPIRequest对象
func NewAlibabaAlihealthPregnancyDataSyncRequest() *AlibabaAlihealthPregnancyDataSyncAPIRequest {
	return &AlibabaAlihealthPregnancyDataSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPregnancyDataSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pregnancy.data.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPregnancyDataSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UserId Setter
// 用户id
func (r *AlibabaAlihealthPregnancyDataSyncAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r AlibabaAlihealthPregnancyDataSyncAPIRequest) GetUserId() int64 {
	return r._userId
}

// Set is OuterId Setter
// 三方id
func (r *AlibabaAlihealthPregnancyDataSyncAPIRequest) SetOuterId(_outerId int64) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// Get OuterId Getter
func (r AlibabaAlihealthPregnancyDataSyncAPIRequest) GetOuterId() int64 {
	return r._outerId
}

// Set is EventType Setter
// 4经期调整；1基础体温；2排卵试纸；3 B超测排
func (r *AlibabaAlihealthPregnancyDataSyncAPIRequest) SetEventType(_eventType int64) error {
	r._eventType = _eventType
	r.Set("event_type", _eventType)
	return nil
}

// Get EventType Getter
func (r AlibabaAlihealthPregnancyDataSyncAPIRequest) GetEventType() int64 {
	return r._eventType
}

// Set is Data Setter
// 四类数据定制化详情
func (r *AlibabaAlihealthPregnancyDataSyncAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// Get Data Getter
func (r AlibabaAlihealthPregnancyDataSyncAPIRequest) GetData() string {
	return r._data
}

// Set is MeasureDate Setter
// 测量日期
func (r *AlibabaAlihealthPregnancyDataSyncAPIRequest) SetMeasureDate(_measureDate int64) error {
	r._measureDate = _measureDate
	r.Set("measure_date", _measureDate)
	return nil
}

// Get MeasureDate Getter
func (r AlibabaAlihealthPregnancyDataSyncAPIRequest) GetMeasureDate() int64 {
	return r._measureDate
}

// Set is OperationType Setter
// 0-新增 1-修改 2-删除
func (r *AlibabaAlihealthPregnancyDataSyncAPIRequest) SetOperationType(_operationType int64) error {
	r._operationType = _operationType
	r.Set("operation_type", _operationType)
	return nil
}

// Get OperationType Getter
func (r AlibabaAlihealthPregnancyDataSyncAPIRequest) GetOperationType() int64 {
	return r._operationType
}

// Set is PeriodMsg Setter
// 经期数据json串
func (r *AlibabaAlihealthPregnancyDataSyncAPIRequest) SetPeriodMsg(_periodMsg string) error {
	r._periodMsg = _periodMsg
	r.Set("period_msg", _periodMsg)
	return nil
}

// Get PeriodMsg Getter
func (r AlibabaAlihealthPregnancyDataSyncAPIRequest) GetPeriodMsg() string {
	return r._periodMsg
}
