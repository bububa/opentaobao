package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBenefitSendAPIRequest 发奖接口 API请求
// alibaba.benefit.send
//
// 发奖接口
type AlibabaBenefitSendAPIRequest struct {
	model.Params
	// 发放的权益(奖品)唯一标识
	_rightEname string
	// 接收奖品的用户openId
	_receiverId string
	// 规定填taobao即可
	_userType string
	// 幂等校验id，业务重试需要，自定义唯一字段即可
	_uniqueId string
	// mtop
	_appName string
	// 调用应用ip，非必填
	_ip string
}

// NewAlibabaBenefitSendRequest 初始化AlibabaBenefitSendAPIRequest对象
func NewAlibabaBenefitSendRequest() *AlibabaBenefitSendAPIRequest {
	return &AlibabaBenefitSendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBenefitSendAPIRequest) GetApiMethodName() string {
	return "alibaba.benefit.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBenefitSendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaBenefitSendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRightEname is RightEname Setter
// 发放的权益(奖品)唯一标识
func (r *AlibabaBenefitSendAPIRequest) SetRightEname(_rightEname string) error {
	r._rightEname = _rightEname
	r.Set("right_ename", _rightEname)
	return nil
}

// GetRightEname RightEname Getter
func (r AlibabaBenefitSendAPIRequest) GetRightEname() string {
	return r._rightEname
}

// SetReceiverId is ReceiverId Setter
// 接收奖品的用户openId
func (r *AlibabaBenefitSendAPIRequest) SetReceiverId(_receiverId string) error {
	r._receiverId = _receiverId
	r.Set("receiver_id", _receiverId)
	return nil
}

// GetReceiverId ReceiverId Getter
func (r AlibabaBenefitSendAPIRequest) GetReceiverId() string {
	return r._receiverId
}

// SetUserType is UserType Setter
// 规定填taobao即可
func (r *AlibabaBenefitSendAPIRequest) SetUserType(_userType string) error {
	r._userType = _userType
	r.Set("user_type", _userType)
	return nil
}

// GetUserType UserType Getter
func (r AlibabaBenefitSendAPIRequest) GetUserType() string {
	return r._userType
}

// SetUniqueId is UniqueId Setter
// 幂等校验id，业务重试需要，自定义唯一字段即可
func (r *AlibabaBenefitSendAPIRequest) SetUniqueId(_uniqueId string) error {
	r._uniqueId = _uniqueId
	r.Set("unique_id", _uniqueId)
	return nil
}

// GetUniqueId UniqueId Getter
func (r AlibabaBenefitSendAPIRequest) GetUniqueId() string {
	return r._uniqueId
}

// SetAppName is AppName Setter
// mtop
func (r *AlibabaBenefitSendAPIRequest) SetAppName(_appName string) error {
	r._appName = _appName
	r.Set("app_name", _appName)
	return nil
}

// GetAppName AppName Getter
func (r AlibabaBenefitSendAPIRequest) GetAppName() string {
	return r._appName
}

// SetIp is Ip Setter
// 调用应用ip，非必填
func (r *AlibabaBenefitSendAPIRequest) SetIp(_ip string) error {
	r._ip = _ip
	r.Set("ip", _ip)
	return nil
}

// GetIp Ip Getter
func (r AlibabaBenefitSendAPIRequest) GetIp() string {
	return r._ip
}
