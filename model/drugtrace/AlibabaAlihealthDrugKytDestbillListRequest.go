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
    _refEntId   string
    // 开始时间，格式yyyy-MM-dd
    _beginDate   string
    // 结束时间，格式yyyy-MM-dd
    _endDate   string
    // 单据编号
    _billCode   string
    // 审核状态，1：未审核；2：审核通过；3：审核失败
    _approvalStatus   string
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
func (r *AlibabaAlihealthDrugKytDestbillListRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDestbillListRequest) GetRefEntId() string {
    return r._refEntId
}
// BeginDate Setter
// 开始时间，格式yyyy-MM-dd
func (r *AlibabaAlihealthDrugKytDestbillListRequest) SetBeginDate(_beginDate string) error {
    r._beginDate = _beginDate
    r.Set("begin_date", _beginDate)
    return nil
}

// BeginDate Getter
func (r AlibabaAlihealthDrugKytDestbillListRequest) GetBeginDate() string {
    return r._beginDate
}
// EndDate Setter
// 结束时间，格式yyyy-MM-dd
func (r *AlibabaAlihealthDrugKytDestbillListRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthDrugKytDestbillListRequest) GetEndDate() string {
    return r._endDate
}
// BillCode Setter
// 单据编号
func (r *AlibabaAlihealthDrugKytDestbillListRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytDestbillListRequest) GetBillCode() string {
    return r._billCode
}
// ApprovalStatus Setter
// 审核状态，1：未审核；2：审核通过；3：审核失败
func (r *AlibabaAlihealthDrugKytDestbillListRequest) SetApprovalStatus(_approvalStatus string) error {
    r._approvalStatus = _approvalStatus
    r.Set("approval_status", _approvalStatus)
    return nil
}

// ApprovalStatus Getter
func (r AlibabaAlihealthDrugKytDestbillListRequest) GetApprovalStatus() string {
    return r._approvalStatus
}
