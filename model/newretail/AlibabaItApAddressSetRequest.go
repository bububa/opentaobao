package newretail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
setApAddressNew APIRequest
alibaba.it.ap.address.set

该接口可为ISV系统提供 ap位置信息维护的功能
*/
type AlibabaItApAddressSetRequest struct {
    model.Params

    // 城市
    apCityName   string 

    // 经度
    lng   string 

    // 签名
    signature   string 

    // 园区/门店
    apCampusName   string 

    // 区域
    apAreaName   string 

    // 省份
    apProvinceName   string 

    // ap mac地址
    mac   string 

    // ap空间单元名称
    apUnitName   string 

    // 楼层
    apFloor   string 

    // 楼栋
    apBuildingName   string 

    // 分配的内部ak
    appKeyInternal   string 

    // 国家
    apNationName   string 

    // 纬度
    lat   string 

    // 方位
    direction   string 

    // 时间戳，毫秒
    timestampInternal   int64 

}

func NewAlibabaItApAddressSetRequest() *AlibabaItApAddressSetRequest{
    return &AlibabaItApAddressSetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaItApAddressSetRequest) GetApiMethodName() string {
    return "alibaba.it.ap.address.set"
}

func (r AlibabaItApAddressSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaItApAddressSetRequest) SetApCityName(apCityName string) error {
    r.apCityName = apCityName
    r.Set("ap_city_name", apCityName)
    return nil
}

func (r AlibabaItApAddressSetRequest) GetApCityName() string {
    return r.apCityName
}

func (r *AlibabaItApAddressSetRequest) SetLng(lng string) error {
    r.lng = lng
    r.Set("lng", lng)
    return nil
}

func (r AlibabaItApAddressSetRequest) GetLng() string {
    return r.lng
}

func (r *AlibabaItApAddressSetRequest) SetSignature(signature string) error {
    r.signature = signature
    r.Set("signature", signature)
    return nil
}

func (r AlibabaItApAddressSetRequest) GetSignature() string {
    return r.signature
}

func (r *AlibabaItApAddressSetRequest) SetApCampusName(apCampusName string) error {
    r.apCampusName = apCampusName
    r.Set("ap_campus_name", apCampusName)
    return nil
}

func (r AlibabaItApAddressSetRequest) GetApCampusName() string {
    return r.apCampusName
}

func (r *AlibabaItApAddressSetRequest) SetApAreaName(apAreaName string) error {
    r.apAreaName = apAreaName
    r.Set("ap_area_name", apAreaName)
    return nil
}

func (r AlibabaItApAddressSetRequest) GetApAreaName() string {
    return r.apAreaName
}

func (r *AlibabaItApAddressSetRequest) SetApProvinceName(apProvinceName string) error {
    r.apProvinceName = apProvinceName
    r.Set("ap_province_name", apProvinceName)
    return nil
}

func (r AlibabaItApAddressSetRequest) GetApProvinceName() string {
    return r.apProvinceName
}

func (r *AlibabaItApAddressSetRequest) SetMac(mac string) error {
    r.mac = mac
    r.Set("mac", mac)
    return nil
}

func (r AlibabaItApAddressSetRequest) GetMac() string {
    return r.mac
}

func (r *AlibabaItApAddressSetRequest) SetApUnitName(apUnitName string) error {
    r.apUnitName = apUnitName
    r.Set("ap_unit_name", apUnitName)
    return nil
}

func (r AlibabaItApAddressSetRequest) GetApUnitName() string {
    return r.apUnitName
}

func (r *AlibabaItApAddressSetRequest) SetApFloor(apFloor string) error {
    r.apFloor = apFloor
    r.Set("ap_floor", apFloor)
    return nil
}

func (r AlibabaItApAddressSetRequest) GetApFloor() string {
    return r.apFloor
}

func (r *AlibabaItApAddressSetRequest) SetApBuildingName(apBuildingName string) error {
    r.apBuildingName = apBuildingName
    r.Set("ap_building_name", apBuildingName)
    return nil
}

func (r AlibabaItApAddressSetRequest) GetApBuildingName() string {
    return r.apBuildingName
}

func (r *AlibabaItApAddressSetRequest) SetAppKeyInternal(appKeyInternal string) error {
    r.appKeyInternal = appKeyInternal
    r.Set("app_key_internal", appKeyInternal)
    return nil
}

func (r AlibabaItApAddressSetRequest) GetAppKeyInternal() string {
    return r.appKeyInternal
}

func (r *AlibabaItApAddressSetRequest) SetApNationName(apNationName string) error {
    r.apNationName = apNationName
    r.Set("ap_nation_name", apNationName)
    return nil
}

func (r AlibabaItApAddressSetRequest) GetApNationName() string {
    return r.apNationName
}

func (r *AlibabaItApAddressSetRequest) SetLat(lat string) error {
    r.lat = lat
    r.Set("lat", lat)
    return nil
}

func (r AlibabaItApAddressSetRequest) GetLat() string {
    return r.lat
}

func (r *AlibabaItApAddressSetRequest) SetDirection(direction string) error {
    r.direction = direction
    r.Set("direction", direction)
    return nil
}

func (r AlibabaItApAddressSetRequest) GetDirection() string {
    return r.direction
}

func (r *AlibabaItApAddressSetRequest) SetTimestampInternal(timestampInternal int64) error {
    r.timestampInternal = timestampInternal
    r.Set("timestamp_internal", timestampInternal)
    return nil
}

func (r AlibabaItApAddressSetRequest) GetTimestampInternal() int64 {
    return r.timestampInternal
}

