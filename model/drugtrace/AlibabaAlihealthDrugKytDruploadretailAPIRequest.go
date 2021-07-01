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
type AlibabaAlihealthDrugKytDruploadretailAPIRequest struct {
    model.Params
    // 单据编号（唯一）
    _billCode   string
    // 单据生成时间（一般写当前时间）
    _billTime   string
    // 单据类型[321,零售出库][322,疫苗接种][116,消费者退货入库]
    _billType   int64
    // 药品类型[2,特药，3,普药]
    _physicType   int64
    // 码上放心平台企业唯一编码（门店或医疗机构）
    _refUserId   string
    // 发货企业(可为空)
    _fromUserId   string
    // 单据提交者(appkey编号)
    _operIcCode   string
    // 单据提交者姓名(可为空)
    _operIcName   string
    // 20位追溯码（多个时用半角逗号分隔）
    _traceCodes   []string
    // 购买人证件类型【1身份证2护照3 军官证4 医保卡5接种卡6学生证9其它】
    _customerIdType   string
    // 购买人证件编号
    _customerId   string
    // 药品购买者电话
    _userTel   string
    // 互联网标识 1是 0代表否
    _networkBillFlag   string
    // 医师名称
    _medicDoctor   string
    // 药品发药者
    _medicDispenser   string
    // 药品使用者名称
    _userName   string
    // 药品使用者代理人
    _userAgent   string
}

