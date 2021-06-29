package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
直调单据查询 API请求
alibaba.alihealth.drug.kyt.destbill.list

为药企提供直调单据查询功能
*/
type AlibabaAlihealthDrugKytDestbillListRequest struct {
    model.Params
    // 企业ID
    refEntId   string
    // 开始时间，格式yyyy-MM-dd
    beginDate   string
    // 结束时间，格式yyyy-MM-dd
    endDate   string
    // 单据编号
    billCode   string
    // 审核状态，1：未审核；2：审核通过；3：审核失败
    approvalStatus   string
}

// 初始化AlibabaAlihealthDrugKytDestbillListRequest对象
func NewAlibabaAlihealthDrugKytDestbillListRequest() *AlibabaAlihealthDrugKytDestbillListRequest{
    return &AlibabaAlihealthDrugKytDestbillListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDestbillListRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.destbill.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDestbillListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytDestbillListRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDestbillListRequest) GetRefEntId() string {
    return r.refEntId
}
// BeginDate Setter
// 开始时间，格式yyyy-MM-dd
func (r *AlibabaAlihealthDrugKytDestbillListRequest) SetBeginDate(beginDate string) error {
    r.beginDate = beginDate
    r.Set("begin_date", beginDate)
    return nil
}

// BeginDate Getter
func (r AlibabaAlihealthDrugKytDestbillListRequest) GetBeginDate() string {
    return r.beginDate
}
// EndDate Setter
// 结束时间，格式yyyy-MM-dd
func (r *AlibabaAlihealthDrugKytDestbillListRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthDrugKytDestbillListRequest) GetEndDate() string {
    return r.endDate
}
// BillCode Setter
// 单据编号
func (r *AlibabaAlihealthDrugKytDestbillListRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytDestbillListRequest) GetBillCode() string {
    return r.billCode
}
// ApprovalStatus Setter
// 审核状态，1：未审核；2：审核通过；3：审核失败
func (r *AlibabaAlihealthDrugKytDestbillListRequest) SetApprovalStatus(approvalStatus string) error {
    r.approvalStatus = approvalStatus
    r.Set("approval_status", approvalStatus)
    return nil
}

// ApprovalStatus Getter
func (r AlibabaAlihealthDrugKytDestbillListRequest) GetApprovalStatus() string {
    return r.approvalStatus
}
