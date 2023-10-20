package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest 查询药品码段信息 API请求
// alibaba.alihealth.drug.kyt.wes.drugrescode
//
// 查询药品码段信息
type AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 服务时间
	_licenseToken string
	// 药品通用名
	_physicName string
	// 批准文号
	_approvalLicenceNo string
	// 开始日期
	_startDate string
	// 结束日期
	_endDate string
	// 企业名称
	_entName string
	// 包装规格
	_packageSpec string
	// 制剂规格
	_prepnSpec string
	// 页大小
	_pageSize int64
	// 页码
	_page int64
}

// NewAlibabaAlihealthDrugKytWesDrugrescodeRequest 初始化AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest对象
func NewAlibabaAlihealthDrugKytWesDrugrescodeRequest() *AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest {
	return &AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest{
		Params: model.NewParams(11),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) Reset() {
	r._refEntId = ""
	r._licenseToken = ""
	r._physicName = ""
	r._approvalLicenceNo = ""
	r._startDate = ""
	r._endDate = ""
	r._entName = ""
	r._packageSpec = ""
	r._prepnSpec = ""
	r._pageSize = 0
	r._page = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.drugrescode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// 服务时间
func (r *AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetPhysicName is PhysicName Setter
// 药品通用名
func (r *AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) SetPhysicName(_physicName string) error {
	r._physicName = _physicName
	r.Set("physic_name", _physicName)
	return nil
}

// GetPhysicName PhysicName Getter
func (r AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) GetPhysicName() string {
	return r._physicName
}

// SetApprovalLicenceNo is ApprovalLicenceNo Setter
// 批准文号
func (r *AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) SetApprovalLicenceNo(_approvalLicenceNo string) error {
	r._approvalLicenceNo = _approvalLicenceNo
	r.Set("approval_licence_no", _approvalLicenceNo)
	return nil
}

// GetApprovalLicenceNo ApprovalLicenceNo Getter
func (r AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) GetApprovalLicenceNo() string {
	return r._approvalLicenceNo
}

// SetStartDate is StartDate Setter
// 开始日期
func (r *AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 结束日期
func (r *AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetEntName is EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) GetEntName() string {
	return r._entName
}

// SetPackageSpec is PackageSpec Setter
// 包装规格
func (r *AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) SetPackageSpec(_packageSpec string) error {
	r._packageSpec = _packageSpec
	r.Set("package_spec", _packageSpec)
	return nil
}

// GetPackageSpec PackageSpec Getter
func (r AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) GetPackageSpec() string {
	return r._packageSpec
}

// SetPrepnSpec is PrepnSpec Setter
// 制剂规格
func (r *AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) SetPrepnSpec(_prepnSpec string) error {
	r._prepnSpec = _prepnSpec
	r.Set("prepn_spec", _prepnSpec)
	return nil
}

// GetPrepnSpec PrepnSpec Getter
func (r AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) GetPrepnSpec() string {
	return r._prepnSpec
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPage is Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) GetPage() int64 {
	return r._page
}

var poolAlibabaAlihealthDrugKytWesDrugrescodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytWesDrugrescodeRequest()
	},
}

// GetAlibabaAlihealthDrugKytWesDrugrescodeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest
func GetAlibabaAlihealthDrugKytWesDrugrescodeAPIRequest() *AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest {
	return poolAlibabaAlihealthDrugKytWesDrugrescodeAPIRequest.Get().(*AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytWesDrugrescodeAPIRequest 将 AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytWesDrugrescodeAPIRequest(v *AlibabaAlihealthDrugKytWesDrugrescodeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytWesDrugrescodeAPIRequest.Put(v)
}
