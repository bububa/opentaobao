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
type AlibabaInteractSupplierAwardResourceGetCuntaoAPIRequest struct {
    model.Params
    // 用户昵称
    _userNick   string
    // 活动code
    _activityKey   string
    // 经度
    _lng   string
    // 纬度
    _lat   string
}

// 初始化AlibabaInteractSupplierAwardResourceGetCuntaoAPIRequest对象
func NewAlibabaInteractSupplierAwardResourceGetCuntaoRequest() *AlibabaInteractSupplierAwardResourceGetCuntaoAPIRequest{
    return &AlibabaInteractSupplierAwardResourceGetCuntaoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractSupplierAwardResourceGetCuntaoAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.supplier.award.resource.get.cuntao"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractSupplierAwardResourceGetCuntaoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserNick Setter
// 用户昵称
func (r *AlibabaInteractSupplierAwardResourceGetCuntaoAPIRequest) SetUserNick(_userNick string) error {
    r._userNick = _userNick
    r.Set("user_nick", _userNick)
    return nil
}

// UserNick Getter
func (r AlibabaInteractSupplierAwardResourceGetCuntaoAPIRequest) GetUserNick() string {
    return r._userNick
}
// ActivityKey Setter
// 活动code
func (r *AlibabaInteractSupplierAwardResourceGetCuntaoAPIRequest) SetActivityKey(_activityKey string) error {
    r._activityKey = _activityKey
    r.Set("activity_key", _activityKey)
    return nil
}

// ActivityKey Getter
func (r AlibabaInteractSupplierAwardResourceGetCuntaoAPIRequest) GetActivityKey() string {
    return r._activityKey
}
// Lng Setter
// 经度
func (r *AlibabaInteractSupplierAwardResourceGetCuntaoAPIRequest) SetLng(_lng string) error {
    r._lng = _lng
    r.Set("lng", _lng)
    return nil
}

// Lng Getter
func (r AlibabaInteractSupplierAwardResourceGetCuntaoAPIRequest) GetLng() string {
    return r._lng
}
// Lat Setter
// 纬度
func (r *AlibabaInteractSupplierAwardResourceGetCuntaoAPIRequest) SetLat(_lat string) error {
    r._lat = _lat
    r.Set("lat", _lat)
    return nil
}

// Lat Getter
func (r AlibabaInteractSupplierAwardResourceGetCuntaoAPIRequest) GetLat() string {
    return r._lat
}
