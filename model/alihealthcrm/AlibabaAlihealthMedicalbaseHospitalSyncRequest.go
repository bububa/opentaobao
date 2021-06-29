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
    _isAuth   string
    // 主院区纬度
    _lat   string
    // 主院区经度
    _lon   string
    // 主院区地址
    _hosAddress   string
    // 主院区的联系电话
    _telephone   string
    // 院区名称
    _regionName   string
    // 是否公立医院（Y／N）
    _isPublic   string
    // 标签
    _serviceInfo   string
    // 自定义科室
    _special   string
    // 生活号或者服务窗url
    _serviceWindowUrl   string
    // 医院简介url
    _descriptionUrl   string
    // 是否支持医保（Y/N）
    _isInsurance   string
    // 医院等级
    _grade   string
    // 综合(general)、专科（special）
    _category   string
    // 医院简称
    _shortName   string
    // 医院pid
    _pid   string
    // 机构编码
    _unifyCode   string
    // 所在城市code
    _cityCode   string
    // 营业执照上的医院全称
    _hosName   string
    // 公司名称
    _companyName   string
    // 支付宝BD的姓名
    _aliInterfaceMan   string
    // 邮箱地址
    _email   string
    // 联系人
    _technicalMan   string
    // 联系手机
    _phone   string
    // 单医院（main）／ 平台（platform）
    _hosType   string
    // 服务项列表
    _functions   string
    // isv库里面的hosCode
    _isvHosCode   string
    // 投放阵地alipay aliyy uc quark
    _deliveryChannel   string
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
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetIsAuth(_isAuth string) error {
    r._isAuth = _isAuth
    r.Set("is_auth", _isAuth)
    return nil
}

// IsAuth Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetIsAuth() string {
    return r._isAuth
}
// Lat Setter
// 主院区纬度
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetLat(_lat string) error {
    r._lat = _lat
    r.Set("lat", _lat)
    return nil
}

// Lat Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetLat() string {
    return r._lat
}
// Lon Setter
// 主院区经度
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetLon(_lon string) error {
    r._lon = _lon
    r.Set("lon", _lon)
    return nil
}

// Lon Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetLon() string {
    return r._lon
}
// HosAddress Setter
// 主院区地址
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetHosAddress(_hosAddress string) error {
    r._hosAddress = _hosAddress
    r.Set("hos_address", _hosAddress)
    return nil
}

// HosAddress Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetHosAddress() string {
    return r._hosAddress
}
// Telephone Setter
// 主院区的联系电话
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetTelephone(_telephone string) error {
    r._telephone = _telephone
    r.Set("telephone", _telephone)
    return nil
}

// Telephone Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetTelephone() string {
    return r._telephone
}
// RegionName Setter
// 院区名称
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetRegionName(_regionName string) error {
    r._regionName = _regionName
    r.Set("region_name", _regionName)
    return nil
}

// RegionName Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetRegionName() string {
    return r._regionName
}
// IsPublic Setter
// 是否公立医院（Y／N）
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetIsPublic(_isPublic string) error {
    r._isPublic = _isPublic
    r.Set("is_public", _isPublic)
    return nil
}

// IsPublic Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetIsPublic() string {
    return r._isPublic
}
// ServiceInfo Setter
// 标签
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetServiceInfo(_serviceInfo string) error {
    r._serviceInfo = _serviceInfo
    r.Set("service_info", _serviceInfo)
    return nil
}

// ServiceInfo Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetServiceInfo() string {
    return r._serviceInfo
}
// Special Setter
// 自定义科室
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetSpecial(_special string) error {
    r._special = _special
    r.Set("special", _special)
    return nil
}

// Special Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetSpecial() string {
    return r._special
}
// ServiceWindowUrl Setter
// 生活号或者服务窗url
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetServiceWindowUrl(_serviceWindowUrl string) error {
    r._serviceWindowUrl = _serviceWindowUrl
    r.Set("service_window_url", _serviceWindowUrl)
    return nil
}

// ServiceWindowUrl Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetServiceWindowUrl() string {
    return r._serviceWindowUrl
}
// DescriptionUrl Setter
// 医院简介url
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetDescriptionUrl(_descriptionUrl string) error {
    r._descriptionUrl = _descriptionUrl
    r.Set("description_url", _descriptionUrl)
    return nil
}

