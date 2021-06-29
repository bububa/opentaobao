package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
医院上传出库信息 APIRequest
alibaba.alihealth.drug.kyt.yy.uploadretail

医院上传出库信息接口
*/
type AlibabaAlihealthDrugKytYyUploadretailRequest struct {
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

    // 药房名称(可为空)
    drugstoreName   string 

    // 20位追溯码（多个时用半角逗号分隔）
    traceCodes   []string 

    // 购买人证件类型【1身份证2护照3 军官证4 医保卡5接种卡6学生证9其它】
    customerIdType   string 

    // 购买人证件编号
    customerId   string 

    // 购买人电话：
    userTel   string 

    // 互联网标识 1是 0代表否
    networkBillFlag   string 

    // 开药医师名称
    medicDoctor   string 

    // 药品使用分药者
    medicDispenser   string 

    // 药品使用者名称
    userName   string 

    // 药品使用者代理人
    userAgent   string 

}

func NewAlibabaAlihealthDrugKytYyUploadretailRequest() *AlibabaAlihealthDrugKytYyUploadretailRequest{
    return &AlibabaAlihealthDrugKytYyUploadretailRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytYyUploadretailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.yy.uploadretail"
}

func (r AlibabaAlihealthDrugKytYyUploadretailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytYyUploadretailRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadretailRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthDrugKytYyUploadretailRequest) SetBillTime(billTime string) error {
    r.billTime = billTime
    r.Set("bill_time", billTime)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadretailRequest) GetBillTime() string {
    return r.billTime
}

func (r *AlibabaAlihealthDrugKytYyUploadretailRequest) SetBillType(billType int64) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadretailRequest) GetBillType() int64 {
    return r.billType
}

func (r *AlibabaAlihealthDrugKytYyUploadretailRequest) SetPhysicType(physicType int64) error {
    r.physicType = physicType
    r.Set("physic_type", physicType)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadretailRequest) GetPhysicType() int64 {
    return r.physicType
}

func (r *AlibabaAlihealthDrugKytYyUploadretailRequest) SetRefUserId(refUserId string) error {
    r.refUserId = refUserId
    r.Set("ref_user_id", refUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadretailRequest) GetRefUserId() string {
    return r.refUserId
}

func (r *AlibabaAlihealthDrugKytYyUploadretailRequest) SetFromUserId(fromUserId string) error {
    r.fromUserId = fromUserId
    r.Set("from_user_id", fromUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadretailRequest) GetFromUserId() string {
    return r.fromUserId
}

func (r *AlibabaAlihealthDrugKytYyUploadretailRequest) SetOperIcCode(operIcCode string) error {
    r.operIcCode = operIcCode
    r.Set("oper_ic_code", operIcCode)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadretailRequest) GetOperIcCode() string {
    return r.operIcCode
}

func (r *AlibabaAlihealthDrugKytYyUploadretailRequest) SetDrugstoreName(drugstoreName string) error {
    r.drugstoreName = drugstoreName
    r.Set("drugstore_name", drugstoreName)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadretailRequest) GetDrugstoreName() string {
    return r.drugstoreName
}

func (r *AlibabaAlihealthDrugKytYyUploadretailRequest) SetTraceCodes(traceCodes []string) error {
    r.traceCodes = traceCodes
    r.Set("trace_codes", traceCodes)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadretailRequest) GetTraceCodes() []string {
    return r.traceCodes
}

func (r *AlibabaAlihealthDrugKytYyUploadretailRequest) SetCustomerIdType(customerIdType string) error {
    r.customerIdType = customerIdType
    r.Set("customer_id_type", customerIdType)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadretailRequest) GetCustomerIdType() string {
    return r.customerIdType
}

func (r *AlibabaAlihealthDrugKytYyUploadretailRequest) SetCustomerId(customerId string) error {
    r.customerId = customerId
    r.Set("customer_id", customerId)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadretailRequest) GetCustomerId() string {
    return r.customerId
}

func (r *AlibabaAlihealthDrugKytYyUploadretailRequest) SetUserTel(userTel string) error {
    r.userTel = userTel
    r.Set("user_tel", userTel)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadretailRequest) GetUserTel() string {
    return r.userTel
}

func (r *AlibabaAlihealthDrugKytYyUploadretailRequest) SetNetworkBillFlag(networkBillFlag string) error {
    r.networkBillFlag = networkBillFlag
    r.Set("network_bill_flag", networkBillFlag)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadretailRequest) GetNetworkBillFlag() string {
    return r.networkBillFlag
}

func (r *AlibabaAlihealthDrugKytYyUploadretailRequest) SetMedicDoctor(medicDoctor string) error {
    r.medicDoctor = medicDoctor
    r.Set("medic_doctor", medicDoctor)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadretailRequest) GetMedicDoctor() string {
    return r.medicDoctor
}

func (r *AlibabaAlihealthDrugKytYyUploadretailRequest) SetMedicDispenser(medicDispenser string) error {
    r.medicDispenser = medicDispenser
    r.Set("medic_dispenser", medicDispenser)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadretailRequest) GetMedicDispenser() string {
    return r.medicDispenser
}

func (r *AlibabaAlihealthDrugKytYyUploadretailRequest) SetUserName(userName string) error {
    r.userName = userName
    r.Set("user_name", userName)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadretailRequest) GetUserName() string {
    return r.userName
}

func (r *AlibabaAlihealthDrugKytYyUploadretailRequest) SetUserAgent(userAgent string) error {
    r.userAgent = userAgent
    r.Set("user_agent", userAgent)
    return nil
}

func (r AlibabaAlihealthDrugKytYyUploadretailRequest) GetUserAgent() string {
    return r.userAgent
}

