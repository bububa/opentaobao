package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReserveConfirmAPIRequest 体检机构对接_体检套餐预定确认 API请求
// alibaba.alihealth.examination.reserve.confirm
//
// 向体检机构确认用户购买的体检套餐信息
type AlibabaAlihealthExaminationReserveConfirmAPIRequest struct {
	model.Params
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
	// 加项列表
	_addItems []AddItem
	// 加项包列表
	_addPacks []AddPack
	// 0没报告1有报告
	_havaReport string
	// 员工号
	_employeeNumber string
	// 服务类型，ONSITE_SERVICE（到店检测）、DOOR_TO_DOOR_SERVICE（上门检测）
	_serviceType string
	// 上门服务的上门地址
	_serviceAddress *AddAddress
	// 预约时间段开始
	_reserveTimeStart string
	// 预约时间段截止
	_reserveTimeEnd string
}

// NewAlibabaAlihealthExaminationReserveConfirmRequest 初始化AlibabaAlihealthExaminationReserveConfirmAPIRequest对象
func NewAlibabaAlihealthExaminationReserveConfirmRequest() *AlibabaAlihealthExaminationReserveConfirmAPIRequest {
	return &AlibabaAlihealthExaminationReserveConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.examination.reserve.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MerchantCode Setter
// 商户唯一码
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetMerchantCode(_merchantCode string) error {
	r._merchantCode = _merchantCode
	r.Set("merchant_code", _merchantCode)
	return nil
}

// Get MerchantCode Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetMerchantCode() string {
	return r._merchantCode
}

// Set is Name Setter
// 体检人姓名
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetName() string {
	return r._name
}

// Set is ReserveNumber Setter
// 阿里健康预约唯一标识
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetReserveNumber(_reserveNumber string) error {
	r._reserveNumber = _reserveNumber
	r.Set("reserve_number", _reserveNumber)
	return nil
}

// Get ReserveNumber Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetReserveNumber() string {
	return r._reserveNumber
}

// Set is Gender Setter
// 性别(0-男;1-女;)
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetGender(_gender string) error {
	r._gender = _gender
	r.Set("gender", _gender)
	return nil
}

// Get Gender Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetGender() string {
	return r._gender
}

// Set is Birthday Setter
// 出生日期
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetBirthday(_birthday string) error {
	r._birthday = _birthday
	r.Set("birthday", _birthday)
	return nil
}

// Get Birthday Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetBirthday() string {
	return r._birthday
}

// Set is ReserveDate Setter
// 预约时间
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetReserveDate(_reserveDate string) error {
	r._reserveDate = _reserveDate
	r.Set("reserve_date", _reserveDate)
	return nil
}

// Get ReserveDate Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetReserveDate() string {
	return r._reserveDate
}

// Set is PackageCode Setter
// 体检套餐编码
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetPackageCode(_packageCode string) error {
	r._packageCode = _packageCode
	r.Set("package_code", _packageCode)
	return nil
}

// Get PackageCode Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetPackageCode() string {
	return r._packageCode
}

// Set is Married Setter
// 婚否(0-未婚; 1-已婚)
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetMarried(_married string) error {
	r._married = _married
	r.Set("married", _married)
	return nil
}

// Get Married Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetMarried() string {
	return r._married
}

// Set is StoreId Setter
// 店铺ID
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetStoreId() string {
	return r._storeId
}

// Set is Phone Setter
// 电话号码
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// Get Phone Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetPhone() string {
	return r._phone
}

// Set is CertType Setter
// 证件类型(0-身份证; 1-护照; 2-军官证)
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetCertType(_certType string) error {
	r._certType = _certType
	r.Set("cert_type", _certType)
	return nil
}

// Get CertType Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetCertType() string {
	return r._certType
}

// Set is CertNumber Setter
// 证件号
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetCertNumber(_certNumber string) error {
	r._certNumber = _certNumber
	r.Set("cert_number", _certNumber)
	return nil
}

// Get CertNumber Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetCertNumber() string {
	return r._certNumber
}

// Set is Company Setter
// 所属公司
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetCompany(_company string) error {
	r._company = _company
	r.Set("company", _company)
	return nil
}

// Get Company Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetCompany() string {
	return r._company
}

// Set is Department Setter
// 所属部门
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetDepartment(_department string) error {
	r._department = _department
	r.Set("department", _department)
	return nil
}

// Get Department Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetDepartment() string {
	return r._department
}

// Set is Address Setter
// 报告邮寄地址
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// Get Address Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetAddress() string {
	return r._address
}

// Set is AddItems Setter
// 加项列表
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetAddItems(_addItems []AddItem) error {
	r._addItems = _addItems
	r.Set("add_items", _addItems)
	return nil
}

// Get AddItems Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetAddItems() []AddItem {
	return r._addItems
}

// Set is AddPacks Setter
// 加项包列表
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetAddPacks(_addPacks []AddPack) error {
	r._addPacks = _addPacks
	r.Set("add_packs", _addPacks)
	return nil
}

// Get AddPacks Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetAddPacks() []AddPack {
	return r._addPacks
}

// Set is HavaReport Setter
// 0没报告1有报告
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetHavaReport(_havaReport string) error {
	r._havaReport = _havaReport
	r.Set("hava_report", _havaReport)
	return nil
}

// Get HavaReport Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetHavaReport() string {
	return r._havaReport
}

// Set is EmployeeNumber Setter
// 员工号
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetEmployeeNumber(_employeeNumber string) error {
	r._employeeNumber = _employeeNumber
	r.Set("employee_number", _employeeNumber)
	return nil
}

// Get EmployeeNumber Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetEmployeeNumber() string {
	return r._employeeNumber
}

// Set is ServiceType Setter
// 服务类型，ONSITE_SERVICE（到店检测）、DOOR_TO_DOOR_SERVICE（上门检测）
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetServiceType(_serviceType string) error {
	r._serviceType = _serviceType
	r.Set("service_type", _serviceType)
	return nil
}

// Get ServiceType Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetServiceType() string {
	return r._serviceType
}

// Set is ServiceAddress Setter
// 上门服务的上门地址
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetServiceAddress(_serviceAddress *AddAddress) error {
	r._serviceAddress = _serviceAddress
	r.Set("service_address", _serviceAddress)
	return nil
}

// Get ServiceAddress Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetServiceAddress() *AddAddress {
	return r._serviceAddress
}

// Set is ReserveTimeStart Setter
// 预约时间段开始
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetReserveTimeStart(_reserveTimeStart string) error {
	r._reserveTimeStart = _reserveTimeStart
	r.Set("reserve_time_start", _reserveTimeStart)
	return nil
}

// Get ReserveTimeStart Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetReserveTimeStart() string {
	return r._reserveTimeStart
}

// Set is ReserveTimeEnd Setter
// 预约时间段截止
func (r *AlibabaAlihealthExaminationReserveConfirmAPIRequest) SetReserveTimeEnd(_reserveTimeEnd string) error {
	r._reserveTimeEnd = _reserveTimeEnd
	r.Set("reserve_time_end", _reserveTimeEnd)
	return nil
}

// Get ReserveTimeEnd Getter
func (r AlibabaAlihealthExaminationReserveConfirmAPIRequest) GetReserveTimeEnd() string {
	return r._reserveTimeEnd
}
