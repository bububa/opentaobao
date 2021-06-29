package util

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
权益池信息查询 API请求
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

// 初始化AlibabaInteractSupplierAwardResourceGetCuntaoRequest对象
func NewAlibabaInteractSupplierAwardResourceGetCuntaoRequest() *AlibabaInteractSupplierAwardResourceGetCuntaoRequest{
    return &AlibabaInteractSupplierAwardResourceGetCuntaoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSupplierAwardResourceGetCuntaoRequest) GetApiMethodName() string {
    return "alibaba.interact.supplier.award.resource.get.cuntao"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSupplierAwardResourceGetCuntaoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserNick Setter
// 用户昵称
func (r *AlibabaInteractSupplierAwardResourceGetCuntaoRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

// UserNick Getter
func (r AlibabaInteractSupplierAwardResourceGetCuntaoRequest) GetUserNick() string {
    return r.userNick
}
// ActivityKey Setter
// 活动code
func (r *AlibabaInteractSupplierAwardResourceGetCuntaoRequest) SetActivityKey(activityKey string) error {
    r.activityKey = activityKey
    r.Set("activity_key", activityKey)
    return nil
}

// ActivityKey Getter
func (r AlibabaInteractSupplierAwardResourceGetCuntaoRequest) GetActivityKey() string {
    return r.activityKey
}
// Lng Setter
// 经度
func (r *AlibabaInteractSupplierAwardResourceGetCuntaoRequest) SetLng(lng string) error {
    r.lng = lng
    r.Set("lng", lng)
    return nil
}

// Lng Getter
func (r AlibabaInteractSupplierAwardResourceGetCuntaoRequest) GetLng() string {
    return r.lng
}
// Lat Setter
// 纬度
func (r *AlibabaInteractSupplierAwardResourceGetCuntaoRequest) SetLat(lat string) error {
    r.lat = lat
    r.Set("lat", lat)
    return nil
}

// Lat Getter
func (r AlibabaInteractSupplierAwardResourceGetCuntaoRequest) GetLat() string {
    return r.lat
}
