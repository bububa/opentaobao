package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreserveconfirmAPIRequest 体检机构对接_体检套餐预定确认 API请求
// alibaba.alihealth.examination.reserve.confirm
//
// 向体检机构确认用户购买的体检套餐信息
type AlibabaalihealthexaminationreserveconfirmAPIRequest struct {
	model.Params
	// 加项列表
	_addItems []AddItem
	// 加项包列表
	_addPacks []AddPack
	// 商户唯一码
	_merchantCode string
	// 体检人姓名
	_name string
	// 阿里健康预约唯一标识
	_reserveNumber string
	// 性别(0-男;1-女;)
	_gender string
	// 出生日期
	_birthday string
	// 预约时间
	_reserveDate string
	// 体检套餐编码
	_packageCode string
	// 婚否(0-未婚; 1-已婚)
	_married string
	// 店铺ID
	_storeId string
	// 电话号码
	_phone string
	// 证件类型(0-身份证; 1-护照; 2-军官证)
	_certType string
	// 证件号
	_certNumber string
	// 所属公司
	_company string
	// 所属部门
	_department string
	// 报告邮寄地址
	_address string
	// 0没报告1有报告
	_havaReport string
	// 员工号
	_employeeNumber string
	// 服务类型，ONSITE_SERVICE（到店检测）、DOOR_TO_DOOR_SERVICE（上门检测）
	_serviceType string
	// 预约时间段开始
	_reserveTimeStart string
	// 预约时间段截止
	_reserveTimeEnd string
	// 上门服务的上门地址
	_serviceAddress *AddAddress
}

// NewAlibabaalihealthexaminationreserveconfirmRequest 初始化AlibabaalihealthexaminationreserveconfirmAPIRequest对象
func NewAlibabaalihealthexaminationreserveconfirmRequest() *AlibabaalihealthexaminationreserveconfirmAPIRequest {
	return &AlibabaalihealthexaminationreserveconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.reserve.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAddItems is AddItems Setter
// 加项列表
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetAddItems(_addItems []AddItem) error {
	r._addItems = _addItems
	r.Set("add_items", _addItems)
	return nil
}

// GetAddItems AddItems Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetAddItems() []AddItem {
	return r._addItems
}

// SetAddPacks is AddPacks Setter
// 加项包列表
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetAddPacks(_addPacks []AddPack) error {
	r._addPacks = _addPacks
	r.Set("add_packs", _addPacks)
	return nil
}

// GetAddPacks AddPacks Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetAddPacks() []AddPack {
	return r._addPacks
}

// SetMerchantCode is MerchantCode Setter
// 商户唯一码
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetMerchantCode(_merchantCode string) error {
	r._merchantCode = _merchantCode
	r.Set("merchant_code", _merchantCode)
	return nil
}

// GetMerchantCode MerchantCode Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetMerchantCode() string {
	return r._merchantCode
}

// SetName is Name Setter
// 体检人姓名
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetName() string {
	return r._name
}

// SetReserveNumber is ReserveNumber Setter
// 阿里健康预约唯一标识
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetReserveNumber(_reserveNumber string) error {
	r._reserveNumber = _reserveNumber
	r.Set("reserve_number", _reserveNumber)
	return nil
}

// GetReserveNumber ReserveNumber Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetReserveNumber() string {
	return r._reserveNumber
}

// SetGender is Gender Setter
// 性别(0-男;1-女;)
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetGender(_gender string) error {
	r._gender = _gender
	r.Set("gender", _gender)
	return nil
}

// GetGender Gender Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetGender() string {
	return r._gender
}

// SetBirthday is Birthday Setter
// 出生日期
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetBirthday(_birthday string) error {
	r._birthday = _birthday
	r.Set("birthday", _birthday)
	return nil
}

// GetBirthday Birthday Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetBirthday() string {
	return r._birthday
}

// SetReserveDate is ReserveDate Setter
// 预约时间
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetReserveDate(_reserveDate string) error {
	r._reserveDate = _reserveDate
	r.Set("reserve_date", _reserveDate)
	return nil
}

// GetReserveDate ReserveDate Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetReserveDate() string {
	return r._reserveDate
}

