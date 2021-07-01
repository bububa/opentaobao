package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机设备注册 API请求
alibaba.retail.device.vending.register

贩卖机注册
*/
type AlibabaRetailDeviceVendingRegisterAPIRequest struct {
    model.Params
    // 设备名称
    _deviceName   string
    // 设备地址
    _address   string
    // 对接过阿里atm传入
    _deviceSn   string
    // 业务编码，联系对接人申请
    _bizCode   string
    // 外部设备编号
    _deviceUuid   string
    // 设备类型
    _deviceModel   string
    // COMMUNITY：小区,SCHOOL：学校,OFFICE：写字楼,SHOPPING_MALL：商场,AIRPORT：机场,SUBWAY：地铁,HOSPITAL：医院,PLAYGROUNDS：游乐场所,FACTORY：工厂,VIEWPOINT：旅游景点,OTHERS：其他
    _scene   string
    // 场地名称，根据场地类型来，如：学校名称，商场名称。如果不传系统会根据address传入计算，为确保准确性请传入并确保address完整
    _siteName   string
    // 楼栋信息。如果不传系统会根据address传入计算，为确保准确性请传入并确保address完整
    _floor   string
    // 层。如果不传系统会根据address传入计算，为确保准确性请传入并确保address完整
    _layer   string
    // 室内地址描述。如果不传系统会根据address传入计算，为确保准确性请传入并确保address完整
    _location   string
}

// 初始化AlibabaRetailDeviceVendingRegisterAPIRequest对象
func NewAlibabaRetailDeviceVendingRegisterRequest() *AlibabaRetailDeviceVendingRegisterAPIRequest{
    return &AlibabaRetailDeviceVendingRegisterAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailDeviceVendingRegisterAPIRequest) GetApiMethodName() string {
    return "alibaba.retail.device.vending.register"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailDeviceVendingRegisterAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeviceName Setter
// 设备名称
func (r *AlibabaRetailDeviceVendingRegisterAPIRequest) SetDeviceName(_deviceName string) error {
    r._deviceName = _deviceName
    r.Set("device_name", _deviceName)
    return nil
}

// DeviceName Getter
func (r AlibabaRetailDeviceVendingRegisterAPIRequest) GetDeviceName() string {
    return r._deviceName
}
// Address Setter
// 设备地址
func (r *AlibabaRetailDeviceVendingRegisterAPIRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r AlibabaRetailDeviceVendingRegisterAPIRequest) GetAddress() string {
    return r._address
}
// DeviceSn Setter
// 对接过阿里atm传入
func (r *AlibabaRetailDeviceVendingRegisterAPIRequest) SetDeviceSn(_deviceSn string) error {
    r._deviceSn = _deviceSn
    r.Set("device_sn", _deviceSn)
    return nil
}

// DeviceSn Getter
func (r AlibabaRetailDeviceVendingRegisterAPIRequest) GetDeviceSn() string {
    return r._deviceSn
}
// BizCode Setter
// 业务编码，联系对接人申请
func (r *AlibabaRetailDeviceVendingRegisterAPIRequest) SetBizCode(_bizCode string) error {
    r._bizCode = _bizCode
    r.Set("biz_code", _bizCode)
    return nil
}

// BizCode Getter
func (r AlibabaRetailDeviceVendingRegisterAPIRequest) GetBizCode() string {
    return r._bizCode
}
// DeviceUuid Setter
// 外部设备编号
func (r *AlibabaRetailDeviceVendingRegisterAPIRequest) SetDeviceUuid(_deviceUuid string) error {
    r._deviceUuid = _deviceUuid
    r.Set("device_uuid", _deviceUuid)
    return nil
}

// DeviceUuid Getter
func (r AlibabaRetailDeviceVendingRegisterAPIRequest) GetDeviceUuid() string {
    return r._deviceUuid
}
// DeviceModel Setter
// 设备类型
func (r *AlibabaRetailDeviceVendingRegisterAPIRequest) SetDeviceModel(_deviceModel string) error {
    r._deviceModel = _deviceModel
    r.Set("device_model", _deviceModel)
    return nil
}

// DeviceModel Getter
func (r AlibabaRetailDeviceVendingRegisterAPIRequest) GetDeviceModel() string {
    return r._deviceModel
}
// Scene Setter
// COMMUNITY：小区,SCHOOL：学校,OFFICE：写字楼,SHOPPING_MALL：商场,AIRPORT：机场,SUBWAY：地铁,HOSPITAL：医院,PLAYGROUNDS：游乐场所,FACTORY：工厂,VIEWPOINT：旅游景点,OTHERS：其他
func (r *AlibabaRetailDeviceVendingRegisterAPIRequest) SetScene(_scene string) error {
    r._scene = _scene
    r.Set("scene", _scene)
    return nil
}

// Scene Getter
func (r AlibabaRetailDeviceVendingRegisterAPIRequest) GetScene() string {
    return r._scene
}
// SiteName Setter
// 场地名称，根据场地类型来，如：学校名称，商场名称。如果不传系统会根据address传入计算，为确保准确性请传入并确保address完整
func (r *AlibabaRetailDeviceVendingRegisterAPIRequest) SetSiteName(_siteName string) error {
    r._siteName = _siteName
    r.Set("site_name", _siteName)
    return nil
}

// SiteName Getter
func (r AlibabaRetailDeviceVendingRegisterAPIRequest) GetSiteName() string {
    return r._siteName
}
// Floor Setter
// 楼栋信息。如果不传系统会根据address传入计算，为确保准确性请传入并确保address完整
func (r *AlibabaRetailDeviceVendingRegisterAPIRequest) SetFloor(_floor string) error {
    r._floor = _floor
    r.Set("floor", _floor)
    return nil
}

// Floor Getter
func (r AlibabaRetailDeviceVendingRegisterAPIRequest) GetFloor() string {
    return r._floor
}
// Layer Setter
// 层。如果不传系统会根据address传入计算，为确保准确性请传入并确保address完整
func (r *AlibabaRetailDeviceVendingRegisterAPIRequest) SetLayer(_layer string) error {
    r._layer = _layer
    r.Set("layer", _layer)
    return nil
}

// Layer Getter
func (r AlibabaRetailDeviceVendingRegisterAPIRequest) GetLayer() string {
    return r._layer
}
// Location Setter
// 室内地址描述。如果不传系统会根据address传入计算，为确保准确性请传入并确保address完整
func (r *AlibabaRetailDeviceVendingRegisterAPIRequest) SetLocation(_location string) error {
    r._location = _location
    r.Set("location", _location)
    return nil
}

// Location Getter
func (r AlibabaRetailDeviceVendingRegisterAPIRequest) GetLocation() string {
    return r._location
}
