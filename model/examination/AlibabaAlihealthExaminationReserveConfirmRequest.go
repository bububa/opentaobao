package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构对接_体检套餐预定确认 APIRequest
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

func NewAlibabaAlihealthExaminationReserveConfirmRequest() *AlibabaAlihealthExaminationReserveConfirmRequest{
    return &AlibabaAlihealthExaminationReserveConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.reserve.confirm"
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetMerchantCode(merchantCode string) error {
    r.merchantCode = merchantCode
    r.Set("merchant_code", merchantCode)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetMerchantCode() string {
    return r.merchantCode
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetName() string {
    return r.name
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetReserveNumber(reserveNumber string) error {
    r.reserveNumber = reserveNumber
    r.Set("reserve_number", reserveNumber)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetReserveNumber() string {
    return r.reserveNumber
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetGender(gender string) error {
    r.gender = gender
    r.Set("gender", gender)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetGender() string {
    return r.gender
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetBirthday(birthday string) error {
    r.birthday = birthday
    r.Set("birthday", birthday)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetBirthday() string {
    return r.birthday
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetReserveDate(reserveDate string) error {
    r.reserveDate = reserveDate
    r.Set("reserve_date", reserveDate)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetReserveDate() string {
    return r.reserveDate
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetPackageCode(packageCode string) error {
    r.packageCode = packageCode
    r.Set("package_code", packageCode)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetPackageCode() string {
    return r.packageCode
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetMarried(married string) error {
    r.married = married
    r.Set("married", married)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetMarried() string {
    return r.married
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetStoreId() string {
    return r.storeId
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetPhone() string {
    return r.phone
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetCertType(certType string) error {
    r.certType = certType
    r.Set("cert_type", certType)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetCertType() string {
    return r.certType
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetCertNumber(certNumber string) error {
    r.certNumber = certNumber
    r.Set("cert_number", certNumber)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetCertNumber() string {
    return r.certNumber
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetCompany(company string) error {
    r.company = company
    r.Set("company", company)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetCompany() string {
    return r.company
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetDepartment(department string) error {
    r.department = department
    r.Set("department", department)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetDepartment() string {
    return r.department
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetAddress() string {
    return r.address
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetAddItems(addItems []AddItem) error {
    r.addItems = addItems
    r.Set("add_items", addItems)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetAddItems() []AddItem {
    return r.addItems
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetAddPacks(addPacks []AddPack) error {
    r.addPacks = addPacks
    r.Set("add_packs", addPacks)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetAddPacks() []AddPack {
    return r.addPacks
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetHavaReport(havaReport string) error {
    r.havaReport = havaReport
    r.Set("hava_report", havaReport)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetHavaReport() string {
    return r.havaReport
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetEmployeeNumber(employeeNumber string) error {
    r.employeeNumber = employeeNumber
    r.Set("employee_number", employeeNumber)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetEmployeeNumber() string {
    return r.employeeNumber
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetServiceType(serviceType string) error {
    r.serviceType = serviceType
    r.Set("service_type", serviceType)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetServiceType() string {
    return r.serviceType
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetServiceAddress(serviceAddress *AddAddress) error {
    r.serviceAddress = serviceAddress
    r.Set("service_address", serviceAddress)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetServiceAddress() *AddAddress {
    return r.serviceAddress
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetReserveTimeStart(reserveTimeStart string) error {
    r.reserveTimeStart = reserveTimeStart
    r.Set("reserve_time_start", reserveTimeStart)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetReserveTimeStart() string {
    return r.reserveTimeStart
}

func (r *AlibabaAlihealthExaminationReserveConfirmRequest) SetReserveTimeEnd(reserveTimeEnd string) error {
    r.reserveTimeEnd = reserveTimeEnd
    r.Set("reserve_time_end", reserveTimeEnd)
    return nil
}

func (r AlibabaAlihealthExaminationReserveConfirmRequest) GetReserveTimeEnd() string {
    return r.reserveTimeEnd
}

