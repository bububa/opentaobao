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
    _apCityName   string
    // 经度
    _lng   string
    // 签名
    _signature   string
    // 园区/门店
    _apCampusName   string
    // 区域
    _apAreaName   string
    // 省份
    _apProvinceName   string
    // ap mac地址
    _mac   string
    // ap空间单元名称
    _apUnitName   string
    // 楼层
    _apFloor   string
    // 楼栋
    _apBuildingName   string
    // 分配的内部ak
    _appKeyInternal   string
    // 国家
    _apNationName   string
    // 纬度
    _lat   string
    // 方位
    _direction   string
    // 时间戳，毫秒
    _timestampInternal   int64
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
func (r *AlibabaItApAddressSetRequest) SetApCityName(_apCityName string) error {
    r._apCityName = _apCityName
    r.Set("ap_city_name", _apCityName)
    return nil
}

// ApCityName Getter
func (r AlibabaItApAddressSetRequest) GetApCityName() string {
    return r._apCityName
}
// Lng Setter
// 经度
func (r *AlibabaItApAddressSetRequest) SetLng(_lng string) error {
    r._lng = _lng
    r.Set("lng", _lng)
    return nil
}

// Lng Getter
func (r AlibabaItApAddressSetRequest) GetLng() string {
    return r._lng
}
// Signature Setter
// 签名
func (r *AlibabaItApAddressSetRequest) SetSignature(_signature string) error {
    r._signature = _signature
    r.Set("signature", _signature)
    return nil
}

// Signature Getter
func (r AlibabaItApAddressSetRequest) GetSignature() string {
    return r._signature
}
// ApCampusName Setter
// 园区/门店
func (r *AlibabaItApAddressSetRequest) SetApCampusName(_apCampusName string) error {
    r._apCampusName = _apCampusName
    r.Set("ap_campus_name", _apCampusName)
    return nil
}

// ApCampusName Getter
func (r AlibabaItApAddressSetRequest) GetApCampusName() string {
    return r._apCampusName
}
// ApAreaName Setter
// 区域
func (r *AlibabaItApAddressSetRequest) SetApAreaName(_apAreaName string) error {
    r._apAreaName = _apAreaName
    r.Set("ap_area_name", _apAreaName)
    return nil
}

// ApAreaName Getter
func (r AlibabaItApAddressSetRequest) GetApAreaName() string {
    return r._apAreaName
}
// ApProvinceName Setter
// 省份
func (r *AlibabaItApAddressSetRequest) SetApProvinceName(_apProvinceName string) error {
    r._apProvinceName = _apProvinceName
    r.Set("ap_province_name", _apProvinceName)
    return nil
}

// ApProvinceName Getter
func (r AlibabaItApAddressSetRequest) GetApProvinceName() string {
    return r._apProvinceName
}
// Mac Setter
// ap mac地址
func (r *AlibabaItApAddressSetRequest) SetMac(_mac string) error {
    r._mac = _mac
    r.Set("mac", _mac)
    return nil
}

// Mac Getter
func (r AlibabaItApAddressSetRequest) GetMac() string {
    return r._mac
}
// ApUnitName Setter
// ap空间单元名称
func (r *AlibabaItApAddressSetRequest) SetApUnitName(_apUnitName string) error {
    r._apUnitName = _apUnitName
    r.Set("ap_unit_name", _apUnitName)
    return nil
}

// ApUnitName Getter
func (r AlibabaItApAddressSetRequest) GetApUnitName() string {
    return r._apUnitName
}
// ApFloor Setter
// 楼层
func (r *AlibabaItApAddressSetRequest) SetApFloor(_apFloor string) error {
    r._apFloor = _apFloor
    r.Set("ap_floor", _apFloor)
    return nil
}

// ApFloor Getter
func (r AlibabaItApAddressSetRequest) GetApFloor() string {
    return r._apFloor
}
// ApBuildingName Setter
// 楼栋
func (r *AlibabaItApAddressSetRequest) SetApBuildingName(_apBuildingName string) error {
    r._apBuildingName = _apBuildingName
    r.Set("ap_building_name", _apBuildingName)
    return nil
}

// ApBuildingName Getter
func (r AlibabaItApAddressSetRequest) GetApBuildingName() string {
    return r._apBuildingName
}
// AppKeyInternal Setter
// 分配的内部ak
func (r *AlibabaItApAddressSetRequest) SetAppKeyInternal(_appKeyInternal string) error {
    r._appKeyInternal = _appKeyInternal
    r.Set("app_key_internal", _appKeyInternal)
    return nil
}

// AppKeyInternal Getter
func (r AlibabaItApAddressSetRequest) GetAppKeyInternal() string {
    return r._appKeyInternal
}
// ApNationName Setter
// 国家
func (r *AlibabaItApAddressSetRequest) SetApNationName(_apNationName string) error {
    r._apNationName = _apNationName
    r.Set("ap_nation_name", _apNationName)
    return nil
}

// ApNationName Getter
func (r AlibabaItApAddressSetRequest) GetApNationName() string {
    return r._apNationName
}
// Lat Setter
// 纬度
func (r *AlibabaItApAddressSetRequest) SetLat(_lat string) error {
    r._lat = _lat
    r.Set("lat", _lat)
    return nil
}

// Lat Getter
func (r AlibabaItApAddressSetRequest) GetLat() string {
    return r._lat
}
// Direction Setter
// 方位
func (r *AlibabaItApAddressSetRequest) SetDirection(_direction string) error {
    r._direction = _direction
    r.Set("direction", _direction)
    return nil
}

// Direction Getter
func (r AlibabaItApAddressSetRequest) GetDirection() string {
    return r._direction
}
// TimestampInternal Setter
// 时间戳，毫秒
func (r *AlibabaItApAddressSetRequest) SetTimestampInternal(_timestampInternal int64) error {
    r._timestampInternal = _timestampInternal
    r.Set("timestamp_internal", _timestampInternal)
    return nil
}

// TimestampInternal Getter
func (r AlibabaItApAddressSetRequest) GetTimestampInternal() int64 {
    return r._timestampInternal
}
