package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityCharitytimeQuerytimeAPIRequest 外部查询公益时 API请求
// alibaba.charity.charitytime.querytime
//
// 外部查询公益时
type AlibabaCharityCharitytimeQuerytimeAPIRequest struct {
	model.Params
	// 结束时间，不包含
	_endTime string
	// 开始时间，包含
	_startTime string
	// 用户ID
	_userKey string
	// 用户类型
	_userType string
}

// NewAlibabaCharityCharitytimeQuerytimeRequest 初始化AlibabaCharityCharitytimeQuerytimeAPIRequest对象
func NewAlibabaCharityCharitytimeQuerytimeRequest() *AlibabaCharityCharitytimeQuerytimeAPIRequest {
	return &AlibabaCharityCharitytimeQuerytimeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCharityCharitytimeQuerytimeAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.charitytime.querytime"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCharityCharitytimeQuerytimeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCharityCharitytimeQuerytimeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEndTime is EndTime Setter
// 结束时间，不包含
func (r *AlibabaCharityCharitytimeQuerytimeAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AlibabaCharityCharitytimeQuerytimeAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetStartTime is StartTime Setter
// 开始时间，包含
func (r *AlibabaCharityCharitytimeQuerytimeAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r AlibabaCharityCharitytimeQuerytimeAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetUserKey is UserKey Setter
// 用户ID
func (r *AlibabaCharityCharitytimeQuerytimeAPIRequest) SetUserKey(_userKey string) error {
	r._userKey = _userKey
	r.Set("user_key", _userKey)
	return nil
}

// GetUserKey UserKey Getter
func (r AlibabaCharityCharitytimeQuerytimeAPIRequest) GetUserKey() string {
	return r._userKey
}

// SetUserType is UserType Setter
// 用户类型
func (r *AlibabaCharityCharitytimeQuerytimeAPIRequest) SetUserType(_userType string) error {
	r._userType = _userType
	r.Set("user_type", _userType)
	return nil
}

// GetUserType UserType Getter
func (r AlibabaCharityCharitytimeQuerytimeAPIRequest) GetUserType() string {
	return r._userType
}
