package smartstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSmartstoreDeviceAddAPIRequest 智慧门店设备创建 API请求
// taobao.smartstore.device.add
//
// 智慧门店设备创建
type TaobaoSmartstoreDeviceAddAPIRequest struct {
	model.Params
	// mac地址
	_mac string
	// 设备在店内的位置，以文字描述
	_indoorPosition string
	// 设备名称
	_deviceName string
	// 门店ID
	_storeId int64
	// 操作系统类型：WINDOWS("WINDOWS", "WINDOWS"),     ANDROID("ANDROID", "ANDROID"),     IOS("IOS", "IOS"),     LINUX("LINUX", "LINUX"),     OTHER("OTHER", "OTHER");
	_osType string
	// 设备类型：     CAMERA("CAMERA", "客流摄像头"),     SHELF("SHELF", "云货架"),     MAKEUP_MIRROR("MAKEUP_MIRROR", "试妆镜"),     FITTING_MIRROR("FITTING_MIRROR", "试衣镜"),     VENDOR("VENDOR", "售货机"),     SAMPLE_MACHINE("SAMPLE_MACHINE","派样机"),     DOLL_MACHINE("DOLL_MACHINE", "娃娃机"),     INTERACTIVE_PHOTO("INTERACTIVE_PHOTO", "互动拍照"),     INTERACTIVE_GAME("INTERACTIVE_GAME", "互动游戏"),     USHER_SCREEN("USHER_SCREEN", "智慧迎宾屏"),     DRESSING("DRESSING", "闪电换装"),     MAGIC_MIRROR("MAGIC_MIRROR", "百搭魔镜"),     SHOES_FITTING_MIRROR("SHOES_FITTING_MIRROR", "试鞋镜"),     SKIN_DETECTION("SKIN_DETECTION", "肌肤测试仪"),     FOOT_DETECTION("FOOT_DETECTION", "测脚仪"),     RFID_SENSOR("RFID_SENSOR", "RFID"),touch_machine("touch_machine","导购一体屏")
	_deviceType string
	// 商家自定义设备编码
	_outerCode string
}

// NewTaobaoSmartstoreDeviceAddRequest 初始化TaobaoSmartstoreDeviceAddAPIRequest对象
func NewTaobaoSmartstoreDeviceAddRequest() *TaobaoSmartstoreDeviceAddAPIRequest {
	return &TaobaoSmartstoreDeviceAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSmartstoreDeviceAddAPIRequest) GetApiMethodName() string {
	return "taobao.smartstore.device.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSmartstoreDeviceAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Mac Setter
// mac地址
func (r *TaobaoSmartstoreDeviceAddAPIRequest) SetMac(_mac string) error {
	r._mac = _mac
	r.Set("mac", _mac)
	return nil
}

// Get Mac Getter
func (r TaobaoSmartstoreDeviceAddAPIRequest) GetMac() string {
	return r._mac
}

// Set is IndoorPosition Setter
// 设备在店内的位置，以文字描述
func (r *TaobaoSmartstoreDeviceAddAPIRequest) SetIndoorPosition(_indoorPosition string) error {
	r._indoorPosition = _indoorPosition
	r.Set("indoor_position", _indoorPosition)
	return nil
}

// Get IndoorPosition Getter
func (r TaobaoSmartstoreDeviceAddAPIRequest) GetIndoorPosition() string {
	return r._indoorPosition
}

// Set is DeviceName Setter
// 设备名称
func (r *TaobaoSmartstoreDeviceAddAPIRequest) SetDeviceName(_deviceName string) error {
	r._deviceName = _deviceName
	r.Set("device_name", _deviceName)
	return nil
}

// Get DeviceName Getter
func (r TaobaoSmartstoreDeviceAddAPIRequest) GetDeviceName() string {
	return r._deviceName
}

// Set is StoreId Setter
// 门店ID
func (r *TaobaoSmartstoreDeviceAddAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r TaobaoSmartstoreDeviceAddAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// Set is OsType Setter
// 操作系统类型：WINDOWS("WINDOWS", "WINDOWS"),     ANDROID("ANDROID", "ANDROID"),     IOS("IOS", "IOS"),     LINUX("LINUX", "LINUX"),     OTHER("OTHER", "OTHER");
func (r *TaobaoSmartstoreDeviceAddAPIRequest) SetOsType(_osType string) error {
	r._osType = _osType
	r.Set("os_type", _osType)
	return nil
}

// Get OsType Getter
func (r TaobaoSmartstoreDeviceAddAPIRequest) GetOsType() string {
	return r._osType
}

// Set is DeviceType Setter
// 设备类型：     CAMERA("CAMERA", "客流摄像头"),     SHELF("SHELF", "云货架"),     MAKEUP_MIRROR("MAKEUP_MIRROR", "试妆镜"),     FITTING_MIRROR("FITTING_MIRROR", "试衣镜"),     VENDOR("VENDOR", "售货机"),     SAMPLE_MACHINE("SAMPLE_MACHINE","派样机"),     DOLL_MACHINE("DOLL_MACHINE", "娃娃机"),     INTERACTIVE_PHOTO("INTERACTIVE_PHOTO", "互动拍照"),     INTERACTIVE_GAME("INTERACTIVE_GAME", "互动游戏"),     USHER_SCREEN("USHER_SCREEN", "智慧迎宾屏"),     DRESSING("DRESSING", "闪电换装"),     MAGIC_MIRROR("MAGIC_MIRROR", "百搭魔镜"),     SHOES_FITTING_MIRROR("SHOES_FITTING_MIRROR", "试鞋镜"),     SKIN_DETECTION("SKIN_DETECTION", "肌肤测试仪"),     FOOT_DETECTION("FOOT_DETECTION", "测脚仪"),     RFID_SENSOR("RFID_SENSOR", "RFID"),touch_machine("touch_machine","导购一体屏")
func (r *TaobaoSmartstoreDeviceAddAPIRequest) SetDeviceType(_deviceType string) error {
	r._deviceType = _deviceType
	r.Set("device_type", _deviceType)
	return nil
}

// Get DeviceType Getter
func (r TaobaoSmartstoreDeviceAddAPIRequest) GetDeviceType() string {
	return r._deviceType
}

// Set is OuterCode Setter
// 商家自定义设备编码
func (r *TaobaoSmartstoreDeviceAddAPIRequest) SetOuterCode(_outerCode string) error {
	r._outerCode = _outerCode
	r.Set("outer_code", _outerCode)
	return nil
}

// Get OuterCode Getter
func (r TaobaoSmartstoreDeviceAddAPIRequest) GetOuterCode() string {
	return r._outerCode
}
