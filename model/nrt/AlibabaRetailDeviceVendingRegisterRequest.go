package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机设备注册 APIRequest
alibaba.retail.device.vending.register

贩卖机注册
*/
type AlibabaRetailDeviceVendingRegisterRequest struct {
    model.Params

    // 设备名称
    deviceName   string 

    // 设备地址
    address   string 

    // 对接过阿里atm传入
    deviceSn   string 

    // 业务编码，联系对接人申请
    bizCode   string 

    // 外部设备编号
    deviceUuid   string 

    // 设备类型
    deviceModel   string 

    // COMMUNITY：小区,SCHOOL：学校,OFFICE：写字楼,SHOPPING_MALL：商场,AIRPORT：机场,SUBWAY：地铁,HOSPITAL：医院,PLAYGROUNDS：游乐场所,FACTORY：工厂,VIEWPOINT：旅游景点,OTHERS：其他
    scene   string 

    // 场地名称，根据场地类型来，如：学校名称，商场名称。如果不传系统会根据address传入计算，为确保准确性请传入并确保address完整
    siteName   string 

    // 楼栋信息。如果不传系统会根据address传入计算，为确保准确性请传入并确保address完整
    floor   string 

    // 层。如果不传系统会根据address传入计算，为确保准确性请传入并确保address完整
    layer   string 

    // 室内地址描述。如果不传系统会根据address传入计算，为确保准确性请传入并确保address完整
    location   string 

}

func NewAlibabaRetailDeviceVendingRegisterRequest() *AlibabaRetailDeviceVendingRegisterRequest{
    return &AlibabaRetailDeviceVendingRegisterRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailDeviceVendingRegisterRequest) GetApiMethodName() string {
    return "alibaba.retail.device.vending.register"
}

func (r AlibabaRetailDeviceVendingRegisterRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailDeviceVendingRegisterRequest) SetDeviceName(deviceName string) error {
    r.deviceName = deviceName
    r.Set("device_name", deviceName)
    return nil
}

func (r AlibabaRetailDeviceVendingRegisterRequest) GetDeviceName() string {
    return r.deviceName
}

func (r *AlibabaRetailDeviceVendingRegisterRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

func (r AlibabaRetailDeviceVendingRegisterRequest) GetAddress() string {
    return r.address
}

func (r *AlibabaRetailDeviceVendingRegisterRequest) SetDeviceSn(deviceSn string) error {
    r.deviceSn = deviceSn
    r.Set("device_sn", deviceSn)
    return nil
}

func (r AlibabaRetailDeviceVendingRegisterRequest) GetDeviceSn() string {
    return r.deviceSn
}

func (r *AlibabaRetailDeviceVendingRegisterRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

func (r AlibabaRetailDeviceVendingRegisterRequest) GetBizCode() string {
    return r.bizCode
}

func (r *AlibabaRetailDeviceVendingRegisterRequest) SetDeviceUuid(deviceUuid string) error {
    r.deviceUuid = deviceUuid
    r.Set("device_uuid", deviceUuid)
    return nil
}

func (r AlibabaRetailDeviceVendingRegisterRequest) GetDeviceUuid() string {
    return r.deviceUuid
}

func (r *AlibabaRetailDeviceVendingRegisterRequest) SetDeviceModel(deviceModel string) error {
    r.deviceModel = deviceModel
    r.Set("device_model", deviceModel)
    return nil
}

func (r AlibabaRetailDeviceVendingRegisterRequest) GetDeviceModel() string {
    return r.deviceModel
}

func (r *AlibabaRetailDeviceVendingRegisterRequest) SetScene(scene string) error {
    r.scene = scene
    r.Set("scene", scene)
    return nil
}

func (r AlibabaRetailDeviceVendingRegisterRequest) GetScene() string {
    return r.scene
}

func (r *AlibabaRetailDeviceVendingRegisterRequest) SetSiteName(siteName string) error {
    r.siteName = siteName
    r.Set("site_name", siteName)
    return nil
}

func (r AlibabaRetailDeviceVendingRegisterRequest) GetSiteName() string {
    return r.siteName
}

func (r *AlibabaRetailDeviceVendingRegisterRequest) SetFloor(floor string) error {
    r.floor = floor
    r.Set("floor", floor)
    return nil
}

func (r AlibabaRetailDeviceVendingRegisterRequest) GetFloor() string {
    return r.floor
}

func (r *AlibabaRetailDeviceVendingRegisterRequest) SetLayer(layer string) error {
    r.layer = layer
    r.Set("layer", layer)
    return nil
}

func (r AlibabaRetailDeviceVendingRegisterRequest) GetLayer() string {
    return r.layer
}

func (r *AlibabaRetailDeviceVendingRegisterRequest) SetLocation(location string) error {
    r.location = location
    r.Set("location", location)
    return nil
}

func (r AlibabaRetailDeviceVendingRegisterRequest) GetLocation() string {
    return r.location
}

