package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
快易通多融零售上传接口 API请求
alibaba.alihealth.drug.kyt.druploadretail

快易通多融零售上传接口
*/
type AlibabaAlihealthDrugKytDruploadretailRequest struct {
    model.Params
    // 单据编号（唯一）
    billCode   string
    // 单据生成时间（一般写当前时间）
    billTime   string
    // 单据类型[321,零售出库][322,疫苗接种][116,消费者退货入库]
    billType   int64
    // 药品类型[2,特药，3,普药]
    physicType   int64
    // 码上放心平台企业唯一编码（门店或医疗机构）
    refUserId   string
    // 发货企业(可为空)
    fromUserId   string
    // 单据提交者(appkey编号)
    operIcCode   string
    // 单据提交者姓名(可为空)
    operIcName   string
    // 20位追溯码（多个时用半角逗号分隔）
    traceCodes   []string
    // 购买人证件类型【1身份证2护照3 军官证4 医保卡5接种卡6学生证9其它】
    customerIdType   string
    // 购买人证件编号
    customerId   string
    // 药品购买者电话
    userTel   string
    // 互联网标识 1是 0代表否
    networkBillFlag   string
    // 医师名称
    medicDoctor   string
    // 药品发药者
    medicDispenser   string
    // 药品使用者名称
    userName   string
    // 药品使用者代理人
    userAgent   string
}

// 初始化AlibabaAlihealthDrugKytDruploadretailRequest对象
func NewAlibabaAlihealthDrugKytDruploadretailRequest() *AlibabaAlihealthDrugKytDruploadretailRequest{
    return &AlibabaAlihealthDrugKytDruploadretailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDruploadretailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.druploadretail"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDruploadretailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillCode Setter
// 单据编号（唯一）
func (r *AlibabaAlihealthDrugKytDruploadretailRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytDruploadretailRequest) GetBillCode() string {
    return r.billCode
}
// BillTime Setter
// 单据生成时间（一般写当前时间）
func (r *AlibabaAlihealthDrugKytDruploadretailRequest) SetBillTime(billTime string) error {
    r.billTime = billTime
    r.Set("bill_time", billTime)
    return nil
}

// BillTime Getter
func (r AlibabaAlihealthDrugKytDruploadretailRequest) GetBillTime() string {
    return r.billTime
}
// BillType Setter
// 单据类型[321,零售出库][322,疫苗接种][116,消费者退货入库]
func (r *AlibabaAlihealthDrugKytDruploadretailRequest) SetBillType(billType int64) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytDruploadretailRequest) GetBillType() int64 {
    return r.billType
}
// PhysicType Setter
// 药品类型[2,特药，3,普药]
func (r *AlibabaAlihealthDrugKytDruploadretailRequest) SetPhysicType(physicType int64) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

// PhysicType Getter
func (r AlibabaAlihealthDrugKytDruploadretailRequest) GetPhysicType() int64 {
    return r.physicType
}
// RefUserId Setter
// 码上放心平台企业唯一编码（门店或医疗机构）
func (r *AlibabaAlihealthDrugKytDruploadretailRequest) SetRefUserId(refUserId string) error {
    r.refUserId = refUserId
    r.Set("ref_user_id", refUserId)
    return nil
}

