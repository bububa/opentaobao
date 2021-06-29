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
    merchantCode   string
    // 体检人姓名
    name   string
    // 阿里健康预约唯一标识
    reserveNumber   string
    // 性别(0-男;1-女;)
    gender   string
    // 出生日期
    birthday   string
    // 预约时间
    reserveDate   string
    // 体检套餐编码
    packageCode   string
    // 婚否(0-未婚; 1-已婚)
    married   string
    // 店铺ID
    storeId   string
    // 电话号码
    phone   string
    // 证件类型(0-身份证; 1-护照; 2-军官证)
    certType   string
    // 证件号
    certNumber   string
    // 所属公司
    company   string
    // 所属部门
    department   string
    // 报告邮寄地址
    address   string
    // 加项列表
    addItems   []AddItem
    // 加项包列表
    addPacks   []AddPack
    // 0没报告1有报告
    havaReport   string
    // 员工号
    employeeNumber   string
    // 服务类型，ONSITE_SERVICE（到店检测）、DOOR_TO_DOOR_SERVICE（上门检测）
    serviceType   string
    // 上门服务的上门地址
    serviceAddress   *AddAddress
    // 预约时间段开始
    reserveTimeStart   string
    // 预约时间段截止
    reserveTimeEnd   string
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
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetMerchantCode(merchantCode string) error {
    r.merchantCode = merchantCode
    r.Set("merchant_code", merchantCode)
    return nil
}

// MerchantCode Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetMerchantCode() string {
    return r.merchantCode
}
// Name Setter
// 体检人姓名
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetName() string {
    return r.name
}
// ReserveNumber Setter
// 阿里健康预约唯一标识
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetReserveNumber(reserveNumber string) error {
    r.reserveNumber = reserveNumber
    r.Set("reserve_number", reserveNumber)
    return nil
}

// ReserveNumber Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetReserveNumber() string {
    return r.reserveNumber
}
// Gender Setter
// 性别(0-男;1-女;)
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetGender(gender string) error {
    r.gender = gender
    r.Set("gender", gender)
    return nil
}

// Gender Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetGender() string {
    return r.gender
}
// Birthday Setter
// 出生日期
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetBirthday(birthday string) error {
    r.birthday = birthday
    r.Set("birthday", birthday)
    return nil
}

// Birthday Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetBirthday() string {
    return r.birthday
}
// ReserveDate Setter
// 预约时间
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetReserveDate(reserveDate string) error {
    r.reserveDate = reserveDate
    r.Set("reserve_date", reserveDate)
    return nil
}

// ReserveDate Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetReserveDate() string {
    return r.reserveDate
}
// PackageCode Setter
// 体检套餐编码
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetPackageCode(packageCode string) error {
    r.packageCode = packageCode
    r.Set("package_code", packageCode)
    return nil
}

// PackageCode Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetPackageCode() string {
    return r.packageCode
}
// Married Setter
// 婚否(0-未婚; 1-已婚)
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetMarried(married string) error {
    r.married = married
    r.Set("married", married)
    return nil
}

// Married Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetMarried() string {
    return r.married
}
// StoreId Setter
// 店铺ID
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetStoreId() string {
    return r.storeId
}
// Phone Setter
// 电话号码
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

// Phone Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetPhone() string {
    return r.phone
}
// CertType Setter
// 证件类型(0-身份证; 1-护照; 2-军官证)
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetCertType(certType string) error {
    r.certType = certType
    r.Set("cert_type", certType)
    return nil
}

// CertType Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetCertType() string {
    return r.certType
}
// CertNumber Setter
// 证件号
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetCertNumber(certNumber string) error {
    r.certNumber = certNumber
    r.Set("cert_number", certNumber)
    return nil
}

// CertNumber Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetCertNumber() string {
    return r.certNumber
}
// Company Setter
// 所属公司
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetCompany(company string) error {
    r.company = company
    r.Set("company", company)
    return nil
}

// Company Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetCompany() string {
    return r.company
}
// Department Setter
// 所属部门
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetDepartment(department string) error {
    r.department = department
    r.Set("department", department)
    return nil
}

// Department Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetDepartment() string {
    return r.department
}
// Address Setter
// 报告邮寄地址
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

// Address Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetAddress() string {
    return r.address
}
// AddItems Setter
// 加项列表
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetAddItems(addItems []AddItem) error {
    r.addItems = addItems
    r.Set("add_items", addItems)
    return nil
}

// AddItems Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetAddItems() []AddItem {
    return r.addItems
}
// AddPacks Setter
// 加项包列表
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetAddPacks(addPacks []AddPack) error {
    r.addPacks = addPacks
    r.Set("add_packs", addPacks)
    return nil
}

// AddPacks Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetAddPacks() []AddPack {
    return r.addPacks
}
// HavaReport Setter
// 0没报告1有报告
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetHavaReport(havaReport string) error {
    r.havaReport = havaReport
    r.Set("hava_report", havaReport)
    return nil
}

// HavaReport Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetHavaReport() string {
    return r.havaReport
}
// EmployeeNumber Setter
// 员工号
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetEmployeeNumber(employeeNumber string) error {
    r.employeeNumber = employeeNumber
    r.Set("employee_number", employeeNumber)
    return nil
}

// EmployeeNumber Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetEmployeeNumber() string {
    return r.employeeNumber
}
// ServiceType Setter
// 服务类型，ONSITE_SERVICE（到店检测）、DOOR_TO_DOOR_SERVICE（上门检测）
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetServiceType(serviceType string) error {
    r.serviceType = serviceType
    r.Set("service_type", serviceType)
    return nil
}

// ServiceType Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetServiceType() string {
    return r.serviceType
}
// ServiceAddress Setter
// 上门服务的上门地址
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetServiceAddress(serviceAddress *AddAddress) error {
    r.serviceAddress = serviceAddress
    r.Set("service_address", serviceAddress)
    return nil
}

// ServiceAddress Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetServiceAddress() *AddAddress {
    return r.serviceAddress
}
// ReserveTimeStart Setter
// 预约时间段开始
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetReserveTimeStart(reserveTimeStart string) error {
    r.reserveTimeStart = reserveTimeStart
    r.Set("reserve_time_start", reserveTimeStart)
    return nil
}

// ReserveTimeStart Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetReserveTimeStart() string {
    return r.reserveTimeStart
}
// ReserveTimeEnd Setter
// 预约时间段截止
func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetReserveTimeEnd(reserveTimeEnd string) error {
    r.reserveTimeEnd = reserveTimeEnd
    r.Set("reserve_time_end", reserveTimeEnd)
    return nil
}

// ReserveTimeEnd Getter
func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetReserveTimeEnd() string {
    return r.reserveTimeEnd
}
