package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
互联网医院批量导入接口 API请求
alibaba.alihealth.medicalbase.hospital.sync

互联网医院isv批量通过接口批量导入
*/
type AlibabaAlihealthMedicalbaseHospitalSyncRequest struct {
    model.Params
    // 是否需要用户授权
    isAuth   string
    // 主院区纬度
    lat   string
    // 主院区经度
    lon   string
    // 主院区地址
    hosAddress   string
    // 主院区的联系电话
    telephone   string
    // 院区名称
    regionName   string
    // 是否公立医院（Y／N）
    isPublic   string
    // 标签
    serviceInfo   string
    // 自定义科室
    special   string
    // 生活号或者服务窗url
    serviceWindowUrl   string
    // 医院简介url
    descriptionUrl   string
    // 是否支持医保（Y/N）
    isInsurance   string
    // 医院等级
    grade   string
    // 综合(general)、专科（special）
    category   string
    // 医院简称
    shortName   string
    // 医院pid
    pid   string
    // 机构编码
    unifyCode   string
    // 所在城市code
    cityCode   string
    // 营业执照上的医院全称
    hosName   string
    // 公司名称
    companyName   string
    // 支付宝BD的姓名
    aliInterfaceMan   string
    // 邮箱地址
    email   string
    // 联系人
    technicalMan   string
    // 联系手机
    phone   string
    // 单医院（main）／ 平台（platform）
    hosType   string
    // 服务项列表
    functions   string
    // isv库里面的hosCode
    isvHosCode   string
    // 投放阵地alipay aliyy uc quark
    deliveryChannel   string
}

// 初始化AlibabaAlihealthMedicalbaseHospitalSyncRequest对象
func NewAlibabaAlihealthMedicalbaseHospitalSyncRequest() *AlibabaAlihealthMedicalbaseHospitalSyncRequest{
    return &AlibabaAlihealthMedicalbaseHospitalSyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetApiMethodName() string {
    return "alibaba.alihealth.medicalbase.hospital.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IsAuth Setter
// 是否需要用户授权
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetIsAuth(isAuth string) error {
    r.isAuth = isAuth
    r.Set("is_auth", isAuth)
    return nil
}

// IsAuth Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetIsAuth() string {
    return r.isAuth
}
// Lat Setter
// 主院区纬度
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetLat(lat string) error {
    r.lat = lat
    r.Set("lat", lat)
    return nil
}

// Lat Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetLat() string {
    return r.lat
}
// Lon Setter
// 主院区经度
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetLon(lon string) error {
    r.lon = lon
    r.Set("lon", lon)
    return nil
}

// Lon Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetLon() string {
    return r.lon
}
// HosAddress Setter
// 主院区地址
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetHosAddress(hosAddress string) error {
    r.hosAddress = hosAddress
    r.Set("hos_address", hosAddress)
    return nil
}

// HosAddress Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetHosAddress() string {
    return r.hosAddress
}
// Telephone Setter
// 主院区的联系电话
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetTelephone(telephone string) error {
    r.telephone = telephone
    r.Set("telephone", telephone)
    return nil
}

// Telephone Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetTelephone() string {
    return r.telephone
}
// RegionName Setter
// 院区名称
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetRegionName(regionName string) error {
    r.regionName = regionName
    r.Set("region_name", regionName)
    return nil
}

// RegionName Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetRegionName() string {
    return r.regionName
}
// IsPublic Setter
// 是否公立医院（Y／N）
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetIsPublic(isPublic string) error {
    r.isPublic = isPublic
    r.Set("is_public", isPublic)
    return nil
}

// IsPublic Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetIsPublic() string {
    return r.isPublic
}
// ServiceInfo Setter
// 标签
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetServiceInfo(serviceInfo string) error {
    r.serviceInfo = serviceInfo
    r.Set("service_info", serviceInfo)
    return nil
}

// ServiceInfo Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetServiceInfo() string {
    return r.serviceInfo
}
// Special Setter
// 自定义科室
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetSpecial(special string) error {
    r.special = special
    r.Set("special", special)
    return nil
}

// Special Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetSpecial() string {
    return r.special
}
// ServiceWindowUrl Setter
// 生活号或者服务窗url
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetServiceWindowUrl(serviceWindowUrl string) error {
    r.serviceWindowUrl = serviceWindowUrl
    r.Set("service_window_url", serviceWindowUrl)
    return nil
}

