package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
权益池信息查询 APIRequest
alibaba.interact.supplier.award.resource.get.cuntao

农村淘宝营销互动权益池信息查询
*/
type AlibabaInteractSupplierAwardResourceGetCuntaoRequest struct {
    model.Params

    // 用户昵称
    userNick   string 

    // 活动code
    activityKey   string 

    // 经度
    lng   string 

    // 纬度
    lat   string 

}

func NewAlibabaInteractSupplierAwardResourceGetCuntaoRequest() *AlibabaInteractSupplierAwardResourceGetCuntaoRequest{
    return &AlibabaInteractSupplierAwardResourceGetCuntaoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractSupplierAwardResourceGetCuntaoRequest) GetApiMethodName() string {
    return "alibaba.interact.supplier.award.resource.get.cuntao"
}

func (r AlibabaInteractSupplierAwardResourceGetCuntaoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractSupplierAwardResourceGetCuntaoRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

func (r AlibabaInteractSupplierAwardResourceGetCuntaoRequest) GetUserNick() string {
    return r.userNick
}

func (r *AlibabaInteractSupplierAwardResourceGetCuntaoRequest) SetActivityKey(activityKey string) error {
    r.activityKey = activityKey
    r.Set("activity_key", activityKey)
    return nil
}

func (r AlibabaInteractSupplierAwardResourceGetCuntaoRequest) GetActivityKey() string {
    return r.activityKey
}

func (r *AlibabaInteractSupplierAwardResourceGetCuntaoRequest) SetLng(lng string) error {
    r.lng = lng
    r.Set("lng", lng)
    return nil
}

func (r AlibabaInteractSupplierAwardResourceGetCuntaoRequest) GetLng() string {
    return r.lng
}

func (r *AlibabaInteractSupplierAwardResourceGetCuntaoRequest) SetLat(lat string) error {
    r.lat = lat
    r.Set("lat", lat)
    return nil
}

func (r AlibabaInteractSupplierAwardResourceGetCuntaoRequest) GetLat() string {
    return r.lat
}

