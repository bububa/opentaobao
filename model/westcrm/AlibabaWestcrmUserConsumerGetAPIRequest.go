package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawestcrmuserconsumergetAPIRequest 获取指定用户的消费总额 API请求
// alibaba.westcrm.user.consumer.get
//
// 获取指定用户的消费总额
type AlibabawestcrmuserconsumergetAPIRequest struct {
	model.Params
	// 开始时间
	_timeBegin string
	// 结束时间
	_timeEnd string
	// 园区id
	_campusId int64
	// 用户id
	_ibUserId int64
}

// NewAlibabawestcrmuserconsumergetRequest 初始化AlibabawestcrmuserconsumergetAPIRequest对象
func NewAlibabawestcrmuserconsumergetRequest() *AlibabawestcrmuserconsumergetAPIRequest {
	return &AlibabawestcrmuserconsumergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawestcrmuserconsumergetAPIRequest) GetApiMethodName() string {
	return "alibaba.westcrm.user.consumer.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawestcrmuserconsumergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawestcrmuserconsumergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTimeBegin is TimeBegin Setter
// 开始时间
func (r *AlibabawestcrmuserconsumergetAPIRequest) SetTimeBegin(_timeBegin string) error {
	r._timeBegin = _timeBegin
	r.Set("time_begin", _timeBegin)
	return nil
}

// GetTimeBegin TimeBegin Getter
func (r AlibabawestcrmuserconsumergetAPIRequest) GetTimeBegin() string {
	return r._timeBegin
}

// SetTimeEnd is TimeEnd Setter
// 结束时间
func (r *AlibabawestcrmuserconsumergetAPIRequest) SetTimeEnd(_timeEnd string) error {
	r._timeEnd = _timeEnd
	r.Set("time_end", _timeEnd)
	return nil
}

// GetTimeEnd TimeEnd Getter
func (r AlibabawestcrmuserconsumergetAPIRequest) GetTimeEnd() string {
	return r._timeEnd
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabawestcrmuserconsumergetAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabawestcrmuserconsumergetAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// SetIbUserId is IbUserId Setter
// 用户id
func (r *AlibabawestcrmuserconsumergetAPIRequest) SetIbUserId(_ibUserId int64) error {
	r._ibUserId = _ibUserId
	r.Set("ib_user_id", _ibUserId)
	return nil
}

// GetIbUserId IbUserId Getter
func (r AlibabawestcrmuserconsumergetAPIRequest) GetIbUserId() int64 {
	return r._ibUserId
}
