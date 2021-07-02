package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytYyUploadretailAPIRequest 医院上传出库信息 API请求
// alibaba.alihealth.drug.kyt.yy.uploadretail
//
// 医院上传出库信息接口
type AlibabaAlihealthDrugKytYyUploadretailAPIRequest struct {
	model.Params
	// 单据编号（唯一）
	_billCode string
	// 单据生成时间（一般写当前时间）
	_billTime string
	// 单据类型[321,零售出库][322,疫苗接种]
	_billType int64
	// 药品类型[2,特药，3,普药]
	_physicType int64
	// 码上放心平台企业唯一编码（门店或医疗机构）
	_refUserId string
	// 发货企业(可为空)
	_fromUserId string
	// 单据提交者(appkey编号)
	_operIcCode string
	// 药房名称(可为空)
	_drugstoreName string
	// 20位追溯码（多个时用半角逗号分隔）
	_traceCodes []string
	// 购买人证件类型【1身份证2护照3 军官证4 医保卡5接种卡6学生证9其它】
	_customerIdType string
	// 购买人证件编号
	_customerId string
	// 购买人电话：
	_userTel string
	// 互联网标识 1是 0代表否
	_networkBillFlag string
	// 开药医师名称
	_medicDoctor string
	// 药品使用分药者
	_medicDispenser string
	// 药品使用者名称
	_userName string
	// 药品使用者代理人
	_userAgent string
}

