package alihealthcrm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPregnancyDataSyncAPIRequest 四类数据同步 API请求
// alibaba.alihealth.pregnancy.data.sync
//
// 经期调整；基础体温；排卵试纸；B超测排数据同步
type AlibabaAlihealthPregnancyDataSyncAPIRequest struct {
	model.Params
	// 四类数据定制化详情
	_data string
	// 经期数据json串
	_periodMsg string
	// 用户id
	_userId int64
	// 三方id
	_outerId int64
	// 4经期调整；1基础体温；2排卵试纸；3 B超测排
	_eventType int64
	// 测量日期
	_measureDate int64
	// 0-新增 1-修改 2-删除
	_operationType int64
}

// NewAlibabaAlihealthPregnancyDataSyncRequest 初始化AlibabaAlihealthPregnancyDataSyncAPIRequest对象
func NewAlibabaAlihealthPregnancyDataSyncRequest() *AlibabaAlihealthPregnancyDataSyncAPIRequest {
	return &AlibabaAlihealthPregnancyDataSyncAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthPregnancyDataSyncAPIRequest) Reset() {
	r._data = ""
	r._periodMsg = ""
	r._userId = 0
	r._outerId = 0
	r._eventType = 0
	r._measureDate = 0
	r._operationType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPregnancyDataSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pregnancy.data.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPregnancyDataSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthPregnancyDataSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetData is Data Setter
// 四类数据定制化详情
func (r *AlibabaAlihealthPregnancyDataSyncAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabaAlihealthPregnancyDataSyncAPIRequest) GetData() string {
	return r._data
}

// SetPeriodMsg is PeriodMsg Setter
// 经期数据json串
func (r *AlibabaAlihealthPregnancyDataSyncAPIRequest) SetPeriodMsg(_periodMsg string) error {
	r._periodMsg = _periodMsg
	r.Set("period_msg", _periodMsg)
	return nil
}

// GetPeriodMsg PeriodMsg Getter
func (r AlibabaAlihealthPregnancyDataSyncAPIRequest) GetPeriodMsg() string {
	return r._periodMsg
}

// SetUserId is UserId Setter
// 用户id
func (r *AlibabaAlihealthPregnancyDataSyncAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaAlihealthPregnancyDataSyncAPIRequest) GetUserId() int64 {
	return r._userId
}

// SetOuterId is OuterId Setter
// 三方id
func (r *AlibabaAlihealthPregnancyDataSyncAPIRequest) SetOuterId(_outerId int64) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaAlihealthPregnancyDataSyncAPIRequest) GetOuterId() int64 {
	return r._outerId
}

// SetEventType is EventType Setter
// 4经期调整；1基础体温；2排卵试纸；3 B超测排
func (r *AlibabaAlihealthPregnancyDataSyncAPIRequest) SetEventType(_eventType int64) error {
	r._eventType = _eventType
	r.Set("event_type", _eventType)
	return nil
}

// GetEventType EventType Getter
func (r AlibabaAlihealthPregnancyDataSyncAPIRequest) GetEventType() int64 {
	return r._eventType
}

// SetMeasureDate is MeasureDate Setter
// 测量日期
func (r *AlibabaAlihealthPregnancyDataSyncAPIRequest) SetMeasureDate(_measureDate int64) error {
	r._measureDate = _measureDate
	r.Set("measure_date", _measureDate)
	return nil
}

// GetMeasureDate MeasureDate Getter
func (r AlibabaAlihealthPregnancyDataSyncAPIRequest) GetMeasureDate() int64 {
	return r._measureDate
}

// SetOperationType is OperationType Setter
// 0-新增 1-修改 2-删除
func (r *AlibabaAlihealthPregnancyDataSyncAPIRequest) SetOperationType(_operationType int64) error {
	r._operationType = _operationType
	r.Set("operation_type", _operationType)
	return nil
}

// GetOperationType OperationType Getter
func (r AlibabaAlihealthPregnancyDataSyncAPIRequest) GetOperationType() int64 {
	return r._operationType
}

var poolAlibabaAlihealthPregnancyDataSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthPregnancyDataSyncRequest()
	},
}

// GetAlibabaAlihealthPregnancyDataSyncRequest 从 sync.Pool 获取 AlibabaAlihealthPregnancyDataSyncAPIRequest
func GetAlibabaAlihealthPregnancyDataSyncAPIRequest() *AlibabaAlihealthPregnancyDataSyncAPIRequest {
	return poolAlibabaAlihealthPregnancyDataSyncAPIRequest.Get().(*AlibabaAlihealthPregnancyDataSyncAPIRequest)
}

// ReleaseAlibabaAlihealthPregnancyDataSyncAPIRequest 将 AlibabaAlihealthPregnancyDataSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthPregnancyDataSyncAPIRequest(v *AlibabaAlihealthPregnancyDataSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthPregnancyDataSyncAPIRequest.Put(v)
}
