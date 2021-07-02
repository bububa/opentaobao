package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest 互联网医院批量导入接口 API请求
// alibaba.alihealth.medicalbase.hospital.sync
//
// 互联网医院isv批量通过接口批量导入
type AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest struct {
	model.Params
	// 是否需要用户授权
	_isAuth string
	// 主院区纬度
	_lat string
	// 主院区经度
	_lon string
	// 主院区地址
	_hosAddress string
	// 主院区的联系电话
	_telephone string
	// 院区名称
	_regionName string
	// 是否公立医院（Y／N）
	_isPublic string
	// 标签
	_serviceInfo string
	// 自定义科室
	_special string
	// 生活号或者服务窗url
	_serviceWindowUrl string
	// 医院简介url
	_descriptionUrl string
	// 是否支持医保（Y/N）
	_isInsurance string
	// 医院等级
	_grade string
	// 综合(general)、专科（special）
	_category string
	// 医院简称
	_shortName string
	// 医院pid
	_pid string
	// 机构编码
	_unifyCode string
	// 所在城市code
	_cityCode string
	// 营业执照上的医院全称
	_hosName string
	// 公司名称
	_companyName string
	// 支付宝BD的姓名
	_aliInterfaceMan string
	// 邮箱地址
	_email string
	// 联系人
	_technicalMan string
	// 联系手机
	_phone string
	// 单医院（main）／ 平台（platform）
	_hosType string
	// 服务项列表
	_functions string
	// isv库里面的hosCode
	_isvHosCode string
	// 投放阵地alipay aliyy uc quark
	_deliveryChannel string
}

// NewAlibabaAlihealthMedicalbaseHospitalSyncRequest 初始化AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest对象
func NewAlibabaAlihealthMedicalbaseHospitalSyncRequest() *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest {
	return &AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.medicalbase.hospital.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetIsAuth is IsAuth Setter
// 是否需要用户授权
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetIsAuth(_isAuth string) error {
	r._isAuth = _isAuth
	r.Set("is_auth", _isAuth)
	return nil
}

// GetIsAuth IsAuth Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetIsAuth() string {
	return r._isAuth
}

// SetLat is Lat Setter
// 主院区纬度
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetLat(_lat string) error {
	r._lat = _lat
	r.Set("lat", _lat)
	return nil
}

// GetLat Lat Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetLat() string {
	return r._lat
}

// SetLon is Lon Setter
// 主院区经度
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetLon(_lon string) error {
	r._lon = _lon
	r.Set("lon", _lon)
	return nil
}

// GetLon Lon Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetLon() string {
	return r._lon
}

// SetHosAddress is HosAddress Setter
// 主院区地址
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetHosAddress(_hosAddress string) error {
	r._hosAddress = _hosAddress
	r.Set("hos_address", _hosAddress)
	return nil
}

// GetHosAddress HosAddress Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetHosAddress() string {
	return r._hosAddress
}

// SetTelephone is Telephone Setter
// 主院区的联系电话
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetTelephone(_telephone string) error {
	r._telephone = _telephone
	r.Set("telephone", _telephone)
	return nil
}

// GetTelephone Telephone Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetTelephone() string {
	return r._telephone
}

// SetRegionName is RegionName Setter
// 院区名称
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetRegionName(_regionName string) error {
	r._regionName = _regionName
	r.Set("region_name", _regionName)
	return nil
}

// GetRegionName RegionName Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetRegionName() string {
	return r._regionName
}

// SetIsPublic is IsPublic Setter
// 是否公立医院（Y／N）
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetIsPublic(_isPublic string) error {
	r._isPublic = _isPublic
	r.Set("is_public", _isPublic)
	return nil
}

// GetIsPublic IsPublic Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetIsPublic() string {
	return r._isPublic
}

// SetServiceInfo is ServiceInfo Setter
// 标签
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetServiceInfo(_serviceInfo string) error {
	r._serviceInfo = _serviceInfo
	r.Set("service_info", _serviceInfo)
	return nil
}

// GetServiceInfo ServiceInfo Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetServiceInfo() string {
	return r._serviceInfo
}

// SetSpecial is Special Setter
// 自定义科室
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetSpecial(_special string) error {
	r._special = _special
	r.Set("special", _special)
	return nil
}

// GetSpecial Special Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetSpecial() string {
	return r._special
}

// SetServiceWindowUrl is ServiceWindowUrl Setter
// 生活号或者服务窗url
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetServiceWindowUrl(_serviceWindowUrl string) error {
	r._serviceWindowUrl = _serviceWindowUrl
	r.Set("service_window_url", _serviceWindowUrl)
	return nil
}

// GetServiceWindowUrl ServiceWindowUrl Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetServiceWindowUrl() string {
	return r._serviceWindowUrl
}

// SetDescriptionUrl is DescriptionUrl Setter
// 医院简介url
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetDescriptionUrl(_descriptionUrl string) error {
	r._descriptionUrl = _descriptionUrl
	r.Set("description_url", _descriptionUrl)
	return nil
}