// DescriptionUrl Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetDescriptionUrl() string {
    return r._descriptionUrl
}
// IsInsurance Setter
// 是否支持医保（Y/N）
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetIsInsurance(_isInsurance string) error {
    r._isInsurance = _isInsurance
    r.Set("is_insurance", _isInsurance)
    return nil
}

// IsInsurance Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetIsInsurance() string {
    return r._isInsurance
}
// Grade Setter
// 医院等级
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetGrade(_grade string) error {
    r._grade = _grade
    r.Set("grade", _grade)
    return nil
}

// Grade Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetGrade() string {
    return r._grade
}
// Category Setter
// 综合(general)、专科（special）
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetCategory(_category string) error {
    r._category = _category
    r.Set("category", _category)
    return nil
}

// Category Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetCategory() string {
    return r._category
}
// ShortName Setter
// 医院简称
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetShortName(_shortName string) error {
    r._shortName = _shortName
    r.Set("short_name", _shortName)
    return nil
}

// ShortName Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetShortName() string {
    return r._shortName
}
// Pid Setter
// 医院pid
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetPid(_pid string) error {
    r._pid = _pid
    r.Set("pid", _pid)
    return nil
}

// Pid Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetPid() string {
    return r._pid
}
// UnifyCode Setter
// 机构编码
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetUnifyCode(_unifyCode string) error {
    r._unifyCode = _unifyCode
    r.Set("unify_code", _unifyCode)
    return nil
}

// UnifyCode Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetUnifyCode() string {
    return r._unifyCode
}
// CityCode Setter
// 所在城市code
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetCityCode(_cityCode string) error {
    r._cityCode = _cityCode
    r.Set("city_code", _cityCode)
    return nil
}

// CityCode Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetCityCode() string {
    return r._cityCode
}
// HosName Setter
// 营业执照上的医院全称
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetHosName(_hosName string) error {
    r._hosName = _hosName
    r.Set("hos_name", _hosName)
    return nil
}

// HosName Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetHosName() string {
    return r._hosName
}
// CompanyName Setter
// 公司名称
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetCompanyName(_companyName string) error {
    r._companyName = _companyName
    r.Set("company_name", _companyName)
    return nil
}

// CompanyName Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetCompanyName() string {
    return r._companyName
}
// AliInterfaceMan Setter
// 支付宝BD的姓名
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetAliInterfaceMan(_aliInterfaceMan string) error {
    r._aliInterfaceMan = _aliInterfaceMan
    r.Set("ali_interface_man", _aliInterfaceMan)
    return nil
}

// AliInterfaceMan Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetAliInterfaceMan() string {
    return r._aliInterfaceMan
}
// Email Setter
// 邮箱地址
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetEmail(_email string) error {
    r._email = _email
    r.Set("email", _email)
    return nil
}

// Email Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetEmail() string {
    return r._email
}
// TechnicalMan Setter
// 联系人
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetTechnicalMan(_technicalMan string) error {
    r._technicalMan = _technicalMan
    r.Set("technical_man", _technicalMan)
    return nil
}

// TechnicalMan Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetTechnicalMan() string {
    return r._technicalMan
}
// Phone Setter
// 联系手机
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetPhone(_phone string) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetPhone() string {
    return r._phone
}
// HosType Setter
// 单医院（main）／ 平台（platform）
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetHosType(_hosType string) error {
    r._hosType = _hosType
    r.Set("hos_type", _hosType)
    return nil
}

// HosType Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetHosType() string {
    return r._hosType
}
// Functions Setter
// 服务项列表
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetFunctions(_functions string) error {
    r._functions = _functions
    r.Set("functions", _functions)
    return nil
}

// Functions Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetFunctions() string {
    return r._functions
}
// IsvHosCode Setter
// isv库里面的hosCode
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetIsvHosCode(_isvHosCode string) error {
    r._isvHosCode = _isvHosCode
    r.Set("isv_hos_code", _isvHosCode)
    return nil
}

// IsvHosCode Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetIsvHosCode() string {
    return r._isvHosCode
}
// DeliveryChannel Setter
// 投放阵地alipay aliyy uc quark
func (r *AlibabaAlihealthMedicalbaseHospitalSyncRequest) SetDeliveryChannel(_deliveryChannel string) error {
    r._deliveryChannel = _deliveryChannel
    r.Set("delivery_channel", _deliveryChannel)
    return nil
}

// DeliveryChannel Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncRequest) GetDeliveryChannel() string {
    return r._deliveryChannel
}
