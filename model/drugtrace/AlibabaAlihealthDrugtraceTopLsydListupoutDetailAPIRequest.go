package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest
上游出库单单据明细查询 API请求
alibaba.alihealth.drugtrace.top.lsyd.listupout.detail

查询上游出库单明细(带追溯码信息) */
type AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
	// 单据编码
	_billCode string
	// 发货企业renEntId
	_fromRefUserId string
	// 收货企业refEntId
	_toRefUserId string
}

// NewAlibabaAlihealthDrugtraceTopLsydListupoutDetailRequest 初始化AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopLsydListupoutDetailRequest() *AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest {
	return &AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.lsyd.listupout.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// Set is BillCode Setter
// 单据编码
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// Get BillCode Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) GetBillCode() string {
	return r._billCode
}

// Set is FromRefUserId Setter
// 发货企业renEntId
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) SetFromRefUserId(_fromRefUserId string) error {
	r._fromRefUserId = _fromRefUserId
	r.Set("from_ref_user_id", _fromRefUserId)
	return nil
}

// Get FromRefUserId Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) GetFromRefUserId() string {
	return r._fromRefUserId
}

// Set is ToRefUserId Setter
// 收货企业refEntId
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) SetToRefUserId(_toRefUserId string) error {
	r._toRefUserId = _toRefUserId
	r.Set("to_ref_user_id", _toRefUserId)
	return nil
}

// Get ToRefUserId Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) GetToRefUserId() string {
	return r._toRefUserId
}
