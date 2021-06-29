package aliqin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物联卡用户管理停开机功能 API请求
alibaba.aliqin.fc.iot.useroscontrol

物联网针对用户级管理停开支持
*/
type AlibabaAliqinFcIotUseroscontrolRequest struct {
    model.Params
    // 物联卡的iccid
    _iccid   string
    // 用户停开的操作类型：MANAGE_RESUME、MANAGE_STOP
    _action   string
}

// 初始化AlibabaAliqinFcIotUseroscontrolRequest对象
func NewAlibabaAliqinFcIotUseroscontrolRequest() *AlibabaAliqinFcIotUseroscontrolRequest{
    return &AlibabaAliqinFcIotUseroscontrolRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotUseroscontrolRequest) GetApiMethodName() string {
    return "alibaba.aliqin.fc.iot.useroscontrol"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotUseroscontrolRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Iccid Setter
// 物联卡的iccid
func (r *AlibabaAliqinFcIotUseroscontrolRequest) SetIccid(_iccid string) error {
    r._iccid = _iccid
    r.Set("iccid", _iccid)
    return nil
}

// Iccid Getter
func (r AlibabaAliqinFcIotUseroscontrolRequest) GetIccid() string {
    return r._iccid
}
// Action Setter
// 用户停开的操作类型：MANAGE_RESUME、MANAGE_STOP
func (r *AlibabaAliqinFcIotUseroscontrolRequest) SetAction(_action string) error {
    r._action = _action
    r.Set("action", _action)
    return nil
}

// Action Getter
func (r AlibabaAliqinFcIotUseroscontrolRequest) GetAction() string {
    return r._action
}