// NewAlibabaAlihealthDrugKytYyUploadretailRequest 初始化AlibabaAlihealthDrugKytYyUploadretailAPIRequest对象
func NewAlibabaAlihealthDrugKytYyUploadretailRequest() *AlibabaAlihealthDrugKytYyUploadretailAPIRequest {
	return &AlibabaAlihealthDrugKytYyUploadretailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytYyUploadretailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.yy.uploadretail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytYyUploadretailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBillCode is BillCode Setter
// 单据编号（唯一）
func (r *AlibabaAlihealthDrugKytYyUploadretailAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugKytYyUploadretailAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetBillTime is BillTime Setter
// 单据生成时间（一般写当前时间）
func (r *AlibabaAlihealthDrugKytYyUploadretailAPIRequest) SetBillTime(_billTime string) error {
	r._billTime = _billTime
	r.Set("bill_time", _billTime)
	return nil
}

// GetBillTime BillTime Getter
func (r AlibabaAlihealthDrugKytYyUploadretailAPIRequest) GetBillTime() string {
	return r._billTime
}

// SetBillType is BillType Setter
// 单据类型[321,零售出库][322,疫苗接种]
func (r *AlibabaAlihealthDrugKytYyUploadretailAPIRequest) SetBillType(_billType int64) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaAlihealthDrugKytYyUploadretailAPIRequest) GetBillType() int64 {
	return r._billType
}

// SetPhysicType is PhysicType Setter
// 药品类型[2,特药，3,普药]
func (r *AlibabaAlihealthDrugKytYyUploadretailAPIRequest) SetPhysicType(_physicType int64) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// GetPhysicType PhysicType Getter
func (r AlibabaAlihealthDrugKytYyUploadretailAPIRequest) GetPhysicType() int64 {
	return r._physicType
}

// SetRefUserId is RefUserId Setter
// 码上放心平台企业唯一编码（门店或医疗机构）
func (r *AlibabaAlihealthDrugKytYyUploadretailAPIRequest) SetRefUserId(_refUserId string) error {
	r._refUserId = _refUserId
	r.Set("ref_user_id", _refUserId)
	return nil
}

// GetRefUserId RefUserId Getter
func (r AlibabaAlihealthDrugKytYyUploadretailAPIRequest) GetRefUserId() string {
	return r._refUserId
}

// SetFromUserId is FromUserId Setter
// 发货企业(可为空)
func (r *AlibabaAlihealthDrugKytYyUploadretailAPIRequest) SetFromUserId(_fromUserId string) error {
	r._fromUserId = _fromUserId
	r.Set("from_user_id", _fromUserId)
	return nil
}

// GetFromUserId FromUserId Getter
func (r AlibabaAlihealthDrugKytYyUploadretailAPIRequest) GetFromUserId() string {
	return r._fromUserId
}

// SetOperIcCode is OperIcCode Setter
// 单据提交者(appkey编号)
func (r *AlibabaAlihealthDrugKytYyUploadretailAPIRequest) SetOperIcCode(_operIcCode string) error {
	r._operIcCode = _operIcCode
	r.Set("oper_ic_code", _operIcCode)
	return nil
}

// GetOperIcCode OperIcCode Getter
func (r AlibabaAlihealthDrugKytYyUploadretailAPIRequest) GetOperIcCode() string {
	return r._operIcCode
}

// SetDrugstoreName is DrugstoreName Setter
// 药房名称(可为空)
func (r *AlibabaAlihealthDrugKytYyUploadretailAPIRequest) SetDrugstoreName(_drugstoreName string) error {
	r._drugstoreName = _drugstoreName
	r.Set("drugstore_name", _drugstoreName)
	return nil
}

// GetDrugstoreName DrugstoreName Getter
func (r AlibabaAlihealthDrugKytYyUploadretailAPIRequest) GetDrugstoreName() string {
	return r._drugstoreName
}

// SetTraceCodes is TraceCodes Setter
// 20位追溯码（多个时用半角逗号分隔）
func (r *AlibabaAlihealthDrugKytYyUploadretailAPIRequest) SetTraceCodes(_traceCodes []string) error {
	r._traceCodes = _traceCodes
	r.Set("trace_codes", _traceCodes)
	return nil
}

// GetTraceCodes TraceCodes Getter
func (r AlibabaAlihealthDrugKytYyUploadretailAPIRequest) GetTraceCodes() []string {
	return r._traceCodes
}

// SetCustomerIdType is CustomerIdType Setter
// 购买人证件类型【1身份证2护照3 军官证4 医保卡5接种卡6学生证9其它】
func (r *AlibabaAlihealthDrugKytYyUploadretailAPIRequest) SetCustomerIdType(_customerIdType string) error {
	r._customerIdType = _customerIdType
	r.Set("customer_id_type", _customerIdType)
	return nil
}

// GetCustomerIdType CustomerIdType Getter
func (r AlibabaAlihealthDrugKytYyUploadretailAPIRequest) GetCustomerIdType() string {
	return r._customerIdType
}

// SetCustomerId is CustomerId Setter
// 购买人证件编号
func (r *AlibabaAlihealthDrugKytYyUploadretailAPIRequest) SetCustomerId(_customerId string) error {
	r._customerId = _customerId
	r.Set("customer_id", _customerId)
	return nil
}

// GetCustomerId CustomerId Getter
func (r AlibabaAlihealthDrugKytYyUploadretailAPIRequest) GetCustomerId() string {
	return r._customerId
}

// SetUserTel is UserTel Setter
// 购买人电话：
func (r *AlibabaAlihealthDrugKytYyUploadretailAPIRequest) SetUserTel(_userTel string) error {
	r._userTel = _userTel
	r.Set("user_tel", _userTel)
	return nil
}

// GetUserTel UserTel Getter
func (r AlibabaAlihealthDrugKytYyUploadretailAPIRequest) GetUserTel() string {
	return r._userTel
}

// SetNetworkBillFlag is NetworkBillFlag Setter
// 互联网标识 1是 0代表否
func (r *AlibabaAlihealthDrugKytYyUploadretailAPIRequest) SetNetworkBillFlag(_networkBillFlag string) error {
	r._networkBillFlag = _networkBillFlag
	r.Set("network_bill_flag", _networkBillFlag)
	return nil
}

// GetNetworkBillFlag NetworkBillFlag Getter
func (r AlibabaAlihealthDrugKytYyUploadretailAPIRequest) GetNetworkBillFlag() string {
	return r._networkBillFlag
}

// SetMedicDoctor is MedicDoctor Setter
// 开药医师名称
func (r *AlibabaAlihealthDrugKytYyUploadretailAPIRequest) SetMedicDoctor(_medicDoctor string) error {
	r._medicDoctor = _medicDoctor
	r.Set("medic_doctor", _medicDoctor)
	return nil
}

// GetMedicDoctor MedicDoctor Getter
func (r AlibabaAlihealthDrugKytYyUploadretailAPIRequest) GetMedicDoctor() string {
	return r._medicDoctor
}

// SetMedicDispenser is MedicDispenser Setter
// 药品使用分药者
func (r *AlibabaAlihealthDrugKytYyUploadretailAPIRequest) SetMedicDispenser(_medicDispenser string) error {
	r._medicDispenser = _medicDispenser
	r.Set("medic_dispenser", _medicDispenser)
	return nil
}

// GetMedicDispenser MedicDispenser Getter
func (r AlibabaAlihealthDrugKytYyUploadretailAPIRequest) GetMedicDispenser() string {
	return r._medicDispenser
}

// SetUserName is UserName Setter
// 药品使用者名称
func (r *AlibabaAlihealthDrugKytYyUploadretailAPIRequest) SetUserName(_userName string) error {
	r._userName = _userName
	r.Set("user_name", _userName)
	return nil
}

// GetUserName UserName Getter
func (r AlibabaAlihealthDrugKytYyUploadretailAPIRequest) GetUserName() string {
	return r._userName
}

// SetUserAgent is UserAgent Setter
// 药品使用者代理人
func (r *AlibabaAlihealthDrugKytYyUploadretailAPIRequest) SetUserAgent(_userAgent string) error {
	r._userAgent = _userAgent
	r.Set("user_agent", _userAgent)
	return nil
}

// GetUserAgent UserAgent Getter
func (r AlibabaAlihealthDrugKytYyUploadretailAPIRequest) GetUserAgent() string {
	return r._userAgent
}
