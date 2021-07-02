package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDestbillListAPIRequest 直调单据查询 API请求
// alibaba.alihealth.drug.kyt.destbill.list
//
// 为药企提供直调单据查询功能
type AlibabaAlihealthDrugKytDestbillListAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 开始时间，格式yyyy-MM-dd
	_beginDate string
	// 结束时间，格式yyyy-MM-dd
	_endDate string
	// 单据编号
	_billCode string
	// 审核状态，1：未审核；2：审核通过；3：审核失败
	_approvalStatus string
}

// NewAlibabaAlihealthDrugKytDestbillListRequest 初始化AlibabaAlihealthDrugKytDestbillListAPIRequest对象
func NewAlibabaAlihealthDrugKytDestbillListRequest() *AlibabaAlihealthDrugKytDestbillListAPIRequest {
	return &AlibabaAlihealthDrugKytDestbillListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDestbillListAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.destbill.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDestbillListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytDestbillListAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugKytDestbillListAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// Set is BeginDate Setter
// 开始时间，格式yyyy-MM-dd
func (r *AlibabaAlihealthDrugKytDestbillListAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// Get BeginDate Getter
func (r AlibabaAlihealthDrugKytDestbillListAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// Set is EndDate Setter
// 结束时间，格式yyyy-MM-dd
func (r *AlibabaAlihealthDrugKytDestbillListAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r AlibabaAlihealthDrugKytDestbillListAPIRequest) GetEndDate() string {
	return r._endDate
}

// Set is BillCode Setter
// 单据编号
func (r *AlibabaAlihealthDrugKytDestbillListAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// Get BillCode Getter
func (r AlibabaAlihealthDrugKytDestbillListAPIRequest) GetBillCode() string {
	return r._billCode
}

// Set is ApprovalStatus Setter
// 审核状态，1：未审核；2：审核通过；3：审核失败
func (r *AlibabaAlihealthDrugKytDestbillListAPIRequest) SetApprovalStatus(_approvalStatus string) error {
	r._approvalStatus = _approvalStatus
	r.Set("approval_status", _approvalStatus)
	return nil
}

// Get ApprovalStatus Getter
func (r AlibabaAlihealthDrugKytDestbillListAPIRequest) GetApprovalStatus() string {
	return r._approvalStatus
}
