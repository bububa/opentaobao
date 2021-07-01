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
type AlibabaAlihealthDrugKytUploadretailAPIRequest struct {
    model.Params
    // 单据编号（唯一）
    _billCode   string
    // 单据生成时间
    _billTime   string
    // 单据类型[321,零售出库][322,疫苗接种]
    _billType   int64
    // 药品类型[3,普药]
    _physicType   int64
    // 码上放心平台企业编码（门店）
    _refUserId   string
    // 发货企业(可为空)
    _fromUserId   string
    // 单据提交者(appkey编号)
    _operIcCode   string
    // 单据提交者姓名(可为空)
    _operIcName   string
    // 请求类型[暂定都写2]
    _clientType   string
    // 追溯码[多个时用逗号分开]
    _traceCodes   []string
    // 患者电话
    _userTel   string
    // 互联网标识
    _networkBillFlag   string
    // 医师名单
    _medicDoctor   string
    // 药品分发者
    _medicDispenser   string
    // 患者名称
    _userName   string
    // 患者代理领药人
    _userAgent   string
    // 购买人证件类型【1身份证2护照3 军官证4 医保卡5接种卡6学生证9其它】
    _customerIdType   string
    // 购买人证件编号
    _customerId   string
}

// 初始化AlibabaAlihealthDrugKytUploadretailAPIRequest对象
func NewAlibabaAlihealthDrugKytUploadretailRequest() *AlibabaAlihealthDrugKytUploadretailAPIRequest{
    return &AlibabaAlihealthDrugKytUploadretailAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.uploadretail"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillCode Setter
// 单据编号（唯一）
func (r *AlibabaAlihealthDrugKytUploadretailAPIRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetBillCode() string {
    return r._billCode
}
// BillTime Setter
// 单据生成时间
func (r *AlibabaAlihealthDrugKytUploadretailAPIRequest) SetBillTime(_billTime string) error {
    r._billTime = _billTime
    r.Set("bill_time", _billTime)
    return nil
}

// BillTime Getter
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetBillTime() string {
    return r._billTime
}
// BillType Setter
// 单据类型[321,零售出库][322,疫苗接种]
func (r *AlibabaAlihealthDrugKytUploadretailAPIRequest) SetBillType(_billType int64) error {
    r._billType = _billType
    r.Set("bill_type", _billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetBillType() int64 {
    return r._billType
}
// PhysicType Setter
// 药品类型[3,普药]
func (r *AlibabaAlihealthDrugKytUploadretailAPIRequest) SetPhysicType(_physicType int64) error {
    r._physicType = _physicType
    r.Set("physic_type", _physicType)
    return nil
}

// PhysicType Getter
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetPhysicType() int64 {
    return r._physicType
}
// RefUserId Setter
// 码上放心平台企业编码（门店）
func (r *AlibabaAlihealthDrugKytUploadretailAPIRequest) SetRefUserId(_refUserId string) error {
    r._refUserId = _refUserId
    r.Set("ref_user_id", _refUserId)
    return nil
}

// RefUserId Getter
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetRefUserId() string {
    return r._refUserId
}
// FromUserId Setter
// 发货企业(可为空)
func (r *AlibabaAlihealthDrugKytUploadretailAPIRequest) SetFromUserId(_fromUserId string) error {
    r._fromUserId = _fromUserId
    r.Set("from_user_id", _fromUserId)
    return nil
}

// FromUserId Getter
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetFromUserId() string {
    return r._fromUserId
}
// OperIcCode Setter
// 单据提交者(appkey编号)
func (r *AlibabaAlihealthDrugKytUploadretailAPIRequest) SetOperIcCode(_operIcCode string) error {
    r._operIcCode = _operIcCode
    r.Set("oper_ic_code", _operIcCode)
    return nil
}

// OperIcCode Getter
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetOperIcCode() string {
    return r._operIcCode
}
// OperIcName Setter
// 单据提交者姓名(可为空)
func (r *AlibabaAlihealthDrugKytUploadretailAPIRequest) SetOperIcName(_operIcName string) error {
    r._operIcName = _operIcName
    r.Set("oper_ic_name", _operIcName)
    return nil
}

// OperIcName Getter
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetOperIcName() string {
    return r._operIcName
}
// ClientType Setter
// 请求类型[暂定都写2]
func (r *AlibabaAlihealthDrugKytUploadretailAPIRequest) SetClientType(_clientType string) error {
    r._clientType = _clientType
    r.Set("client_type", _clientType)
    return nil
}

// ClientType Getter
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetClientType() string {
    return r._clientType
}
// TraceCodes Setter
// 追溯码[多个时用逗号分开]
func (r *AlibabaAlihealthDrugKytUploadretailAPIRequest) SetTraceCodes(_traceCodes []string) error {
    r._traceCodes = _traceCodes
    r.Set("trace_codes", _traceCodes)
    return nil
}

// TraceCodes Getter
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetTraceCodes() []string {
    return r._traceCodes
}
// UserTel Setter
// 患者电话
func (r *AlibabaAlihealthDrugKytUploadretailAPIRequest) SetUserTel(_userTel string) error {
    r._userTel = _userTel
    r.Set("user_tel", _userTel)
    return nil
}

// UserTel Getter
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetUserTel() string {
    return r._userTel
}
// NetworkBillFlag Setter
// 互联网标识
func (r *AlibabaAlihealthDrugKytUploadretailAPIRequest) SetNetworkBillFlag(_networkBillFlag string) error {
    r._networkBillFlag = _networkBillFlag
    r.Set("network_bill_flag", _networkBillFlag)
    return nil
}

// NetworkBillFlag Getter
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetNetworkBillFlag() string {
    return r._networkBillFlag
}
// MedicDoctor Setter
// 医师名单
func (r *AlibabaAlihealthDrugKytUploadretailAPIRequest) SetMedicDoctor(_medicDoctor string) error {
    r._medicDoctor = _medicDoctor
    r.Set("medic_doctor", _medicDoctor)
    return nil
}

// MedicDoctor Getter
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetMedicDoctor() string {
    return r._medicDoctor
}
// MedicDispenser Setter
// 药品分发者
func (r *AlibabaAlihealthDrugKytUploadretailAPIRequest) SetMedicDispenser(_medicDispenser string) error {
    r._medicDispenser = _medicDispenser
    r.Set("medic_dispenser", _medicDispenser)
    return nil
}

// MedicDispenser Getter
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetMedicDispenser() string {
    return r._medicDispenser
}
// UserName Setter
// 患者名称
func (r *AlibabaAlihealthDrugKytUploadretailAPIRequest) SetUserName(_userName string) error {
    r._userName = _userName
    r.Set("user_name", _userName)
    return nil
}

// UserName Getter
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetUserName() string {
    return r._userName
}
// UserAgent Setter
// 患者代理领药人
func (r *AlibabaAlihealthDrugKytUploadretailAPIRequest) SetUserAgent(_userAgent string) error {
    r._userAgent = _userAgent
    r.Set("user_agent", _userAgent)
    return nil
}

// UserAgent Getter
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetUserAgent() string {
    return r._userAgent
}
// CustomerIdType Setter
// 购买人证件类型【1身份证2护照3 军官证4 医保卡5接种卡6学生证9其它】
func (r *AlibabaAlihealthDrugKytUploadretailAPIRequest) SetCustomerIdType(_customerIdType string) error {
    r._customerIdType = _customerIdType
    r.Set("customer_id_type", _customerIdType)
    return nil
}

// CustomerIdType Getter
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetCustomerIdType() string {
    return r._customerIdType
}
// CustomerId Setter
// 购买人证件编号
func (r *AlibabaAlihealthDrugKytUploadretailAPIRequest) SetCustomerId(_customerId string) error {
    r._customerId = _customerId
    r.Set("customer_id", _customerId)
    return nil
}

// CustomerId Getter
func (r AlibabaAlihealthDrugKytUploadretailAPIRequest) GetCustomerId() string {
    return r._customerId
}
