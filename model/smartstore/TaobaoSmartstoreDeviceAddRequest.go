package smartstore

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
智慧门店设备创建 API请求
taobao.smartstore.device.add

智慧门店设备创建
*/
type TaobaoSmartstoreDeviceAddRequest struct {
    model.Params
    // mac地址
    _mac   string
    // 设备在店内的位置，以文字描述
    _indoorPosition   string
    // 设备名称
    _deviceName   string
    // 门店ID
    _storeId   int64
    // 操作系统类型：WINDOWS("WINDOWS", "WINDOWS"),     ANDROID("ANDROID", "ANDROID"),     IOS("IOS", "IOS"),     LINUX("LINUX", "LINUX"),     OTHER("OTHER", "OTHER");
    _osType   string
    // 设备类型：     CAMERA("CAMERA", "客流摄像头"),     SHELF("SHELF", "云货架"),     MAKEUP_MIRROR("MAKEUP_MIRROR", "试妆镜"),     FITTING_MIRROR("FITTING_MIRROR", "试衣镜"),     VENDOR("VENDOR", "售货机"),     SAMPLE_MACHINE("SAMPLE_MACHINE","派样机"),     DOLL_MACHINE("DOLL_MACHINE", "娃娃机"),     INTERACTIVE_PHOTO("INTERACTIVE_PHOTO", "互动拍照"),     INTERACTIVE_GAME("INTERACTIVE_GAME", "互动游戏"),     USHER_SCREEN("USHER_SCREEN", "智慧迎宾屏"),     DRESSING("DRESSING", "闪电换装"),     MAGIC_MIRROR("MAGIC_MIRROR", "百搭魔镜"),     SHOES_FITTING_MIRROR("SHOES_FITTING_MIRROR", "试鞋镜"),     SKIN_DETECTION("SKIN_DETECTION", "肌肤测试仪"),     FOOT_DETECTION("FOOT_DETECTION", "测脚仪"),     RFID_SENSOR("RFID_SENSOR", "RFID"),touch_machine("touch_machine","导购一体屏")
    _deviceType   string
    // 商家自定义设备编码
    _outerCode   string
}

// 初始化TaobaoSmartstoreDeviceAddRequest对象
func NewTaobaoSmartstoreDeviceAddRequest() *TaobaoSmartstoreDeviceAddRequest{
    return &TaobaoSmartstoreDeviceAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSmartstoreDeviceAddRequest) GetApiMethodName() string {
    return "taobao.smartstore.device.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSmartstoreDeviceAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Mac Setter
// mac地址
func (r *TaobaoSmartstoreDeviceAddRequest) SetMac(_mac string) error {
    r._mac = _mac
    r.Set("mac", _mac)
    return nil
}

// Mac Getter
func (r TaobaoSmartstoreDeviceAddRequest) GetMac() string {
    return r._mac
}
// IndoorPosition Setter
// 设备在店内的位置，以文字描述
func (r *TaobaoSmartstoreDeviceAddRequest) SetIndoorPosition(_indoorPosition string) error {
    r._indoorPosition = _indoorPosition
    r.Set("indoor_position", _indoorPosition)
    return nil
}

// IndoorPosition Getter
func (r TaobaoSmartstoreDeviceAddRequest) GetIndoorPosition() string {
    return r._indoorPosition
}
// DeviceName Setter
// 设备名称
func (r *TaobaoSmartstoreDeviceAddRequest) SetDeviceName(_deviceName string) error {
    r._deviceName = _deviceName
    r.Set("device_name", _deviceName)
    return nil
}

// DeviceName Getter
func (r TaobaoSmartstoreDeviceAddRequest) GetDeviceName() string {
    return r._deviceName
}
// StoreId Setter
// 门店ID
func (r *TaobaoSmartstoreDeviceAddRequest) SetStoreId(_storeId int64) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r TaobaoSmartstoreDeviceAddRequest) GetStoreId() int64 {
    return r._storeId
}
// OsType Setter
// 操作系统类型：WINDOWS("WINDOWS", "WINDOWS"),     ANDROID("ANDROID", "ANDROID"),     IOS("IOS", "IOS"),     LINUX("LINUX", "LINUX"),     OTHER("OTHER", "OTHER");
func (r *TaobaoSmartstoreDeviceAddRequest) SetOsType(_osType string) error {
    r._osType = _osType
    r.Set("os_type", _osType)
    return nil
}

// OsType Getter
func (r TaobaoSmartstoreDeviceAddRequest) GetOsType() string {
    return r._osType
}
// DeviceType Setter
// 设备类型：     CAMERA("CAMERA", "客流摄像头"),     SHELF("SHELF", "云货架"),     MAKEUP_MIRROR("MAKEUP_MIRROR", "试妆镜"),     FITTING_MIRROR("FITTING_MIRROR", "试衣镜"),     VENDOR("VENDOR", "售货机"),     SAMPLE_MACHINE("SAMPLE_MACHINE","派样机"),     DOLL_MACHINE("DOLL_MACHINE", "娃娃机"),     INTERACTIVE_PHOTO("INTERACTIVE_PHOTO", "互动拍照"),     INTERACTIVE_GAME("INTERACTIVE_GAME", "互动游戏"),     USHER_SCREEN("USHER_SCREEN", "智慧迎宾屏"),     DRESSING("DRESSING", "闪电换装"),     MAGIC_MIRROR("MAGIC_MIRROR", "百搭魔镜"),     SHOES_FITTING_MIRROR("SHOES_FITTING_MIRROR", "试鞋镜"),     SKIN_DETECTION("SKIN_DETECTION", "肌肤测试仪"),     FOOT_DETECTION("FOOT_DETECTION", "测脚仪"),     RFID_SENSOR("RFID_SENSOR", "RFID"),touch_machine("touch_machine","导购一体屏")
func (r *TaobaoSmartstoreDeviceAddRequest) SetDeviceType(_deviceType string) error {
    r._deviceType = _deviceType
    r.Set("device_type", _deviceType)
    return nil
}

// DeviceType Getter
func (r TaobaoSmartstoreDeviceAddRequest) GetDeviceType() string {
    return r._deviceType
}
// OuterCode Setter
// 商家自定义设备编码
func (r *TaobaoSmartstoreDeviceAddRequest) SetOuterCode(_outerCode string) error {
    r._outerCode = _outerCode
    r.Set("outer_code", _outerCode)
    return nil
}

// OuterCode Getter
func (r TaobaoSmartstoreDeviceAddRequest) GetOuterCode() string {
    return r._outerCode
}
