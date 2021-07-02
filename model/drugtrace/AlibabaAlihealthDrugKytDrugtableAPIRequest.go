package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrugtableAPIRequest 查询药品目录信息 API请求
// alibaba.alihealth.drug.kyt.drugtable
//
// 查询药品目录信息
type AlibabaAlihealthDrugKytDrugtableAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 药品通用名
	_physicName string
	// 批准文号
	_approvalLicenceNo string
	// 开始日期
	_startDate string
	// 结束日期
	_endDate string
	// 页大小
	_pageSize int64
	// 页码
	_page int64
	// 企业名称
	_entName string
	// 包装规格
	_packageSpec string
	// 制剂规格
	_prepnSpec string
}

// NewAlibabaAlihealthDrugKytDrugtableRequest 初始化AlibabaAlihealthDrugKytDrugtableAPIRequest对象
func NewAlibabaAlihealthDrugKytDrugtableRequest() *AlibabaAlihealthDrugKytDrugtableAPIRequest {
	return &AlibabaAlihealthDrugKytDrugtableAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrugtableAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.drugtable"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrugtableAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytDrugtableAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytDrugtableAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetPhysicName is PhysicName Setter
// 药品通用名
func (r *AlibabaAlihealthDrugKytDrugtableAPIRequest) SetPhysicName(_physicName string) error {
	r._physicName = _physicName
	r.Set("physic_name", _physicName)
	return nil
}

// GetPhysicName PhysicName Getter
func (r AlibabaAlihealthDrugKytDrugtableAPIRequest) GetPhysicName() string {
	return r._physicName
}

// SetApprovalLicenceNo is ApprovalLicenceNo Setter
// 批准文号
func (r *AlibabaAlihealthDrugKytDrugtableAPIRequest) SetApprovalLicenceNo(_approvalLicenceNo string) error {
	r._approvalLicenceNo = _approvalLicenceNo
	r.Set("approval_licence_no", _approvalLicenceNo)
	return nil
}

// GetApprovalLicenceNo ApprovalLicenceNo Getter
func (r AlibabaAlihealthDrugKytDrugtableAPIRequest) GetApprovalLicenceNo() string {
	return r._approvalLicenceNo
}

// SetStartDate is StartDate Setter
// 开始日期
func (r *AlibabaAlihealthDrugKytDrugtableAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r AlibabaAlihealthDrugKytDrugtableAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 结束日期
func (r *AlibabaAlihealthDrugKytDrugtableAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaAlihealthDrugKytDrugtableAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytDrugtableAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthDrugKytDrugtableAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPage is Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytDrugtableAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaAlihealthDrugKytDrugtableAPIRequest) GetPage() int64 {
	return r._page
}

// SetEntName is EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytDrugtableAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugKytDrugtableAPIRequest) GetEntName() string {
	return r._entName
}

// SetPackageSpec is PackageSpec Setter
// 包装规格
func (r *AlibabaAlihealthDrugKytDrugtableAPIRequest) SetPackageSpec(_packageSpec string) error {
	r._packageSpec = _packageSpec
	r.Set("package_spec", _packageSpec)
	return nil
}

// GetPackageSpec PackageSpec Getter
func (r AlibabaAlihealthDrugKytDrugtableAPIRequest) GetPackageSpec() string {
	return r._packageSpec
}

// SetPrepnSpec is PrepnSpec Setter
// 制剂规格
func (r *AlibabaAlihealthDrugKytDrugtableAPIRequest) SetPrepnSpec(_prepnSpec string) error {
	r._prepnSpec = _prepnSpec
	r.Set("prepn_spec", _prepnSpec)
	return nil
}

// GetPrepnSpec PrepnSpec Getter
func (r AlibabaAlihealthDrugKytDrugtableAPIRequest) GetPrepnSpec() string {
	return r._prepnSpec
}
