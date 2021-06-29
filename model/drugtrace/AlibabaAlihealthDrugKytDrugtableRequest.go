package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询药品目录信息 APIRequest
alibaba.alihealth.drug.kyt.drugtable

查询药品目录信息
*/
type AlibabaAlihealthDrugKytDrugtableRequest struct {
    model.Params

    // 企业ID
    refEntId   string 

    // 药品通用名
    physicName   string 

    // 批准文号
    approvalLicenceNo   string 

    // 开始日期
    startDate   string 

    // 结束日期
    endDate   string 

    // 页大小
    pageSize   int64 

    // 页码
    page   int64 

    // 企业名称
    entName   string 

    // 包装规格
    packageSpec   string 

    // 制剂规格
    prepnSpec   string 

}

func NewAlibabaAlihealthDrugKytDrugtableRequest() *AlibabaAlihealthDrugKytDrugtableRequest{
    return &AlibabaAlihealthDrugKytDrugtableRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytDrugtableRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.drugtable"
}

func (r AlibabaAlihealthDrugKytDrugtableRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrugtableRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetPhysicName(physicName string) error {
    r.physicName = physicName
    r.Set("physic_name", physicName)
    return nil
}

func (r AlibabaAlihealthDrugKytDrugtableRequest) GetPhysicName() string {
    return r.physicName
}

func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetApprovalLicenceNo(approvalLicenceNo string) error {
    r.approvalLicenceNo = approvalLicenceNo
    r.Set("approval_licence_no", approvalLicenceNo)
    return nil
}

func (r AlibabaAlihealthDrugKytDrugtableRequest) GetApprovalLicenceNo() string {
    return r.approvalLicenceNo
}

func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r AlibabaAlihealthDrugKytDrugtableRequest) GetStartDate() string {
    return r.startDate
}

func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r AlibabaAlihealthDrugKytDrugtableRequest) GetEndDate() string {
    return r.endDate
}

func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlibabaAlihealthDrugKytDrugtableRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlibabaAlihealthDrugKytDrugtableRequest) GetPage() int64 {
    return r.page
}

func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetEntName(entName string) error {
    r.entName = entName
    r.Set("ent_name", entName)
    return nil
}

func (r AlibabaAlihealthDrugKytDrugtableRequest) GetEntName() string {
    return r.entName
}

func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetPackageSpec(packageSpec string) error {
    r.packageSpec = packageSpec
    r.Set("package_spec", packageSpec)
    return nil
}

func (r AlibabaAlihealthDrugKytDrugtableRequest) GetPackageSpec() string {
    return r.packageSpec
}

func (r *AlibabaAlihealthDrugKytDrugtableRequest) SetPrepnSpec(prepnSpec string) error {
    r.prepnSpec = prepnSpec
    r.Set("prepn_spec", prepnSpec)
    return nil
}

func (r AlibabaAlihealthDrugKytDrugtableRequest) GetPrepnSpec() string {
    return r.prepnSpec
}

