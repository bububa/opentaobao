package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
拉取备孕初始化信息 API请求
alibaba.alihealth.pregnancy.baseinfo.get

拉取备孕初始化信息
*/
type AlibabaAlihealthPregnancyBaseinfoGetRequest struct {
    model.Params
    // 用户id
    _userId   int64
}

// 初始化AlibabaAlihealthPregnancyBaseinfoGetRequest对象
func NewAlibabaAlihealthPregnancyBaseinfoGetRequest() *AlibabaAlihealthPregnancyBaseinfoGetRequest{
    return &AlibabaAlihealthPregnancyBaseinfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPregnancyBaseinfoGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.pregnancy.baseinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPregnancyBaseinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户id
func (r *AlibabaAlihealthPregnancyBaseinfoGetRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaAlihealthPregnancyBaseinfoGetRequest) GetUserId() int64 {
    return r._userId
}
