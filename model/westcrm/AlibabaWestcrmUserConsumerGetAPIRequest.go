package westcrm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWestcrmUserConsumerGetAPIRequest 获取指定用户的消费总额 API请求
// alibaba.westcrm.user.consumer.get
//
// 获取指定用户的消费总额
type AlibabaWestcrmUserConsumerGetAPIRequest struct {
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

// NewAlibabaWestcrmUserConsumerGetRequest 初始化AlibabaWestcrmUserConsumerGetAPIRequest对象
func NewAlibabaWestcrmUserConsumerGetRequest() *AlibabaWestcrmUserConsumerGetAPIRequest {
	return &AlibabaWestcrmUserConsumerGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWestcrmUserConsumerGetAPIRequest) Reset() {
	r._timeBegin = ""
	r._timeEnd = ""
	r._campusId = 0
	r._ibUserId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmUserConsumerGetAPIRequest) GetApiMethodName() string {
	return "alibaba.westcrm.user.consumer.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmUserConsumerGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWestcrmUserConsumerGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTimeBegin is TimeBegin Setter
// 开始时间
func (r *AlibabaWestcrmUserConsumerGetAPIRequest) SetTimeBegin(_timeBegin string) error {
	r._timeBegin = _timeBegin
	r.Set("time_begin", _timeBegin)
	return nil
}

// GetTimeBegin TimeBegin Getter
func (r AlibabaWestcrmUserConsumerGetAPIRequest) GetTimeBegin() string {
	return r._timeBegin
}

// SetTimeEnd is TimeEnd Setter
// 结束时间
func (r *AlibabaWestcrmUserConsumerGetAPIRequest) SetTimeEnd(_timeEnd string) error {
	r._timeEnd = _timeEnd
	r.Set("time_end", _timeEnd)
	return nil
}

// GetTimeEnd TimeEnd Getter
func (r AlibabaWestcrmUserConsumerGetAPIRequest) GetTimeEnd() string {
	return r._timeEnd
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabaWestcrmUserConsumerGetAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabaWestcrmUserConsumerGetAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// SetIbUserId is IbUserId Setter
// 用户id
func (r *AlibabaWestcrmUserConsumerGetAPIRequest) SetIbUserId(_ibUserId int64) error {
	r._ibUserId = _ibUserId
	r.Set("ib_user_id", _ibUserId)
	return nil
}

// GetIbUserId IbUserId Getter
func (r AlibabaWestcrmUserConsumerGetAPIRequest) GetIbUserId() int64 {
	return r._ibUserId
}

var poolAlibabaWestcrmUserConsumerGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWestcrmUserConsumerGetRequest()
	},
}

// GetAlibabaWestcrmUserConsumerGetRequest 从 sync.Pool 获取 AlibabaWestcrmUserConsumerGetAPIRequest
func GetAlibabaWestcrmUserConsumerGetAPIRequest() *AlibabaWestcrmUserConsumerGetAPIRequest {
	return poolAlibabaWestcrmUserConsumerGetAPIRequest.Get().(*AlibabaWestcrmUserConsumerGetAPIRequest)
}

// ReleaseAlibabaWestcrmUserConsumerGetAPIRequest 将 AlibabaWestcrmUserConsumerGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWestcrmUserConsumerGetAPIRequest(v *AlibabaWestcrmUserConsumerGetAPIRequest) {
	v.Reset()
	poolAlibabaWestcrmUserConsumerGetAPIRequest.Put(v)
}
