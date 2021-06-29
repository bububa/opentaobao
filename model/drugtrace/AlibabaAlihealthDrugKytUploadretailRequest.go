package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店销售扫码回传接口 API请求
alibaba.alihealth.drug.kyt.uploadretail

门店在销售给顾客时，扫描追溯码的数据按照单据回传；
*/
type AlibabaAlihealthDrugKytUploadretailRequest struct {
    model.Params
    // 单据编号（唯一）
    billCode   string
    // 单据生成时间
    billTime   string
    // 单据类型[321,零售出库][322,疫苗接种]
    billType   int64
    // 药品类型[3,普药]
    physicType   int64
    // 码上放心平台企业编码（门店）
    refUserId   string
    // 发货企业(可为空)
    fromUserId   string
    // 单据提交者(appkey编号)
    operIcCode   string
    // 单据提交者姓名(可为空)
    operIcName   string
    // 请求类型[暂定都写2]
    clientType   string
    // 追溯码[多个时用逗号分开]
    traceCodes   []string
    // 患者电话
    userTel   string
    // 互联网标识
    networkBillFlag   string
    // 医师名单
    medicDoctor   string
    // 药品分发者
    medicDispenser   string
    // 患者名称
    userName   string
    // 患者代理领药人
    userAgent   string
    // 购买人证件类型【1身份证2护照3 军官证4 医保卡5接种卡6学生证9其它】
    customerIdType   string
    // 购买人证件编号
    customerId   string
}