// GetDescriptionUrl DescriptionUrl Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetDescriptionUrl() string {
	return r._descriptionUrl
}

// SetIsInsurance is IsInsurance Setter
// 是否支持医保（Y/N）
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetIsInsurance(_isInsurance string) error {
	r._isInsurance = _isInsurance
	r.Set("is_insurance", _isInsurance)
	return nil
}

// GetIsInsurance IsInsurance Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetIsInsurance() string {
	return r._isInsurance
}

// SetGrade is Grade Setter
// 医院等级
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetGrade(_grade string) error {
	r._grade = _grade
	r.Set("grade", _grade)
	return nil
}

// GetGrade Grade Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetGrade() string {
	return r._grade
}

// SetCategory is Category Setter
// 综合(general)、专科（special）
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetCategory(_category string) error {
	r._category = _category
	r.Set("category", _category)
	return nil
}

// GetCategory Category Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetCategory() string {
	return r._category
}

// SetShortName is ShortName Setter
// 医院简称
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetShortName(_shortName string) error {
	r._shortName = _shortName
	r.Set("short_name", _shortName)
	return nil
}

// GetShortName ShortName Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetShortName() string {
	return r._shortName
}

// SetPid is Pid Setter
// 医院pid
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetPid(_pid string) error {
	r._pid = _pid
	r.Set("pid", _pid)
	return nil
}

// GetPid Pid Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetPid() string {
	return r._pid
}

// SetUnifyCode is UnifyCode Setter
// 机构编码
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetUnifyCode(_unifyCode string) error {
	r._unifyCode = _unifyCode
	r.Set("unify_code", _unifyCode)
	return nil
}

// GetUnifyCode UnifyCode Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetUnifyCode() string {
	return r._unifyCode
}

// SetCityCode is CityCode Setter
// 所在城市code
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetCityCode() string {
	return r._cityCode
}

// SetHosName is HosName Setter
// 营业执照上的医院全称
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetHosName(_hosName string) error {
	r._hosName = _hosName
	r.Set("hos_name", _hosName)
	return nil
}

// GetHosName HosName Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetHosName() string {
	return r._hosName
}

// SetCompanyName is CompanyName Setter
// 公司名称
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetCompanyName(_companyName string) error {
	r._companyName = _companyName
	r.Set("company_name", _companyName)
	return nil
}

// GetCompanyName CompanyName Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetCompanyName() string {
	return r._companyName
}

// SetAliInterfaceMan is AliInterfaceMan Setter
// 支付宝BD的姓名
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetAliInterfaceMan(_aliInterfaceMan string) error {
	r._aliInterfaceMan = _aliInterfaceMan
	r.Set("ali_interface_man", _aliInterfaceMan)
	return nil
}

// GetAliInterfaceMan AliInterfaceMan Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetAliInterfaceMan() string {
	return r._aliInterfaceMan
}

// SetEmail is Email Setter
// 邮箱地址
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetEmail(_email string) error {
	r._email = _email
	r.Set("email", _email)
	return nil
}

// GetEmail Email Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetEmail() string {
	return r._email
}

// SetTechnicalMan is TechnicalMan Setter
// 联系人
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetTechnicalMan(_technicalMan string) error {
	r._technicalMan = _technicalMan
	r.Set("technical_man", _technicalMan)
	return nil
}

// GetTechnicalMan TechnicalMan Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetTechnicalMan() string {
	return r._technicalMan
}

// SetPhone is Phone Setter
// 联系手机
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetPhone() string {
	return r._phone
}

// SetHosType is HosType Setter
// 单医院（main）／ 平台（platform）
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetHosType(_hosType string) error {
	r._hosType = _hosType
	r.Set("hos_type", _hosType)
	return nil
}

// GetHosType HosType Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetHosType() string {
	return r._hosType
}

// SetFunctions is Functions Setter
// 服务项列表
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetFunctions(_functions string) error {
	r._functions = _functions
	r.Set("functions", _functions)
	return nil
}

// GetFunctions Functions Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetFunctions() string {
	return r._functions
}

// SetIsvHosCode is IsvHosCode Setter
// isv库里面的hosCode
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetIsvHosCode(_isvHosCode string) error {
	r._isvHosCode = _isvHosCode
	r.Set("isv_hos_code", _isvHosCode)
	return nil
}

// GetIsvHosCode IsvHosCode Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetIsvHosCode() string {
	return r._isvHosCode
}

// SetDeliveryChannel is DeliveryChannel Setter
// 投放阵地alipay aliyy uc quark
func (r *AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) SetDeliveryChannel(_deliveryChannel string) error {
	r._deliveryChannel = _deliveryChannel
	r.Set("delivery_channel", _deliveryChannel)
	return nil
}

// GetDeliveryChannel DeliveryChannel Getter
func (r AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest) GetDeliveryChannel() string {
	return r._deliveryChannel
}