// RefUserId Getter
func (r AlibabaAlihealthDrugKytDruploadretailRequest) GetRefUserId() string {
    return r.refUserId
}
// FromUserId Setter
// 发货企业(可为空)
func (r *AlibabaAlihealthDrugKytDruploadretailRequest) SetFromUserId(fromUserId string) error {
    r.fromUserId = fromUserId
    r.Set("from_user_id", fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugKytDruploadretailRequest) GetFromUserId() string {
    return r.fromUserId
}
// OperIcCode Setter
// 单据提交者(appkey编号)
func (r *AlibabaAlihealthDrugKytDruploadretailRequest) SetOperIcCode(operIcCode string) error {
    r.operIcCode = operIcCode
    r.Set("oper_ic_code", operIcCode)
    return nil
}

// OperIcCode Getter
func (r AlibabaAlihealthDrugKytDruploadretailRequest) GetOperIcCode() string {
    return r.operIcCode
}
// OperIcName Setter
// 单据提交者姓名(可为空)
func (r *AlibabaAlihealthDrugKytDruploadretailRequest) SetOperIcName(operIcName string) error {
    r.operIcName = operIcName
    r.Set("oper_ic_name", operIcName)
    return nil
}

// OperIcName Getter
func (r AlibabaAlihealthDrugKytDruploadretailRequest) GetOperIcName() string {
    return r.operIcName
}
// TraceCodes Setter
// 20位追溯码（多个时用半角逗号分隔）
func (r *AlibabaAlihealthDrugKytDruploadretailRequest) SetTraceCodes(traceCodes []string) error {
    r.traceCodes = traceCodes
    r.Set("trace_codes", traceCodes)
    return nil
}

// TraceCodes Getter
func (r AlibabaAlihealthDrugKytDruploadretailRequest) GetTraceCodes() []string {
    return r.traceCodes
}
// CustomerIdType Setter
// 购买人证件类型【1身份证2护照3 军官证4 医保卡5接种卡6学生证9其它】
func (r *AlibabaAlihealthDrugKytDruploadretailRequest) SetCustomerIdType(customerIdType string) error {
    r.customerIdType = customerIdType
    r.Set("customer_id_type", customerIdType)
    return nil
}

// CustomerIdType Getter
func (r AlibabaAlihealthDrugKytDruploadretailRequest) GetCustomerIdType() string {
    return r.customerIdType
}
// CustomerId Setter
// 购买人证件编号
func (r *AlibabaAlihealthDrugKytDruploadretailRequest) SetCustomerId(customerId string) error {
    r.customerId = customerId
    r.Set("customer_id", customerId)
    return nil
}

// CustomerId Getter
func (r AlibabaAlihealthDrugKytDruploadretailRequest) GetCustomerId() string {
    return r.customerId
}
// UserTel Setter
// 药品购买者电话
func (r *AlibabaAlihealthDrugKytDruploadretailRequest) SetUserTel(userTel string) error {
    r.userTel = userTel
    r.Set("user_tel", userTel)
    return nil
}

// UserTel Getter
func (r AlibabaAlihealthDrugKytDruploadretailRequest) GetUserTel() string {
    return r.userTel
}
// NetworkBillFlag Setter
// 互联网标识 1是 0代表否
func (r *AlibabaAlihealthDrugKytDruploadretailRequest) SetNetworkBillFlag(networkBillFlag string) error {
    r.networkBillFlag = networkBillFlag
    r.Set("network_bill_flag", networkBillFlag)
    return nil
}

// NetworkBillFlag Getter
func (r AlibabaAlihealthDrugKytDruploadretailRequest) GetNetworkBillFlag() string {
    return r.networkBillFlag
}
// MedicDoctor Setter
// 医师名称
func (r *AlibabaAlihealthDrugKytDruploadretailRequest) SetMedicDoctor(medicDoctor string) error {
    r.medicDoctor = medicDoctor
    r.Set("medic_doctor", medicDoctor)
    return nil
}

// MedicDoctor Getter
func (r AlibabaAlihealthDrugKytDruploadretailRequest) GetMedicDoctor() string {
    return r.medicDoctor
}
// MedicDispenser Setter
// 药品发药者
func (r *AlibabaAlihealthDrugKytDruploadretailRequest) SetMedicDispenser(medicDispenser string) error {
    r.medicDispenser = medicDispenser
    r.Set("medic_dispenser", medicDispenser)
    return nil
}

// MedicDispenser Getter
func (r AlibabaAlihealthDrugKytDruploadretailRequest) GetMedicDispenser() string {
    return r.medicDispenser
}
// UserName Setter
// 药品使用者名称
func (r *AlibabaAlihealthDrugKytDruploadretailRequest) SetUserName(userName string) error {
    r.userName = userName
    r.Set("user_name", userName)
    return nil
}

// UserName Getter
func (r AlibabaAlihealthDrugKytDruploadretailRequest) GetUserName() string {
    return r.userName
}
// UserAgent Setter
// 药品使用者代理人
func (r *AlibabaAlihealthDrugKytDruploadretailRequest) SetUserAgent(userAgent string) error {
    r.userAgent = userAgent
    r.Set("user_agent", userAgent)
    return nil
}

// UserAgent Getter
func (r AlibabaAlihealthDrugKytDruploadretailRequest) GetUserAgent() string {
    return r.userAgent
}
