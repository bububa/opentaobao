package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthpregnancydatasyncAPIRequest 四类数据同步 API请求
// alibaba.alihealth.pregnancy.data.sync
//
// 经期调整；基础体温；排卵试纸；B超测排数据同步
type AlibabaalihealthpregnancydatasyncAPIRequest struct {
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

// NewAlibabaalihealthpregnancydatasyncRequest 初始化AlibabaalihealthpregnancydatasyncAPIRequest对象
func NewAlibabaalihealthpregnancydatasyncRequest() *AlibabaalihealthpregnancydatasyncAPIRequest {
	return &AlibabaalihealthpregnancydatasyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthpregnancydatasyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pregnancy.data.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthpregnancydatasyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthpregnancydatasyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetData is Data Setter
// 四类数据定制化详情
func (r *AlibabaalihealthpregnancydatasyncAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r AlibabaalihealthpregnancydatasyncAPIRequest) GetData() string {
	return r._data
}

// SetPeriodMsg is PeriodMsg Setter
// 经期数据json串
func (r *AlibabaalihealthpregnancydatasyncAPIRequest) SetPeriodMsg(_periodMsg string) error {
	r._periodMsg = _periodMsg
	r.Set("period_msg", _periodMsg)
	return nil
}

// GetPeriodMsg PeriodMsg Getter
func (r AlibabaalihealthpregnancydatasyncAPIRequest) GetPeriodMsg() string {
	return r._periodMsg
}

// SetUserId is UserId Setter
// 用户id
func (r *AlibabaalihealthpregnancydatasyncAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaalihealthpregnancydatasyncAPIRequest) GetUserId() int64 {
	return r._userId
}

// SetOuterId is OuterId Setter
// 三方id
func (r *AlibabaalihealthpregnancydatasyncAPIRequest) SetOuterId(_outerId int64) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaalihealthpregnancydatasyncAPIRequest) GetOuterId() int64 {
	return r._outerId
}

// SetEventType is EventType Setter
// 4经期调整；1基础体温；2排卵试纸；3 B超测排
func (r *AlibabaalihealthpregnancydatasyncAPIRequest) SetEventType(_eventType int64) error {
	r._eventType = _eventType
	r.Set("event_type", _eventType)
	return nil
}

// GetEventType EventType Getter
func (r AlibabaalihealthpregnancydatasyncAPIRequest) GetEventType() int64 {
	return r._eventType
}

// SetMeasureDate is MeasureDate Setter
// 测量日期
func (r *AlibabaalihealthpregnancydatasyncAPIRequest) SetMeasureDate(_measureDate int64) error {
	r._measureDate = _measureDate
	r.Set("measure_date", _measureDate)
	return nil
}

// GetMeasureDate MeasureDate Getter
func (r AlibabaalihealthpregnancydatasyncAPIRequest) GetMeasureDate() int64 {
	return r._measureDate
}

// SetOperationType is OperationType Setter
// 0-新增 1-修改 2-删除
func (r *AlibabaalihealthpregnancydatasyncAPIRequest) SetOperationType(_operationType int64) error {
	r._operationType = _operationType
	r.Set("operation_type", _operationType)
	return nil
}

// GetOperationType OperationType Getter
func (r AlibabaalihealthpregnancydatasyncAPIRequest) GetOperationType() int64 {
	return r._operationType
}