// ServiceWindowUrl Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetServiceWindowUrl() string {
    return r.serviceWindowUrl
}
// DescriptionUrl Setter
// 医院简介url
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetDescriptionUrl(descriptionUrl string) error {
    r.descriptionUrl = descriptionUrl
    r.Set("description_url", descriptionUrl)
    return nil
}

// DescriptionUrl Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetDescriptionUrl() string {
    return r.descriptionUrl
}
// IsInsurance Setter
// 是否支持医保（Y/N）
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetIsInsurance(isInsurance string) error {
    r.isInsurance = isInsurance
    r.Set("is_insurance", isInsurance)
    return nil
}

// IsInsurance Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetIsInsurance() string {
    return r.isInsurance
}
// Grade Setter
// 医院等级
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetGrade(grade string) error {
    r.grade = grade
    r.Set("grade", grade)
    return nil
}

// Grade Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetGrade() string {
    return r.grade
}
// Category Setter
// 综合(general)、专科（special）
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetCategory(category string) error {
    r.category = category
    r.Set("category", category)
    return nil
}

// Category Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetCategory() string {
    return r.category
}
// ShortName Setter
// 医院简称
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetShortName(shortName string) error {
    r.shortName = shortName
    r.Set("short_name", shortName)
    return nil
}

// ShortName Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetShortName() string {
    return r.shortName
}
// Pid Setter
// 医院pid
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetPid(pid string) error {
    r.pid = pid
    r.Set("pid", pid)
    return nil
}

// Pid Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetPid() string {
    return r.pid
}
// UnifyCode Setter
// 机构编码
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetUnifyCode(unifyCode string) error {
    r.unifyCode = unifyCode
    r.Set("unify_code", unifyCode)
    return nil
}

// UnifyCode Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetUnifyCode() string {
    return r.unifyCode
}
// CityCode Setter
// 所在城市code
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

// CityCode Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetCityCode() string {
    return r.cityCode
}
// HosName Setter
// 营业执照上的医院全称
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetHosName(hosName string) error {
    r.hosName = hosName
    r.Set("hos_name", hosName)
    return nil
}

// HosName Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetHosName() string {
    return r.hosName
}
// CompanyName Setter
// 公司名称
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetCompanyName(companyName string) error {
    r.companyName = companyName
    r.Set("company_name", companyName)
    return nil
}

// CompanyName Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetCompanyName() string {
    return r.companyName
}
// AliInterfaceMan Setter
// 支付宝BD的姓名
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetAliInterfaceMan(aliInterfaceMan string) error {
    r.aliInterfaceMan = aliInterfaceMan
    r.Set("ali_interface_man", aliInterfaceMan)
    return nil
}

// AliInterfaceMan Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetAliInterfaceMan() string {
    return r.aliInterfaceMan
}
// Email Setter
// 邮箱地址
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetEmail(email string) error {
    r.email = email
    r.Set("email", email)
    return nil
}

// Email Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetEmail() string {
    return r.email
}
// TechnicalMan Setter
// 联系人
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetTechnicalMan(technicalMan string) error {
    r.technicalMan = technicalMan
    r.Set("technical_man", technicalMan)
    return nil
}

// TechnicalMan Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetTechnicalMan() string {
    return r.technicalMan
}
// Phone Setter
// 联系手机
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

// Phone Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetPhone() string {
    return r.phone
}
// HosType Setter
// 单医院（main）／ 平台（platform）
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetHosType(hosType string) error {
    r.hosType = hosType
    r.Set("hos_type", hosType)
    return nil
}

// HosType Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetHosType() string {
    return r.hosType
}
// Functions Setter
// 服务项列表
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetFunctions(functions string) error {
    r.functions = functions
    r.Set("functions", functions)
    return nil
}

// Functions Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetFunctions() string {
    return r.functions
}
// IsvHosCode Setter
// isv库里面的hosCode
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetIsvHosCode(isvHosCode string) error {
    r.isvHosCode = isvHosCode
    r.Set("isv_hos_code", isvHosCode)
    return nil
}

// IsvHosCode Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetIsvHosCode() string {
    return r.isvHosCode
}
// DeliveryChannel Setter
// 投放阵地alipay aliyy uc quark
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetDeliveryChannel(deliveryChannel string) error {
    r.deliveryChannel = deliveryChannel
    r.Set("delivery_channel", deliveryChannel)
    return nil
}

// DeliveryChannel Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetDeliveryChannel() string {
    return r.deliveryChannel
}
