package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户基础信息 API请求
alibaba.alihealth.user.baseinfo.get

获取用户基础信息
*/
type AlibabaAlihealthUserBaseinfoGetRequest struct {
    model.Params
    // 用户id
    _userId   int64
    // 三方服务商
    _appName   string
}

// 初始化AlibabaAlihealthUserBaseinfoGetRequest对象
func NewAlibabaAlihealthUserBaseinfoGetRequest() *AlibabaAlihealthUserBaseinfoGetRequest{
    return &AlibabaAlihealthUserBaseinfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthUserBaseinfoGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.user.baseinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthUserBaseinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户id
func (r *AlibabaAlihealthUserBaseinfoGetRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaAlihealthUserBaseinfoGetRequest) GetUserId() int64 {
    return r._userId
}
// AppName Setter
// 三方服务商
func (r *AlibabaAlihealthUserBaseinfoGetRequest) SetAppName(_appName string) error {
    r._appName = _appName
    r.Set("app_name", _appName)
    return nil
}

// AppName Getter
func (r AlibabaAlihealthUserBaseinfoGetRequest) GetAppName() string {
    return r._appName
}