// 初始化AlibabaAlihealthDrugKytUploadretailRequest对象
func NewAlibabaAlihealthDrugKytUploadretailRequest() *AlibabaAlihealthDrugKytUploadretailRequest{
    return &AlibabaAlihealthDrugKytUploadretailRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.uploadretail"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillCode Setter
// 单据编号（唯一）
func (r *AlibabaAlihealthDrugKytUploadretailRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetBillCode() string {
    return r.billCode
}
// BillTime Setter
// 单据生成时间
func (r *AlibabaAlihealthDrugKytUploadretailRequest) SetBillTime(billTime string) error {
    r.billTime = billTime
    r.Set("bill_time", billTime)
    return nil
}

// BillTime Getter
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetBillTime() string {
    return r.billTime
}
// BillType Setter
// 单据类型[321,零售出库][322,疫苗接种]
func (r *AlibabaAlihealthDrugKytUploadretailRequest) SetBillType(billType int64) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetBillType() int64 {
    return r.billType
}
// PhysicType Setter
// 药品类型[3,普药]
func (r *AlibabaAlihealthDrugKytUploadretailRequest) SetPhysicType(physicType int64) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

// PhysicType Getter
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetPhysicType() int64 {
    return r.physicType
}
// RefUserId Setter
// 码上放心平台企业编码（门店）
func (r *AlibabaAlihealthDrugKytUploadretailRequest) SetRefUserId(refUserId string) error {
    r.refUserId = refUserId
    r.Set("ref_user_id", refUserId)
    return nil
}

// RefUserId Getter
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetRefUserId() string {
    return r.refUserId
}
// FromUserId Setter
// 发货企业(可为空)
func (r *AlibabaAlihealthDrugKytUploadretailRequest) SetFromUserId(fromUserId string) error {
    r.fromUserId = fromUserId
    r.Set("from_user_id", fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetFromUserId() string {
    return r.fromUserId
}
// OperIcCode Setter
// 单据提交者(appkey编号)
func (r *AlibabaAlihealthDrugKytUploadretailRequest) SetOperIcCode(operIcCode string) error {
    r.operIcCode = operIcCode
    r.Set("oper_ic_code", operIcCode)
    return nil
}

// OperIcCode Getter
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetOperIcCode() string {
    return r.operIcCode
}
// OperIcName Setter
// 单据提交者姓名(可为空)
func (r *AlibabaAlihealthDrugKytUploadretailRequest) SetOperIcName(operIcName string) error {
    r.operIcName = operIcName
    r.Set("oper_ic_name", operIcName)
    return nil
}

// OperIcName Getter
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetOperIcName() string {
    return r.operIcName
}
// ClientType Setter
// 请求类型[暂定都写2]
func (r *AlibabaAlihealthDrugKytUploadretailRequest) SetClientType(clientType string) error {
    r.clientType = clientType
    r.Set("client_type", clientType)
    return nil
}

// ClientType Getter
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetClientType() string {
    return r.clientType
}
// TraceCodes Setter
// 追溯码[多个时用逗号分开]
func (r *AlibabaAlihealthDrugKytUploadretailRequest) SetTraceCodes(traceCodes []string) error {
    r.traceCodes = traceCodes
    r.Set("trace_codes", traceCodes)
    return nil
}

// TraceCodes Getter
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetTraceCodes() []string {
    return r.traceCodes
}
// UserTel Setter
// 患者电话
func (r *AlibabaAlihealthDrugKytUploadretailRequest) SetUserTel(userTel string) error {
    r.userTel = userTel
    r.Set("user_tel", userTel)
    return nil
}

// UserTel Getter
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetUserTel() string {
    return r.userTel
}
// NetworkBillFlag Setter
// 互联网标识
func (r *AlibabaAlihealthDrugKytUploadretailRequest) SetNetworkBillFlag(networkBillFlag string) error {
    r.networkBillFlag = networkBillFlag
    r.Set("network_bill_flag", networkBillFlag)
    return nil
}

// NetworkBillFlag Getter
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetNetworkBillFlag() string {
    return r.networkBillFlag
}
// MedicDoctor Setter
// 医师名单
func (r *AlibabaAlihealthDrugKytUploadretailRequest) SetMedicDoctor(medicDoctor string) error {
    r.medicDoctor = medicDoctor
    r.Set("medic_doctor", medicDoctor)
    return nil
}

// MedicDoctor Getter
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetMedicDoctor() string {
    return r.medicDoctor
}
// MedicDispenser Setter
// 药品分发者
func (r *AlibabaAlihealthDrugKytUploadretailRequest) SetMedicDispenser(medicDispenser string) error {
    r.medicDispenser = medicDispenser
    r.Set("medic_dispenser", medicDispenser)
    return nil
}

// MedicDispenser Getter
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetMedicDispenser() string {
    return r.medicDispenser
}
// UserName Setter
// 患者名称
func (r *AlibabaAlihealthDrugKytUploadretailRequest) SetUserName(userName string) error {
    r.userName = userName
    r.Set("user_name", userName)
    return nil
}

// UserName Getter
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetUserName() string {
    return r.userName
}
// UserAgent Setter
// 患者代理领药人
func (r *AlibabaAlihealthDrugKytUploadretailRequest) SetUserAgent(userAgent string) error {
    r.userAgent = userAgent
    r.Set("user_agent", userAgent)
    return nil
}

// UserAgent Getter
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetUserAgent() string {
    return r.userAgent
}
// CustomerIdType Setter
// 购买人证件类型【1身份证2护照3 军官证4 医保卡5接种卡6学生证9其它】
func (r *AlibabaAlihealthDrugKytUploadretailRequest) SetCustomerIdType(customerIdType string) error {
    r.customerIdType = customerIdType
    r.Set("customer_id_type", customerIdType)
    return nil
}

// CustomerIdType Getter
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetCustomerIdType() string {
    return r.customerIdType
}
// CustomerId Setter
// 购买人证件编号
func (r *AlibabaAlihealthDrugKytUploadretailRequest) SetCustomerId(customerId string) error {
    r.customerId = customerId
    r.Set("customer_id", customerId)
    return nil
}

// CustomerId Getter
func (r AlibabaAlihealthDrugKytUploadretailRequest) GetCustomerId() string {
    return r.customerId
}
