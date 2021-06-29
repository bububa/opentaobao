package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询药品目录信息 API请求
alibaba.alihealth.drug.kyt.drugtable

查询药品目录信息
*/
type AlibabaAlihealthDrugKytDrugtableRequest struct {
    model.Params
    // 企业ID
    _refEntId   string
    // 药品通用名
    _physicName   string
    // 批准文号
    _approvalLicenceNo   string
    // 开始日期
    _startDate   string
    // 结束日期
    _endDate   string
    // 页大小
    _pageSize   int64
    // 页码
    _page   int64
    // 企业名称
    _entName   string
    // 包装规格
    _packageSpec   string
    // 制剂规格
    _prepnSpec   string
}

// 初始化AlibabaAlihealthDrugKytDrugtableRequest对象
func NewAlibabaAlihealthDrugKytDrugtableRequest() *AlibabaAlihealthDrugKytDrugtableRequest{
    return &AlibabaAlihealthDrugKytDrugtableRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrugtableRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.drugtable"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrugtableRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrugtableRequest) GetRefEntId() string {
    return r._refEntId
}
// PhysicName Setter
// 药品通用名
func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetPhysicName(_physicName string) error {
    r._physicName = _physicName
    r.Set("physic_name", _physicName)
    return nil
}

// PhysicName Getter
func (r AlibabaAlihealthDrugKytDrugtableRequest) GetPhysicName() string {
    return r._physicName
}
// ApprovalLicenceNo Setter
// 批准文号
func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetApprovalLicenceNo(_approvalLicenceNo string) error {
    r._approvalLicenceNo = _approvalLicenceNo
    r.Set("approval_licence_no", _approvalLicenceNo)
    return nil
}

// ApprovalLicenceNo Getter
func (r AlibabaAlihealthDrugKytDrugtableRequest) GetApprovalLicenceNo() string {
    return r._approvalLicenceNo
}
// StartDate Setter
// 开始日期
func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r AlibabaAlihealthDrugKytDrugtableRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 结束日期
func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthDrugKytDrugtableRequest) GetEndDate() string {
    return r._endDate
}
// PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugKytDrugtableRequest) GetPageSize() int64 {
    return r._pageSize
}
// Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugKytDrugtableRequest) GetPage() int64 {
    return r._page
}
// EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetEntName(_entName string) error {
    r._entName = _entName
    r.Set("ent_name", _entName)
    return nil
}

// EntName Getter
func (r AlibabaAlihealthDrugKytDrugtableRequest) GetEntName() string {
    return r._entName
}
// PackageSpec Setter
// 包装规格
func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetPackageSpec(_packageSpec string) error {
    r._packageSpec = _packageSpec
    r.Set("package_spec", _packageSpec)
    return nil
}

// PackageSpec Getter
func (r AlibabaAlihealthDrugKytDrugtableRequest) GetPackageSpec() string {
    return r._packageSpec
}
// PrepnSpec Setter
// 制剂规格
func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetPrepnSpec(_prepnSpec string) error {
    r._prepnSpec = _prepnSpec
    r.Set("prepn_spec", _prepnSpec)
    return nil
}

// PrepnSpec Getter
func (r AlibabaAlihealthDrugKytDrugtableRequest) GetPrepnSpec() string {
    return r._prepnSpec
}
