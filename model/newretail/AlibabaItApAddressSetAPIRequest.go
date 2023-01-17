package newretail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItApAddressSetAPIRequest setApAddressNew API请求
// alibaba.it.ap.address.set
//
// 该接口可为ISV系统提供 ap位置信息维护的功能
type AlibabaItApAddressSetAPIRequest struct {
	model.Params
	// 城市
	_apCityName string
	// 经度
	_lng string
	// 签名
	_signature string
	// 园区/门店
	_apCampusName string
	// 区域
	_apAreaName string
	// 省份
	_apProvinceName string
	// ap mac地址
	_mac string
	// ap空间单元名称
	_apUnitName string
	// 楼层
	_apFloor string
	// 楼栋
	_apBuildingName string
	// 分配的内部ak
	_appKeyInternal string
	// 国家
	_apNationName string
	// 纬度
	_lat string
	// 方位
	_direction string
	// 时间戳，毫秒
	_timestampInternal int64
}

// NewAlibabaItApAddressSetRequest 初始化AlibabaItApAddressSetAPIRequest对象
func NewAlibabaItApAddressSetRequest() *AlibabaItApAddressSetAPIRequest {
	return &AlibabaItApAddressSetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItApAddressSetAPIRequest) GetApiMethodName() string {
	return "alibaba.it.ap.address.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItApAddressSetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaItApAddressSetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApCityName is ApCityName Setter
// 城市
func (r *AlibabaItApAddressSetAPIRequest) SetApCityName(_apCityName string) error {
	r._apCityName = _apCityName
	r.Set("ap_city_name", _apCityName)
	return nil
}

// GetApCityName ApCityName Getter
func (r AlibabaItApAddressSetAPIRequest) GetApCityName() string {
	return r._apCityName
}

// SetLng is Lng Setter
// 经度
func (r *AlibabaItApAddressSetAPIRequest) SetLng(_lng string) error {
	r._lng = _lng
	r.Set("lng", _lng)
	return nil
}

// GetLng Lng Getter
func (r AlibabaItApAddressSetAPIRequest) GetLng() string {
	return r._lng
}

// SetSignature is Signature Setter
// 签名
func (r *AlibabaItApAddressSetAPIRequest) SetSignature(_signature string) error {
	r._signature = _signature
	r.Set("signature", _signature)
	return nil
}

// GetSignature Signature Getter
func (r AlibabaItApAddressSetAPIRequest) GetSignature() string {
	return r._signature
}

// SetApCampusName is ApCampusName Setter
// 园区/门店
func (r *AlibabaItApAddressSetAPIRequest) SetApCampusName(_apCampusName string) error {
	r._apCampusName = _apCampusName
	r.Set("ap_campus_name", _apCampusName)
	return nil
}

// GetApCampusName ApCampusName Getter
func (r AlibabaItApAddressSetAPIRequest) GetApCampusName() string {
	return r._apCampusName
}

// SetApAreaName is ApAreaName Setter
// 区域
func (r *AlibabaItApAddressSetAPIRequest) SetApAreaName(_apAreaName string) error {
	r._apAreaName = _apAreaName
	r.Set("ap_area_name", _apAreaName)
	return nil
}

// GetApAreaName ApAreaName Getter
func (r AlibabaItApAddressSetAPIRequest) GetApAreaName() string {
	return r._apAreaName
}

// SetApProvinceName is ApProvinceName Setter
// 省份
func (r *AlibabaItApAddressSetAPIRequest) SetApProvinceName(_apProvinceName string) error {
	r._apProvinceName = _apProvinceName
	r.Set("ap_province_name", _apProvinceName)
	return nil
}

// GetApProvinceName ApProvinceName Getter
func (r AlibabaItApAddressSetAPIRequest) GetApProvinceName() string {
	return r._apProvinceName
}

// SetMac is Mac Setter
// ap mac地址
func (r *AlibabaItApAddressSetAPIRequest) SetMac(_mac string) error {
	r._mac = _mac
	r.Set("mac", _mac)
	return nil
}

// GetMac Mac Getter
func (r AlibabaItApAddressSetAPIRequest) GetMac() string {
	return r._mac
}

// SetApUnitName is ApUnitName Setter
// ap空间单元名称
func (r *AlibabaItApAddressSetAPIRequest) SetApUnitName(_apUnitName string) error {
	r._apUnitName = _apUnitName
	r.Set("ap_unit_name", _apUnitName)
	return nil
}

// GetApUnitName ApUnitName Getter
func (r AlibabaItApAddressSetAPIRequest) GetApUnitName() string {
	return r._apUnitName
}

// SetApFloor is ApFloor Setter
// 楼层
func (r *AlibabaItApAddressSetAPIRequest) SetApFloor(_apFloor string) error {
	r._apFloor = _apFloor
	r.Set("ap_floor", _apFloor)
	return nil
}

// GetApFloor ApFloor Getter
func (r AlibabaItApAddressSetAPIRequest) GetApFloor() string {
	return r._apFloor
}

// SetApBuildingName is ApBuildingName Setter
// 楼栋
func (r *AlibabaItApAddressSetAPIRequest) SetApBuildingName(_apBuildingName string) error {
	r._apBuildingName = _apBuildingName
	r.Set("ap_building_name", _apBuildingName)
	return nil
}

// GetApBuildingName ApBuildingName Getter
func (r AlibabaItApAddressSetAPIRequest) GetApBuildingName() string {
	return r._apBuildingName
}

// SetAppKeyInternal is AppKeyInternal Setter
// 分配的内部ak
func (r *AlibabaItApAddressSetAPIRequest) SetAppKeyInternal(_appKeyInternal string) error {
	r._appKeyInternal = _appKeyInternal
	r.Set("app_key_internal", _appKeyInternal)
	return nil
}

// GetAppKeyInternal AppKeyInternal Getter
func (r AlibabaItApAddressSetAPIRequest) GetAppKeyInternal() string {
	return r._appKeyInternal
}

// SetApNationName is ApNationName Setter
// 国家
func (r *AlibabaItApAddressSetAPIRequest) SetApNationName(_apNationName string) error {
	r._apNationName = _apNationName
	r.Set("ap_nation_name", _apNationName)
	return nil
}

// GetApNationName ApNationName Getter
func (r AlibabaItApAddressSetAPIRequest) GetApNationName() string {
	return r._apNationName
}

// SetLat is Lat Setter
// 纬度
func (r *AlibabaItApAddressSetAPIRequest) SetLat(_lat string) error {
	r._lat = _lat
	r.Set("lat", _lat)
	return nil
}

// GetLat Lat Getter
func (r AlibabaItApAddressSetAPIRequest) GetLat() string {
	return r._lat
}

// SetDirection is Direction Setter
// 方位
func (r *AlibabaItApAddressSetAPIRequest) SetDirection(_direction string) error {
	r._direction = _direction
	r.Set("direction", _direction)
	return nil
}

// GetDirection Direction Getter
func (r AlibabaItApAddressSetAPIRequest) GetDirection() string {
	return r._direction
}

// SetTimestampInternal is TimestampInternal Setter
// 时间戳，毫秒
func (r *AlibabaItApAddressSetAPIRequest) SetTimestampInternal(_timestampInternal int64) error {
	r._timestampInternal = _timestampInternal
	r.Set("timestamp_internal", _timestampInternal)
	return nil
}

// GetTimestampInternal TimestampInternal Getter
func (r AlibabaItApAddressSetAPIRequest) GetTimestampInternal() int64 {
	return r._timestampInternal
}
