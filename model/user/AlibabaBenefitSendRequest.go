package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发奖接口 API请求
alibaba.benefit.send

发奖接口
*/
type AlibabaBenefitSendRequest struct {
    model.Params
    // 发放的权益(奖品)唯一标识
    _rightEname   string
    // 接收奖品的用户openId
    _receiverId   string
    // 规定填taobao即可
    _userType   string
    // 幂等校验id，业务重试需要，自定义唯一字段即可
    _uniqueId   string
    // mtop
    _appName   string
    // 调用应用ip，非必填
    _ip   string
}

// 初始化AlibabaBenefitSendRequest对象
func NewAlibabaBenefitSendRequest() *AlibabaBenefitSendRequest{
    return &AlibabaBenefitSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaBenefitSendRequest) GetApiMethodName() string {
    return "alibaba.benefit.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaBenefitSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RightEname Setter
// 发放的权益(奖品)唯一标识
func (r *AlibabaBenefitSendRequest) SetRightEname(_rightEname string) error {
    r._rightEname = _rightEname
    r.Set("right_ename", _rightEname)
    return nil
}

// RightEname Getter
func (r AlibabaBenefitSendRequest) GetRightEname() string {
    return r._rightEname
}
// ReceiverId Setter
// 接收奖品的用户openId
func (r *AlibabaBenefitSendRequest) SetReceiverId(_receiverId string) error {
    r._receiverId = _receiverId
    r.Set("receiver_id", _receiverId)
    return nil
}

// ReceiverId Getter
func (r AlibabaBenefitSendRequest) GetReceiverId() string {
    return r._receiverId
}
// UserType Setter
// 规定填taobao即可
func (r *AlibabaBenefitSendRequest) SetUserType(_userType string) error {
    r._userType = _userType
    r.Set("user_type", _userType)
    return nil
}

// UserType Getter
func (r AlibabaBenefitSendRequest) GetUserType() string {
    return r._userType
}
// UniqueId Setter
// 幂等校验id，业务重试需要，自定义唯一字段即可
func (r *AlibabaBenefitSendRequest) SetUniqueId(_uniqueId string) error {
    r._uniqueId = _uniqueId
    r.Set("unique_id", _uniqueId)
    return nil
}

// UniqueId Getter
func (r AlibabaBenefitSendRequest) GetUniqueId() string {
    return r._uniqueId
}
// AppName Setter
// mtop
func (r *AlibabaBenefitSendRequest) SetAppName(_appName string) error {
    r._appName = _appName
    r.Set("app_name", _appName)
    return nil
}

// AppName Getter
func (r AlibabaBenefitSendRequest) GetAppName() string {
    return r._appName
}
// Ip Setter
// 调用应用ip，非必填
func (r *AlibabaBenefitSendRequest) SetIp(_ip string) error {
    r._ip = _ip
    r.Set("ip", _ip)
    return nil
}

// Ip Getter
func (r AlibabaBenefitSendRequest) GetIp() string {
    return r._ip
}
