package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
拉取备孕初始化信息 APIRequest
alibaba.alihealth.pregnancy.baseinfo.get

拉取备孕初始化信息
*/
type AlibabaAlihealthPregnancyBaseinfoGetRequest struct {
    model.Params

    // 用户id
    userId   int64 

}

func NewAlibabaAlihealthPregnancyBaseinfoGetRequest() *AlibabaAlihealthPregnancyBaseinfoGetRequest{
    return &AlibabaAlihealthPregnancyBaseinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthPregnancyBaseinfoGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.pregnancy.baseinfo.get"
}

func (r AlibabaAlihealthPregnancyBaseinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthPregnancyBaseinfoGetRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaAlihealthPregnancyBaseinfoGetRequest) GetUserId() int64 {
    return r.userId
}