// 初始化AlibabaAlihealthDrugKytDruploadretailAPIRequest对象
func NewAlibabaAlihealthDrugKytDruploadretailRequest() *AlibabaAlihealthDrugKytDruploadretailAPIRequest{
    return &AlibabaAlihealthDrugKytDruploadretailAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDruploadretailAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.druploadretail"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDruploadretailAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillCode Setter
// 单据编号（唯一）
func (r *AlibabaAlihealthDrugKytDruploadretailAPIRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytDruploadretailAPIRequest) GetBillCode() string {
    return r._billCode
}
// BillTime Setter
// 单据生成时间（一般写当前时间）
func (r *AlibabaAlihealthDrugKytDruploadretailAPIRequest) SetBillTime(_billTime string) error {
    r._billTime = _billTime
    r.Set("bill_time", _billTime)
    return nil
}

// BillTime Getter
func (r AlibabaAlihealthDrugKytDruploadretailAPIRequest) GetBillTime() string {
    return r._billTime
}
// BillType Setter
// 单据类型[321,零售出库][322,疫苗接种][116,消费者退货入库]
func (r *AlibabaAlihealthDrugKytDruploadretailAPIRequest) SetBillType(_billType int64) error {
    r._billType = _billType
    r.Set("bill_type", _billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytDruploadretailAPIRequest) GetBillType() int64 {
    return r._billType
}
// PhysicType Setter
// 药品类型[2,特药，3,普药]
func (r *AlibabaAlihealthDrugKytDruploadretailAPIRequest) SetPhysicType(_physicType int64) error {
    r._physicType = _physicType
    r.Set("physic_type", _physicType)
    return nil
}

// PhysicType Getter
func (r AlibabaAlihealthDrugKytDruploadretailAPIRequest) GetPhysicType() int64 {
    return r._physicType
}
// RefUserId Setter
// 码上放心平台企业唯一编码（门店或医疗机构）
func (r *AlibabaAlihealthDrugKytDruploadretailAPIRequest) SetRefUserId(_refUserId string) error {
    r._refUserId = _refUserId
    r.Set("ref_user_id", _refUserId)
    return nil
}

// RefUserId Getter
func (r AlibabaAlihealthDrugKytDruploadretailAPIRequest) GetRefUserId() string {
    return r._refUserId
}
// FromUserId Setter
// 发货企业(可为空)
func (r *AlibabaAlihealthDrugKytDruploadretailAPIRequest) SetFromUserId(_fromUserId string) error {
    r._fromUserId = _fromUserId
    r.Set("from_user_id", _fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugKytDruploadretailAPIRequest) GetFromUserId() string {
    return r._fromUserId
}
// OperIcCode Setter
// 单据提交者(appkey编号)
func (r *AlibabaAlihealthDrugKytDruploadretailAPIRequest) SetOperIcCode(_operIcCode string) error {
    r._operIcCode = _operIcCode
    r.Set("oper_ic_code", _operIcCode)
    return nil
}

// OperIcCode Getter
func (r AlibabaAlihealthDrugKytDruploadretailAPIRequest) GetOperIcCode() string {
    return r._operIcCode
}
// OperIcName Setter
// 单据提交者姓名(可为空)
func (r *AlibabaAlihealthDrugKytDruploadretailAPIRequest) SetOperIcName(_operIcName string) error {
    r._operIcName = _operIcName
    r.Set("oper_ic_name", _operIcName)
    return nil
}

// OperIcName Getter
func (r AlibabaAlihealthDrugKytDruploadretailAPIRequest) GetOperIcName() string {
    return r._operIcName
}
// TraceCodes Setter
// 20位追溯码（多个时用半角逗号分隔）
func (r *AlibabaAlihealthDrugKytDruploadretailAPIRequest) SetTraceCodes(_traceCodes []string) error {
    r._traceCodes = _traceCodes
    r.Set("trace_codes", _traceCodes)
    return nil
}

// TraceCodes Getter
func (r AlibabaAlihealthDrugKytDruploadretailAPIRequest) GetTraceCodes() []string {
    return r._traceCodes
}
// CustomerIdType Setter
// 购买人证件类型【1身份证2护照3 军官证4 医保卡5接种卡6学生证9其它】
func (r *AlibabaAlihealthDrugKytDruploadretailAPIRequest) SetCustomerIdType(_customerIdType string) error {
    r._customerIdType = _customerIdType
    r.Set("customer_id_type", _customerIdType)
    return nil
}

// CustomerIdType Getter
func (r AlibabaAlihealthDrugKytDruploadretailAPIRequest) GetCustomerIdType() string {
    return r._customerIdType
}
// CustomerId Setter
// 购买人证件编号
func (r *AlibabaAlihealthDrugKytDruploadretailAPIRequest) SetCustomerId(_customerId string) error {
    r._customerId = _customerId
    r.Set("customer_id", _customerId)
    return nil
}

// CustomerId Getter
func (r AlibabaAlihealthDrugKytDruploadretailAPIRequest) GetCustomerId() string {
    return r._customerId
}
// UserTel Setter
// 药品购买者电话
func (r *AlibabaAlihealthDrugKytDruploadretailAPIRequest) SetUserTel(_userTel string) error {
    r._userTel = _userTel
    r.Set("user_tel", _userTel)
    return nil
}

// UserTel Getter
func (r AlibabaAlihealthDrugKytDruploadretailAPIRequest) GetUserTel() string {
    return r._userTel
}
// NetworkBillFlag Setter
// 互联网标识 1是 0代表否
func (r *AlibabaAlihealthDrugKytDruploadretailAPIRequest) SetNetworkBillFlag(_networkBillFlag string) error {
    r._networkBillFlag = _networkBillFlag
    r.Set("network_bill_flag", _networkBillFlag)
    return nil
}

// NetworkBillFlag Getter
func (r AlibabaAlihealthDrugKytDruploadretailAPIRequest) GetNetworkBillFlag() string {
    return r._networkBillFlag
}
// MedicDoctor Setter
// 医师名称
func (r *AlibabaAlihealthDrugKytDruploadretailAPIRequest) SetMedicDoctor(_medicDoctor string) error {
    r._medicDoctor = _medicDoctor
    r.Set("medic_doctor", _medicDoctor)
    return nil
}

// MedicDoctor Getter
func (r AlibabaAlihealthDrugKytDruploadretailAPIRequest) GetMedicDoctor() string {
    return r._medicDoctor
}
// MedicDispenser Setter
// 药品发药者
func (r *AlibabaAlihealthDrugKytDruploadretailAPIRequest) SetMedicDispenser(_medicDispenser string) error {
    r._medicDispenser = _medicDispenser
    r.Set("medic_dispenser", _medicDispenser)
    return nil
}

// MedicDispenser Getter
func (r AlibabaAlihealthDrugKytDruploadretailAPIRequest) GetMedicDispenser() string {
    return r._medicDispenser
}
// UserName Setter
// 药品使用者名称
func (r *AlibabaAlihealthDrugKytDruploadretailAPIRequest) SetUserName(_userName string) error {
    r._userName = _userName
    r.Set("user_name", _userName)
    return nil
}

// UserName Getter
func (r AlibabaAlihealthDrugKytDruploadretailAPIRequest) GetUserName() string {
    return r._userName
}
// UserAgent Setter
// 药品使用者代理人
func (r *AlibabaAlihealthDrugKytDruploadretailAPIRequest) SetUserAgent(_userAgent string) error {
    r._userAgent = _userAgent
    r.Set("user_agent", _userAgent)
    return nil
}

// UserAgent Getter
func (r AlibabaAlihealthDrugKytDruploadretailAPIRequest) GetUserAgent() string {
    return r._userAgent
}
