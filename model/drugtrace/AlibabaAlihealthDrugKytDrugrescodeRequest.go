package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询药品码段信息 API请求
alibaba.alihealth.drug.kyt.drugrescode

查询药品码段信息
*/
type AlibabaAlihealthDrugKytDrugrescodeRequest struct {
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

// 初始化AlibabaAlihealthDrugKytDrugrescodeRequest对象
func NewAlibabaAlihealthDrugKytDrugrescodeRequest() *AlibabaAlihealthDrugKytDrugrescodeRequest{
    return &AlibabaAlihealthDrugKytDrugrescodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.drugrescode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetRefEntId() string {
    return r._refEntId
}
// PhysicName Setter
// 药品通用名
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetPhysicName(_physicName string) error {
    r._physicName = _physicName
    r.Set("physic_name", _physicName)
    return nil
}

// PhysicName Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetPhysicName() string {
    return r._physicName
}
// ApprovalLicenceNo Setter
// 批准文号
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetApprovalLicenceNo(_approvalLicenceNo string) error {
    r._approvalLicenceNo = _approvalLicenceNo
    r.Set("approval_licence_no", _approvalLicenceNo)
    return nil
}

// ApprovalLicenceNo Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetApprovalLicenceNo() string {
    return r._approvalLicenceNo
}
// StartDate Setter
// 开始日期
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 结束日期
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetEndDate() string {
    return r._endDate
}
// PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetPageSize() int64 {
    return r._pageSize
}
// Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetPage() int64 {
    return r._page
}
// EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetEntName(_entName string) error {
    r._entName = _entName
    r.Set("ent_name", _entName)
    return nil
}

// EntName Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetEntName() string {
    return r._entName
}
// PackageSpec Setter
// 包装规格
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetPackageSpec(_packageSpec string) error {
    r._packageSpec = _packageSpec
    r.Set("package_spec", _packageSpec)
    return nil
}

// PackageSpec Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetPackageSpec() string {
    return r._packageSpec
}
// PrepnSpec Setter
// 制剂规格
func (r *AlibabaAlihealthDrugKytDrugrescodeRequest) SetPrepnSpec(_prepnSpec string) error {
    r._prepnSpec = _prepnSpec
    r.Set("prepn_spec", _prepnSpec)
    return nil
}

// PrepnSpec Getter
func (r AlibabaAlihealthDrugKytDrugrescodeRequest) GetPrepnSpec() string {
    return r._prepnSpec
}
