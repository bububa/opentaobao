package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方从我们这获取宝宝信息 APIRequest
alibaba.alihealth.baby.baseinfo.get

三方从我们这获取宝宝信息
*/
type AlibabaAlihealthBabyBaseinfoGetRequest struct {
    model.Params

    // 宝宝id
    babyId   int64 

    // 宝宝所属的用户
    tpUserId   int64 

}

func NewAlibabaAlihealthBabyBaseinfoGetRequest() *AlibabaAlihealthBabyBaseinfoGetRequest{
    return &AlibabaAlihealthBabyBaseinfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthBabyBaseinfoGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.baby.baseinfo.get"
}

func (r AlibabaAlihealthBabyBaseinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthBabyBaseinfoGetRequest) SetBabyId(babyId int64) error {
    r.babyId = babyId
    r.Set("baby_id", babyId)
    return nil
}

func (r AlibabaAlihealthBabyBaseinfoGetRequest) GetBabyId() int64 {
    return r.babyId
}

func (r *AlibabaAlihealthBabyBaseinfoGetRequest) SetTpUserId(tpUserId int64) error {
    r.tpUserId = tpUserId
    r.Set("tp_user_id", tpUserId)
    return nil
}

func (r AlibabaAlihealthBabyBaseinfoGetRequest) GetTpUserId() int64 {
    return r.tpUserId
}

