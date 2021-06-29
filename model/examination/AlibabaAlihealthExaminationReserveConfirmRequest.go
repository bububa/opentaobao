package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构对接_体检套餐预定确认 API请求
alibaba.alihealth.examination.reserve.confirm

向体检机构确认用户购买的体检套餐信息
*/
type AlibabaAlihealthExaminationReserveConfirmRequest struct {
    model.Params
    // 商户唯一码
    _merchantCode   string
    // 体检人姓名
    _name   string
    // 阿里健康预约唯一标识
    _reserveNumber   string
    // 性别(0-男;1-女;)
    _gender   string
    // 出生日期
    _birthday   string
    // 预约时间
    _reserveDate   string
    // 体检套餐编码
    _packageCode   string
    // 婚否(0-未婚; 1-已婚)
    _married   string
    // 店铺ID
    _storeId   string
    // 电话号码
    _phone   string
    // 证件类型(0-身份证; 1-护照; 2-军官证)
    _certType   string
    // 证件号
    _certNumber   string
    // 所属公司
    _company   string
    // 所属部门
    _department   string
    // 报告邮寄地址
    _address   string
    // 加项列表
    _addItems   []AddItem
    // 加项包列表
    _addPacks   []AddPack
    // 0没报告1有报告
    _havaReport   string
    // 员工号
    _employeeNumber   string
    // 服务类型，ONSITE_SERVICE（到店检测）、DOOR_TO_DOOR_SERVICE（上门检测）
    _serviceType   string
    // 上门服务的上门地址
    _serviceAddress   *AddAddress
    // 预约时间段开始
    _reserveTimeStart   string
    // 预约时间段截止
    _reserveTimeEnd   string
}

// 初始化AlibabaAlihealthExaminationReserveConfirmRequest对象
func NewAlibabaAlihealthExaminationReserveConfirmRequest() *AlibabaAlihealthExaminationReserveConfirmRequest{
    return &AlibabaAlihealthExaminationReserveConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.reserve.confirm"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MerchantCode Setter
// 商户唯一码
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetMerchantCode(_merchantCode string) error {
    r._merchantCode = _merchantCode
    r.Set("merchant_code", _merchantCode)
    return nil
}

// MerchantCode Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetMerchantCode() string {
    return r._merchantCode
}
// Name Setter
// 体检人姓名
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetName() string {
    return r._name
}
// ReserveNumber Setter
// 阿里健康预约唯一标识
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetReserveNumber(_reserveNumber string) error {
    r._reserveNumber = _reserveNumber
    r.Set("reserve_number", _reserveNumber)
    return nil
}

// ReserveNumber Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetReserveNumber() string {
    return r._reserveNumber
}
// Gender Setter
// 性别(0-男;1-女;)
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetGender(_gender string) error {
    r._gender = _gender
    r.Set("gender", _gender)
    return nil
}

// Gender Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetGender() string {
    return r._gender
}
// Birthday Setter
// 出生日期
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetBirthday(_birthday string) error {
    r._birthday = _birthday
    r.Set("birthday", _birthday)
    return nil
}

// Birthday Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetBirthday() string {
    return r._birthday
}
// ReserveDate Setter
// 预约时间
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetReserveDate(_reserveDate string) error {
    r._reserveDate = _reserveDate
    r.Set("reserve_date", _reserveDate)
    return nil
}

// ReserveDate Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetReserveDate() string {
    return r._reserveDate
}
// PackageCode Setter
// 体检套餐编码
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetPackageCode(_packageCode string) error {
    r._packageCode = _packageCode
    r.Set("package_code", _packageCode)
    return nil
}

// PackageCode Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetPackageCode() string {
    return r._packageCode
}
// Married Setter
// 婚否(0-未婚; 1-已婚)
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetMarried(_married string) error {
    r._married = _married
    r.Set("married", _married)
    return nil
}

// Married Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetMarried() string {
    return r._married
}
// StoreId Setter
// 店铺ID
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetStoreId() string {
    return r._storeId
}
// Phone Setter
// 电话号码
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetPhone(_phone string) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetPhone() string {
    return r._phone
}
// CertType Setter
// 证件类型(0-身份证; 1-护照; 2-军官证)
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetCertType(_certType string) error {
    r._certType = _certType
    r.Set("cert_type", _certType)
    return nil
}

// CertType Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetCertType() string {
    return r._certType
}
// CertNumber Setter
// 证件号
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetCertNumber(_certNumber string) error {
    r._certNumber = _certNumber
    r.Set("cert_number", _certNumber)
    return nil
}

// CertNumber Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetCertNumber() string {
    return r._certNumber
}
// Company Setter
// 所属公司
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetCompany(_company string) error {
    r._company = _company
    r.Set("company", _company)
    return nil
}

// Company Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetCompany() string {
    return r._company
}
// Department Setter
// 所属部门
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetDepartment(_department string) error {
    r._department = _department
    r.Set("department", _department)
    return nil
}

// Department Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetDepartment() string {
    return r._department
}
// Address Setter
// 报告邮寄地址
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetAddress() string {
    return r._address
}
// AddItems Setter
// 加项列表
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetAddItems(_addItems []AddItem) error {
    r._addItems = _addItems
    r.Set("add_items", _addItems)
    return nil
}

// AddItems Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetAddItems() []AddItem {
    return r._addItems
}
// AddPacks Setter
// 加项包列表
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetAddPacks(_addPacks []AddPack) error {
    r._addPacks = _addPacks
    r.Set("add_packs", _addPacks)
    return nil
}

// AddPacks Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetAddPacks() []AddPack {
    return r._addPacks
}
// HavaReport Setter
// 0没报告1有报告
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetHavaReport(_havaReport string) error {
    r._havaReport = _havaReport
    r.Set("hava_report", _havaReport)
    return nil
}

// HavaReport Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetHavaReport() string {
    return r._havaReport
}
// EmployeeNumber Setter
// 员工号
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetEmployeeNumber(_employeeNumber string) error {
    r._employeeNumber = _employeeNumber
    r.Set("employee_number", _employeeNumber)
    return nil
}

// EmployeeNumber Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetEmployeeNumber() string {
    return r._employeeNumber
}
// ServiceType Setter
// 服务类型，ONSITE_SERVICE（到店检测）、DOOR_TO_DOOR_SERVICE（上门检测）
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetServiceType(_serviceType string) error {
    r._serviceType = _serviceType
    r.Set("service_type", _serviceType)
    return nil
}

// ServiceType Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetServiceType() string {
    return r._serviceType
}
// ServiceAddress Setter
// 上门服务的上门地址
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetServiceAddress(_serviceAddress *AddAddress) error {
    r._serviceAddress = _serviceAddress
    r.Set("service_address", _serviceAddress)
    return nil
}

// ServiceAddress Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetServiceAddress() *AddAddress {
    return r._serviceAddress
}
// ReserveTimeStart Setter
// 预约时间段开始
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetReserveTimeStart(_reserveTimeStart string) error {
    r._reserveTimeStart = _reserveTimeStart
    r.Set("reserve_time_start", _reserveTimeStart)
    return nil
}

// ReserveTimeStart Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetReserveTimeStart() string {
    return r._reserveTimeStart
}
// ReserveTimeEnd Setter
// 预约时间段截止
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetReserveTimeEnd(_reserveTimeEnd string) error {
    r._reserveTimeEnd = _reserveTimeEnd
    r.Set("reserve_time_end", _reserveTimeEnd)
    return nil
}

// ReserveTimeEnd Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetReserveTimeEnd() string {
    return r._reserveTimeEnd
}
