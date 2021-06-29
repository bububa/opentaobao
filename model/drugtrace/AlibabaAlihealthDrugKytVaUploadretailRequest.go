package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售单据上传接口 APIRequest
alibaba.alihealth.drug.kyt.va.uploadretail

零售上传单据信息接口
*/
type AlibabaAlihealthDrugKytVaUploadretailRequest struct {
    model.Params

    // 单据编号（唯一）
    billCode   string 

    // 单据生成时间（一般写当前时间）
    billTime   string 

    // 单据类型[321,零售出库][322,疫苗接种]
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

    // 用药人名称
    userTel   string 

    // 互联网标志 1是 0否
    networkBillFlag   string 

    // 医师名称
    medicDoctor   string 

    // 药品发药人
    medicDispenser   string 

    // 药品使用者：
    userName   string 

    // 药品使用者代理人
    userAgent   string 

}

func NewAlibabaAlihealthDrugKytVaUploadretailRequest() *AlibabaAlihealthDrugKytVaUploadretailRequest{
    return &AlibabaAlihealthDrugKytVaUploadretailRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytVaUploadretailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.va.uploadretail"
}

func (r AlibabaAlihealthDrugKytVaUploadretailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytVaUploadretailRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthDrugKytVaUploadretailRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthDrugKytVaUploadretailRequest) SetBillTime(billTime string) error {
    r.billTime = billTime
    r.Set("bill_time", billTime)
    return nil
}

func (r AlibabaAlihealthDrugKytVaUploadretailRequest) GetBillTime() string {
    return r.billTime
}

func (r *AlibabaAlihealthDrugKytVaUploadretailRequest) SetBillType(billType int64) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

func (r AlibabaAlihealthDrugKytVaUploadretailRequest) GetBillType() int64 {
    return r.billType
}

func (r *AlibabaAlihealthDrugKytVaUploadretailRequest) SetPhysicType(physicType int64) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

func (r AlibabaAlihealthDrugKytVaUploadretailRequest) GetPhysicType() int64 {
    return r.physicType
}

func (r *AlibabaAlihealthDrugKytVaUploadretailRequest) SetRefUserId(refUserId string) error {
    r.refUserId = refUserId
    r.Set("ref_user_id", refUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytVaUploadretailRequest) GetRefUserId() string {
    return r.refUserId
}

func (r *AlibabaAlihealthDrugKytVaUploadretailRequest) SetFromUserId(fromUserId string) error {
    r.fromUserId = fromUserId
    r.Set("from_user_id", fromUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytVaUploadretailRequest) GetFromUserId() string {
    return r.fromUserId
}

func (r *AlibabaAlihealthDrugKytVaUploadretailRequest) SetOperIcCode(operIcCode string) error {
    r.operIcCode = operIcCode
    r.Set("oper_ic_code", operIcCode)
    return nil
}

func (r AlibabaAlihealthDrugKytVaUploadretailRequest) GetOperIcCode() string {
    return r.operIcCode
}

func (r *AlibabaAlihealthDrugKytVaUploadretailRequest) SetOperIcName(operIcName string) error {
    r.operIcName = operIcName
    r.Set("oper_ic_name", operIcName)
    return nil
}

func (r AlibabaAlihealthDrugKytVaUploadretailRequest) GetOperIcName() string {
    return r.operIcName
}

func (r *AlibabaAlihealthDrugKytVaUploadretailRequest) SetTraceCodes(traceCodes []string) error {
    r.traceCodes = traceCodes
    r.Set("trace_codes", traceCodes)
    return nil
}

func (r AlibabaAlihealthDrugKytVaUploadretailRequest) GetTraceCodes() []string {
    return r.traceCodes
}

func (r *AlibabaAlihealthDrugKytVaUploadretailRequest) SetCustomerIdType(customerIdType string) error {
    r.customerIdType = customerIdType
    r.Set("customer_id_type", customerIdType)
    return nil
}

func (r AlibabaAlihealthDrugKytVaUploadretailRequest) GetCustomerIdType() string {
    return r.customerIdType
}

func (r *AlibabaAlihealthDrugKytVaUploadretailRequest) SetCustomerId(customerId string) error {
    r.customerId = customerId
    r.Set("customer_id", customerId)
    return nil
}

func (r AlibabaAlihealthDrugKytVaUploadretailRequest) GetCustomerId() string {
    return r.customerId
}

func (r *AlibabaAlihealthDrugKytVaUploadretailRequest) SetUserTel(userTel string) error {
    r.userTel = userTel
    r.Set("user_tel", userTel)
    return nil
}

func (r AlibabaAlihealthDrugKytVaUploadretailRequest) GetUserTel() string {
    return r.userTel
}

func (r *AlibabaAlihealthDrugKytVaUploadretailRequest) SetNetworkBillFlag(networkBillFlag string) error {
    r.networkBillFlag = networkBillFlag
    r.Set("network_bill_flag", networkBillFlag)
    return nil
}

func (r AlibabaAlihealthDrugKytVaUploadretailRequest) GetNetworkBillFlag() string {
    return r.networkBillFlag
}

func (r *AlibabaAlihealthDrugKytVaUploadretailRequest) SetMedicDoctor(medicDoctor string) error {
    r.medicDoctor = medicDoctor
    r.Set("medic_doctor", medicDoctor)
    return nil
}

func (r AlibabaAlihealthDrugKytVaUploadretailRequest) GetMedicDoctor() string {
    return r.medicDoctor
}

func (r *AlibabaAlihealthDrugKytVaUploadretailRequest) SetMedicDispenser(medicDispenser string) error {
    r.medicDispenser = medicDispenser
    r.Set("medic_dispenser", medicDispenser)
    return nil
}

func (r AlibabaAlihealthDrugKytVaUploadretailRequest) GetMedicDispenser() string {
    return r.medicDispenser
}

func (r *AlibabaAlihealthDrugKytVaUploadretailRequest) SetUserName(userName string) error {
    r.userName = userName
    r.Set("user_name", userName)
    return nil
}

func (r AlibabaAlihealthDrugKytVaUploadretailRequest) GetUserName() string {
    return r.userName
}

func (r *AlibabaAlihealthDrugKytVaUploadretailRequest) SetUserAgent(userAgent string) error {
    r.userAgent = userAgent
    r.Set("user_agent", userAgent)
    return nil
}

func (r AlibabaAlihealthDrugKytVaUploadretailRequest) GetUserAgent() string {
    return r.userAgent
}

