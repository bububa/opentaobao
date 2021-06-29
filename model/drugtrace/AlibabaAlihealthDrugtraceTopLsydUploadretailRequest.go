package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售单据上传接口 APIRequest
alibaba.alihealth.drugtrace.top.lsyd.uploadretail

快易通多融零售上传接口
*/
type AlibabaAlihealthDrugtraceTopLsydUploadretailRequest struct {
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

    // 购买人电话
    userTel   string 

    // 互联网订单标识 0非互联网 1互联网
    networkBillFlag   string 

    // 开药医师
    medicDoctor   string 

    // 药品发药人
    medicDispenser   string 

    // 药品使用者姓名
    userName   string 

    // 药品代理人
    userAgent   string 

}

func NewAlibabaAlihealthDrugtraceTopLsydUploadretailRequest() *AlibabaAlihealthDrugtraceTopLsydUploadretailRequest{
    return &AlibabaAlihealthDrugtraceTopLsydUploadretailRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.lsyd.uploadretail"
}

func (r AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) SetBillTime(billTime string) error {
    r.billTime = billTime
    r.Set("bill_time", billTime)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) GetBillTime() string {
    return r.billTime
}

func (r *AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) SetBillType(billType int64) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) GetBillType() int64 {
    return r.billType
}

func (r *AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) SetPhysicType(physicType int64) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) GetPhysicType() int64 {
    return r.physicType
}

func (r *AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) SetRefUserId(refUserId string) error {
    r.refUserId = refUserId
    r.Set("ref_user_id", refUserId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) GetRefUserId() string {
    return r.refUserId
}

func (r *AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) SetFromUserId(fromUserId string) error {
    r.fromUserId = fromUserId
    r.Set("from_user_id", fromUserId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) GetFromUserId() string {
    return r.fromUserId
}

func (r *AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) SetOperIcCode(operIcCode string) error {
    r.operIcCode = operIcCode
    r.Set("oper_ic_code", operIcCode)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) GetOperIcCode() string {
    return r.operIcCode
}

func (r *AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) SetOperIcName(operIcName string) error {
    r.operIcName = operIcName
    r.Set("oper_ic_name", operIcName)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) GetOperIcName() string {
    return r.operIcName
}

func (r *AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) SetTraceCodes(traceCodes []string) error {
    r.traceCodes = traceCodes
    r.Set("trace_codes", traceCodes)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) GetTraceCodes() []string {
    return r.traceCodes
}

func (r *AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) SetCustomerIdType(customerIdType string) error {
    r.customerIdType = customerIdType
    r.Set("customer_id_type", customerIdType)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) GetCustomerIdType() string {
    return r.customerIdType
}

func (r *AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) SetCustomerId(customerId string) error {
    r.customerId = customerId
    r.Set("customer_id", customerId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) GetCustomerId() string {
    return r.customerId
}

func (r *AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) SetUserTel(userTel string) error {
    r.userTel = userTel
    r.Set("user_tel", userTel)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) GetUserTel() string {
    return r.userTel
}

func (r *AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) SetNetworkBillFlag(networkBillFlag string) error {
    r.networkBillFlag = networkBillFlag
    r.Set("network_bill_flag", networkBillFlag)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) GetNetworkBillFlag() string {
    return r.networkBillFlag
}

func (r *AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) SetMedicDoctor(medicDoctor string) error {
    r.medicDoctor = medicDoctor
    r.Set("medic_doctor", medicDoctor)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) GetMedicDoctor() string {
    return r.medicDoctor
}

func (r *AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) SetMedicDispenser(medicDispenser string) error {
    r.medicDispenser = medicDispenser
    r.Set("medic_dispenser", medicDispenser)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) GetMedicDispenser() string {
    return r.medicDispenser
}

func (r *AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) SetUserName(userName string) error {
    r.userName = userName
    r.Set("user_name", userName)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) GetUserName() string {
    return r.userName
}

func (r *AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) SetUserAgent(userAgent string) error {
    r.userAgent = userAgent
    r.Set("user_agent", userAgent)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydUploadretailRequest) GetUserAgent() string {
    return r.userAgent
}

