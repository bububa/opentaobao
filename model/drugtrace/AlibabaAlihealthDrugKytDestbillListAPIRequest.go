package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytdestbilllistAPIRequest 直调单据查询 API请求
// alibaba.alihealth.drug.kyt.destbill.list
//
// 为药企提供直调单据查询功能
type AlibabaalihealthdrugkytdestbilllistAPIRequest struct {
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

// NewAlibabaalihealthdrugkytdestbilllistRequest 初始化AlibabaalihealthdrugkytdestbilllistAPIRequest对象
func NewAlibabaalihealthdrugkytdestbilllistRequest() *AlibabaalihealthdrugkytdestbilllistAPIRequest {
	return &AlibabaalihealthdrugkytdestbilllistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytdestbilllistAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.destbill.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytdestbilllistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytdestbilllistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaalihealthdrugkytdestbilllistAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytdestbilllistAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBeginDate is BeginDate Setter
// 开始时间，格式yyyy-MM-dd
func (r *AlibabaalihealthdrugkytdestbilllistAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// GetBeginDate BeginDate Getter
func (r AlibabaalihealthdrugkytdestbilllistAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// SetEndDate is EndDate Setter
// 结束时间，格式yyyy-MM-dd
func (r *AlibabaalihealthdrugkytdestbilllistAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaalihealthdrugkytdestbilllistAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetBillCode is BillCode Setter
// 单据编号
func (r *AlibabaalihealthdrugkytdestbilllistAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugkytdestbilllistAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetApprovalStatus is ApprovalStatus Setter
// 审核状态，1：未审核；2：审核通过；3：审核失败
func (r *AlibabaalihealthdrugkytdestbilllistAPIRequest) SetApprovalStatus(_approvalStatus string) error {
	r._approvalStatus = _approvalStatus
	r.Set("approval_status", _approvalStatus)
	return nil
}

// GetApprovalStatus ApprovalStatus Getter
func (r AlibabaalihealthdrugkytdestbilllistAPIRequest) GetApprovalStatus() string {
	return r._approvalStatus
}
