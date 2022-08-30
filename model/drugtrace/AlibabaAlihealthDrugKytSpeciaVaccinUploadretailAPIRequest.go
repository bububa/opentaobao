package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest 零售单据上传接口（疫苗） API请求
// alibaba.alihealth.drug.kyt.specia.vaccin.uploadretail
//
// 零售上传单据信息接口
type AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest struct {
	model.Params
	// 20位追溯码（多个时用半角逗号分隔）
	_traceCodes []string
	// 单据编号（唯一）
	_billCode string
	// 单据生成时间（一般写当前时间）
	_billTime string
	// 码上放心平台企业唯一编码（门店或医疗机构）
	_refUserId string
	// 发货企业(可为空)
	_fromUserId string
	// 单据提交者(appkey编号)
	_operIcCode string
	// 单据提交者姓名(可为空)
	_operIcName string
	// 购买人证件类型【1身份证2护照3 军官证4 医保卡5接种卡6学生证9其它】
	_customerIdType string
	// 购买人证件编号
	_customerId string
	// 用药人名称
	_userTel string
	// 互联网标志 1是 0否
	_networkBillFlag string
	// 医师名称
	_medicDoctor string
	// 药品发药人
	_medicDispenser string
	// 药品使用者：
	_userName string
	// 药品使用者代理人
	_userAgent string
	// 单据类型[321,零售出库][322,疫苗接种]
	_billType int64
	// 药品类型[2,特药，3,普药]
	_physicType int64
}

// NewAlibabaAlihealthDrugKytSpeciaVaccinUploadretailRequest 初始化AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest对象
func NewAlibabaAlihealthDrugKytSpeciaVaccinUploadretailRequest() *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest {
	return &AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.specia.vaccin.uploadretail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTraceCodes is TraceCodes Setter
// 20位追溯码（多个时用半角逗号分隔）
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) SetTraceCodes(_traceCodes []string) error {
	r._traceCodes = _traceCodes
	r.Set("trace_codes", _traceCodes)
	return nil
}

// GetTraceCodes TraceCodes Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) GetTraceCodes() []string {
	return r._traceCodes
}

// SetBillCode is BillCode Setter
// 单据编号（唯一）
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillTime is BillTime Setter
// 单据生成时间（一般写当前时间）
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) SetBillTime(_billTime string) error {
	r._billTime = _billTime
	r.Set("bill_time", _billTime)
	return nil
}

// GetBillTime BillTime Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) GetBillTime() string {
	return r._billTime
}

// SetRefUserId is RefUserId Setter
// 码上放心平台企业唯一编码（门店或医疗机构）
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) SetRefUserId(_refUserId string) error {
	r._refUserId = _refUserId
	r.Set("ref_user_id", _refUserId)
	return nil
}

// GetRefUserId RefUserId Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) GetRefUserId() string {
	return r._refUserId
}

// SetFromUserId is FromUserId Setter
// 发货企业(可为空)
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) SetFromUserId(_fromUserId string) error {
	r._fromUserId = _fromUserId
	r.Set("from_user_id", _fromUserId)
	return nil
}

// GetFromUserId FromUserId Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) GetFromUserId() string {
	return r._fromUserId
}

// SetOperIcCode is OperIcCode Setter
// 单据提交者(appkey编号)
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) SetOperIcCode(_operIcCode string) error {
	r._operIcCode = _operIcCode
	r.Set("oper_ic_code", _operIcCode)
	return nil
}

// GetOperIcCode OperIcCode Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) GetOperIcCode() string {
	return r._operIcCode
}

// SetOperIcName is OperIcName Setter
// 单据提交者姓名(可为空)
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) SetOperIcName(_operIcName string) error {
	r._operIcName = _operIcName
	r.Set("oper_ic_name", _operIcName)
	return nil
}

// GetOperIcName OperIcName Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) GetOperIcName() string {
	return r._operIcName
}

// SetCustomerIdType is CustomerIdType Setter
// 购买人证件类型【1身份证2护照3 军官证4 医保卡5接种卡6学生证9其它】
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) SetCustomerIdType(_customerIdType string) error {
	r._customerIdType = _customerIdType
	r.Set("customer_id_type", _customerIdType)
	return nil
}

// GetCustomerIdType CustomerIdType Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) GetCustomerIdType() string {
	return r._customerIdType
}

// SetCustomerId is CustomerId Setter
// 购买人证件编号
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) SetCustomerId(_customerId string) error {
	r._customerId = _customerId
	r.Set("customer_id", _customerId)
	return nil
}

// GetCustomerId CustomerId Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) GetCustomerId() string {
	return r._customerId
}

// SetUserTel is UserTel Setter
// 用药人名称
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) SetUserTel(_userTel string) error {
	r._userTel = _userTel
	r.Set("user_tel", _userTel)
	return nil
}

// GetUserTel UserTel Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) GetUserTel() string {
	return r._userTel
}

// SetNetworkBillFlag is NetworkBillFlag Setter
// 互联网标志 1是 0否
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) SetNetworkBillFlag(_networkBillFlag string) error {
	r._networkBillFlag = _networkBillFlag
	r.Set("network_bill_flag", _networkBillFlag)
	return nil
}

// GetNetworkBillFlag NetworkBillFlag Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) GetNetworkBillFlag() string {
	return r._networkBillFlag
}

// SetMedicDoctor is MedicDoctor Setter
// 医师名称
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) SetMedicDoctor(_medicDoctor string) error {
	r._medicDoctor = _medicDoctor
	r.Set("medic_doctor", _medicDoctor)
	return nil
}

// GetMedicDoctor MedicDoctor Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) GetMedicDoctor() string {
	return r._medicDoctor
}

// SetMedicDispenser is MedicDispenser Setter
// 药品发药人
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) SetMedicDispenser(_medicDispenser string) error {
	r._medicDispenser = _medicDispenser
	r.Set("medic_dispenser", _medicDispenser)
	return nil
}

// GetMedicDispenser MedicDispenser Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) GetMedicDispenser() string {
	return r._medicDispenser
}

// SetUserName is UserName Setter
// 药品使用者：
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) SetUserName(_userName string) error {
	r._userName = _userName
	r.Set("user_name", _userName)
	return nil
}

// GetUserName UserName Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) GetUserName() string {
	return r._userName
}

// SetUserAgent is UserAgent Setter
// 药品使用者代理人
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) SetUserAgent(_userAgent string) error {
	r._userAgent = _userAgent
	r.Set("user_agent", _userAgent)
	return nil
}

// GetUserAgent UserAgent Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) GetUserAgent() string {
	return r._userAgent
}

// SetBillType is BillType Setter
// 单据类型[321,零售出库][322,疫苗接种]
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) SetBillType(_billType int64) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) GetBillType() int64 {
	return r._billType
}

// SetPhysicType is PhysicType Setter
// 药品类型[2,特药，3,普药]
func (r *AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) SetPhysicType(_physicType int64) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// GetPhysicType PhysicType Getter
func (r AlibabaAlihealthDrugKytSpeciaVaccinUploadretailAPIRequest) GetPhysicType() int64 {
	return r._physicType
}
