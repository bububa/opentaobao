package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacharitycharitytimelistAPIRequest 授权用户的公益时记录查询 API请求
// alibaba.charity.charitytime.list
//
// 查询授权用户的公益时记录
type AlibabacharitycharitytimelistAPIRequest struct {
	model.Params
	// 授权码
	_authCode string
	// 发放公益时查询开始时间，包含
	_startTime string
	// 发放公益时查询截止时间，不包含
	_endTime string
	// 额外取的数据域，ACTIVITY_BASE:活动信息
	_fetches string
	// 用户ID
	_userKey string
	// 用户类型
	_userType string
	// 产生公益时查询开始时间，包含
	_createStartTime string
	// 产生公益时查询截止时间，不包含
	_createEndTime string
	// 最大条数，大于2000取2000
	_limit int64
}

// NewAlibabacharitycharitytimelistRequest 初始化AlibabacharitycharitytimelistAPIRequest对象
func NewAlibabacharitycharitytimelistRequest() *AlibabacharitycharitytimelistAPIRequest {
	return &AlibabacharitycharitytimelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacharitycharitytimelistAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.charitytime.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacharitycharitytimelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacharitycharitytimelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuthCode is AuthCode Setter
// 授权码
func (r *AlibabacharitycharitytimelistAPIRequest) SetAuthCode(_authCode string) error {
	r._authCode = _authCode
	r.Set("auth_code", _authCode)
	return nil
}

// GetAuthCode AuthCode Getter
func (r AlibabacharitycharitytimelistAPIRequest) GetAuthCode() string {
	return r._authCode
}

// SetStartTime is StartTime Setter
// 发放公益时查询开始时间，包含
func (r *AlibabacharitycharitytimelistAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r AlibabacharitycharitytimelistAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 发放公益时查询截止时间，不包含
func (r *AlibabacharitycharitytimelistAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AlibabacharitycharitytimelistAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetFetches is Fetches Setter
// 额外取的数据域，ACTIVITY_BASE:活动信息
func (r *AlibabacharitycharitytimelistAPIRequest) SetFetches(_fetches string) error {
	r._fetches = _fetches
	r.Set("fetches", _fetches)
	return nil
}

// GetFetches Fetches Getter
func (r AlibabacharitycharitytimelistAPIRequest) GetFetches() string {
	return r._fetches
}

// SetUserKey is UserKey Setter
// 用户ID
func (r *AlibabacharitycharitytimelistAPIRequest) SetUserKey(_userKey string) error {
	r._userKey = _userKey
	r.Set("user_key", _userKey)
	return nil
}

// GetUserKey UserKey Getter
func (r AlibabacharitycharitytimelistAPIRequest) GetUserKey() string {
	return r._userKey
}

// SetUserType is UserType Setter
// 用户类型
func (r *AlibabacharitycharitytimelistAPIRequest) SetUserType(_userType string) error {
	r._userType = _userType
	r.Set("user_type", _userType)
	return nil
}

// GetUserType UserType Getter
func (r AlibabacharitycharitytimelistAPIRequest) GetUserType() string {
	return r._userType
}

// SetCreateStartTime is CreateStartTime Setter
// 产生公益时查询开始时间，包含
func (r *AlibabacharitycharitytimelistAPIRequest) SetCreateStartTime(_createStartTime string) error {
	r._createStartTime = _createStartTime
	r.Set("create_start_time", _createStartTime)
	return nil
}

// GetCreateStartTime CreateStartTime Getter
func (r AlibabacharitycharitytimelistAPIRequest) GetCreateStartTime() string {
	return r._createStartTime
}

// SetCreateEndTime is CreateEndTime Setter
// 产生公益时查询截止时间，不包含
func (r *AlibabacharitycharitytimelistAPIRequest) SetCreateEndTime(_createEndTime string) error {
	r._createEndTime = _createEndTime
	r.Set("create_end_time", _createEndTime)
	return nil
}

// GetCreateEndTime CreateEndTime Getter
func (r AlibabacharitycharitytimelistAPIRequest) GetCreateEndTime() string {
	return r._createEndTime
}

// SetLimit is Limit Setter
// 最大条数，大于2000取2000
func (r *AlibabacharitycharitytimelistAPIRequest) SetLimit(_limit int64) error {
	r._limit = _limit
	r.Set("limit", _limit)
	return nil
}

// GetLimit Limit Getter
func (r AlibabacharitycharitytimelistAPIRequest) GetLimit() int64 {
	return r._limit
}