// SetPackageCode is PackageCode Setter
// 体检套餐编码
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetPackageCode(_packageCode string) error {
	r._packageCode = _packageCode
	r.Set("package_code", _packageCode)
	return nil
}

// GetPackageCode PackageCode Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetPackageCode() string {
	return r._packageCode
}

// SetMarried is Married Setter
// 婚否(0-未婚; 1-已婚)
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetMarried(_married string) error {
	r._married = _married
	r.Set("married", _married)
	return nil
}

// GetMarried Married Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetMarried() string {
	return r._married
}

// SetStoreId is StoreId Setter
// 店铺ID
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetPhone is Phone Setter
// 电话号码
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetPhone() string {
	return r._phone
}

// SetCertType is CertType Setter
// 证件类型(0-身份证; 1-护照; 2-军官证)
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetCertType(_certType string) error {
	r._certType = _certType
	r.Set("cert_type", _certType)
	return nil
}

// GetCertType CertType Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetCertType() string {
	return r._certType
}

// SetCertNumber is CertNumber Setter
// 证件号
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetCertNumber(_certNumber string) error {
	r._certNumber = _certNumber
	r.Set("cert_number", _certNumber)
	return nil
}

// GetCertNumber CertNumber Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetCertNumber() string {
	return r._certNumber
}

// SetCompany is Company Setter
// 所属公司
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetCompany(_company string) error {
	r._company = _company
	r.Set("company", _company)
	return nil
}

// GetCompany Company Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetCompany() string {
	return r._company
}

// SetDepartment is Department Setter
// 所属部门
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetDepartment(_department string) error {
	r._department = _department
	r.Set("department", _department)
	return nil
}

// GetDepartment Department Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetDepartment() string {
	return r._department
}

// SetAddress is Address Setter
// 报告邮寄地址
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetAddress() string {
	return r._address
}

// SetHavaReport is HavaReport Setter
// 0没报告1有报告
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetHavaReport(_havaReport string) error {
	r._havaReport = _havaReport
	r.Set("hava_report", _havaReport)
	return nil
}

// GetHavaReport HavaReport Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetHavaReport() string {
	return r._havaReport
}

// SetEmployeeNumber is EmployeeNumber Setter
// 员工号
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetEmployeeNumber(_employeeNumber string) error {
	r._employeeNumber = _employeeNumber
	r.Set("employee_number", _employeeNumber)
	return nil
}

// GetEmployeeNumber EmployeeNumber Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetEmployeeNumber() string {
	return r._employeeNumber
}

// SetServiceType is ServiceType Setter
// 服务类型，ONSITE_SERVICE（到店检测）、DOOR_TO_DOOR_SERVICE（上门检测）
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetServiceType(_serviceType string) error {
	r._serviceType = _serviceType
	r.Set("service_type", _serviceType)
	return nil
}

// GetServiceType ServiceType Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetServiceType() string {
	return r._serviceType
}

// SetReserveTimeStart is ReserveTimeStart Setter
// 预约时间段开始
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetReserveTimeStart(_reserveTimeStart string) error {
	r._reserveTimeStart = _reserveTimeStart
	r.Set("reserve_time_start", _reserveTimeStart)
	return nil
}

// GetReserveTimeStart ReserveTimeStart Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetReserveTimeStart() string {
	return r._reserveTimeStart
}

// SetReserveTimeEnd is ReserveTimeEnd Setter
// 预约时间段截止
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetReserveTimeEnd(_reserveTimeEnd string) error {
	r._reserveTimeEnd = _reserveTimeEnd
	r.Set("reserve_time_end", _reserveTimeEnd)
	return nil
}

// GetReserveTimeEnd ReserveTimeEnd Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetReserveTimeEnd() string {
	return r._reserveTimeEnd
}

// SetServiceAddress is ServiceAddress Setter
// 上门服务的上门地址
func (r *AlibabaalihealthexaminationreserveconfirmAPIRequest) SetServiceAddress(_serviceAddress *AddAddress) error {
	r._serviceAddress = _serviceAddress
	r.Set("service_address", _serviceAddress)
	return nil
}

// GetServiceAddress ServiceAddress Getter
func (r AlibabaalihealthexaminationreserveconfirmAPIRequest) GetServiceAddress() *AddAddress {
	return r._serviceAddress
}
