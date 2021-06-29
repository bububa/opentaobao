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
    rightEname   string
    // 接收奖品的用户openId
    receiverId   string
    // 规定填taobao即可
    userType   string
    // 幂等校验id，业务重试需要，自定义唯一字段即可
    uniqueId   string
    // mtop
    appName   string
    // 调用应用ip，非必填
    ip   string
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
func (r *AlibabaBenefitSendRequest) SetRightEname(rightEname string) error {
    r.rightEname = rightEname
    r.Set("right_ename", rightEname)
    return nil
}

// RightEname Getter
func (r AlibabaBenefitSendRequest) GetRightEname() string {
    return r.rightEname
}
// ReceiverId Setter
// 接收奖品的用户openId
func (r *AlibabaBenefitSendRequest) SetReceiverId(receiverId string) error {
    r.receiverId = receiverId
    r.Set("receiver_id", receiverId)
    return nil
}

// ReceiverId Getter
func (r AlibabaBenefitSendRequest) GetReceiverId() string {
    return r.receiverId
}
// UserType Setter
// 规定填taobao即可
func (r *AlibabaBenefitSendRequest) SetUserType(userType string) error {
    r.userType = userType
    r.Set("user_type", userType)
    return nil
}

// UserType Getter
func (r AlibabaBenefitSendRequest) GetUserType() string {
    return r.userType
}
// UniqueId Setter
// 幂等校验id，业务重试需要，自定义唯一字段即可
func (r *AlibabaBenefitSendRequest) SetUniqueId(uniqueId string) error {
    r.uniqueId = uniqueId
    r.Set("unique_id", uniqueId)
    return nil
}

// UniqueId Getter
func (r AlibabaBenefitSendRequest) GetUniqueId() string {
    return r.uniqueId
}
// AppName Setter
// mtop
func (r *AlibabaBenefitSendRequest) SetAppName(appName string) error {
    r.appName = appName
    r.Set("app_name", appName)
    return nil
}

// AppName Getter
func (r AlibabaBenefitSendRequest) GetAppName() string {
    return r.appName
}
// Ip Setter
// 调用应用ip，非必填
func (r *AlibabaBenefitSendRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

// Ip Getter
func (r AlibabaBenefitSendRequest) GetIp() string {
    return r.ip
}
