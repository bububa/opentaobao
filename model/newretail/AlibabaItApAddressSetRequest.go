package newretail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
setApAddressNew API请求
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

// 初始化AlibabaItApAddressSetRequest对象
func NewAlibabaItApAddressSetRequest() *AlibabaItApAddressSetRequest{
    return &AlibabaItApAddressSetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaItApAddressSetRequest) GetApiMethodName() string {
    return "alibaba.it.ap.address.set"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaItApAddressSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApCityName Setter
// 城市
func (r *AlibabaItApAddressSetRequest) SetApCityName(apCityName string) error {
    r.apCityName = apCityName
    r.Set("ap_city_name", apCityName)
    return nil
}

// ApCityName Getter
func (r AlibabaItApAddressSetRequest) GetApCityName() string {
    return r.apCityName
}
// Lng Setter
// 经度
func (r *AlibabaItApAddressSetRequest) SetLng(lng string) error {
    r.lng = lng
    r.Set("lng", lng)
    return nil
}

// Lng Getter
func (r AlibabaItApAddressSetRequest) GetLng() string {
    return r.lng
}
// Signature Setter
// 签名
func (r *AlibabaItApAddressSetRequest) SetSignature(signature string) error {
    r.signature = signature
    r.Set("signature", signature)
    return nil
}

// Signature Getter
func (r AlibabaItApAddressSetRequest) GetSignature() string {
    return r.signature
}
// ApCampusName Setter
// 园区/门店
func (r *AlibabaItApAddressSetRequest) SetApCampusName(apCampusName string) error {
    r.apCampusName = apCampusName
    r.Set("ap_campus_name", apCampusName)
    return nil
}

// ApCampusName Getter
func (r AlibabaItApAddressSetRequest) GetApCampusName() string {
    return r.apCampusName
}
// ApAreaName Setter
// 区域
func (r *AlibabaItApAddressSetRequest) SetApAreaName(apAreaName string) error {
    r.apAreaName = apAreaName
    r.Set("ap_area_name", apAreaName)
    return nil
}

// ApAreaName Getter
func (r AlibabaItApAddressSetRequest) GetApAreaName() string {
    return r.apAreaName
}
// ApProvinceName Setter
// 省份
func (r *AlibabaItApAddressSetRequest) SetApProvinceName(apProvinceName string) error {
    r.apProvinceName = apProvinceName
    r.Set("ap_province_name", apProvinceName)
    return nil
}

// ApProvinceName Getter
func (r AlibabaItApAddressSetRequest) GetApProvinceName() string {
    return r.apProvinceName
}
// Mac Setter
// ap mac地址
func (r *AlibabaItApAddressSetRequest) SetMac(mac string) error {
    r.mac = mac
    r.Set("mac", mac)
    return nil
}

// Mac Getter
func (r AlibabaItApAddressSetRequest) GetMac() string {
    return r.mac
}
// ApUnitName Setter
// ap空间单元名称
func (r *AlibabaItApAddressSetRequest) SetApUnitName(apUnitName string) error {
    r.apUnitName = apUnitName
    r.Set("ap_unit_name", apUnitName)
    return nil
}

// ApUnitName Getter
func (r AlibabaItApAddressSetRequest) GetApUnitName() string {
    return r.apUnitName
}
// ApFloor Setter
// 楼层
func (r *AlibabaItApAddressSetRequest) SetApFloor(apFloor string) error {
    r.apFloor = apFloor
    r.Set("ap_floor", apFloor)
    return nil
}

// ApFloor Getter
func (r AlibabaItApAddressSetRequest) GetApFloor() string {
    return r.apFloor
}
// ApBuildingName Setter
// 楼栋
func (r *AlibabaItApAddressSetRequest) SetApBuildingName(apBuildingName string) error {
    r.apBuildingName = apBuildingName
    r.Set("ap_building_name", apBuildingName)
    return nil
}

// ApBuildingName Getter
func (r AlibabaItApAddressSetRequest) GetApBuildingName() string {
    return r.apBuildingName
}
// AppKeyInternal Setter
// 分配的内部ak
func (r *AlibabaItApAddressSetRequest) SetAppKeyInternal(appKeyInternal string) error {
    r.appKeyInternal = appKeyInternal
    r.Set("app_key_internal", appKeyInternal)
    return nil
}

// AppKeyInternal Getter
func (r AlibabaItApAddressSetRequest) GetAppKeyInternal() string {
    return r.appKeyInternal
}
// ApNationName Setter
// 国家
func (r *AlibabaItApAddressSetRequest) SetApNationName(apNationName string) error {
    r.apNationName = apNationName
    r.Set("ap_nation_name", apNationName)
    return nil
}

// ApNationName Getter
func (r AlibabaItApAddressSetRequest) GetApNationName() string {
    return r.apNationName
}
// Lat Setter
// 纬度
func (r *AlibabaItApAddressSetRequest) SetLat(lat string) error {
    r.lat = lat
    r.Set("lat", lat)
    return nil
}

// Lat Getter
func (r AlibabaItApAddressSetRequest) GetLat() string {
    return r.lat
}
// Direction Setter
// 方位
func (r *AlibabaItApAddressSetRequest) SetDirection(direction string) error {
    r.direction = direction
    r.Set("direction", direction)
    return nil
}

// Direction Getter
func (r AlibabaItApAddressSetRequest) GetDirection() string {
    return r.direction
}
// TimestampInternal Setter
// 时间戳，毫秒
func (r *AlibabaItApAddressSetRequest) SetTimestampInternal(timestampInternal int64) error {
    r.timestampInternal = timestampInternal
    r.Set("timestamp_internal", timestampInternal)
    return nil
}

// TimestampInternal Getter
func (r AlibabaItApAddressSetRequest) GetTimestampInternal() int64 {
    return r.timestampInternal
}
