package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户基础信息 APIRequest
alibaba.alihealth.user.baseinfo.get

获取用户基础信息
*/
type AlibabaAlihealthUserBaseinfoGetRequest struct {
    model.Params

    // 用户id
    userId   int64 

    // 三方服务商
    appName   string 

}

func NewAlibabaAlihealthUserBaseinfoGetRequest() *AlibabaAlihealthUserBaseinfoGetRequest{
    return &AlibabaAlihealthUserBaseinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthUserBaseinfoGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.user.baseinfo.get"
}

func (r AlibabaAlihealthUserBaseinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthUserBaseinfoGetRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaAlihealthUserBaseinfoGetRequest) GetUserId() int64 {
    return r.userId
}

func (r *AlibabaAlihealthUserBaseinfoGetRequest) SetAppName(appName string) error {
    r.appName = appName
    r.Set("app_name", appName)
    return nil
}

func (r AlibabaAlihealthUserBaseinfoGetRequest) GetAppName() string {
    return r.appName
}

