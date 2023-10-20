package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest 上游出库单单据明细查询 API请求
// alibaba.alihealth.drugtrace.top.lsyd.listupout.detail
//
// 查询上游出库单明细(带追溯码信息)
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
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) Reset() {
	r._refEntId = ""
	r._billCode = ""
	r._fromRefUserId = ""
	r._toRefUserId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.lsyd.listupout.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBillCode is BillCode Setter
// 单据编码
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetFromRefUserId is FromRefUserId Setter
// 发货企业renEntId
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) SetFromRefUserId(_fromRefUserId string) error {
	r._fromRefUserId = _fromRefUserId
	r.Set("from_ref_user_id", _fromRefUserId)
	return nil
}

// GetFromRefUserId FromRefUserId Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) GetFromRefUserId() string {
	return r._fromRefUserId
}

// SetToRefUserId is ToRefUserId Setter
// 收货企业refEntId
func (r *AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) SetToRefUserId(_toRefUserId string) error {
	r._toRefUserId = _toRefUserId
	r.Set("to_ref_user_id", _toRefUserId)
	return nil
}

// GetToRefUserId ToRefUserId Getter
func (r AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) GetToRefUserId() string {
	return r._toRefUserId
}

var poolAlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugtraceTopLsydListupoutDetailRequest()
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydListupoutDetailRequest 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest
func GetAlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest() *AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest {
	return poolAlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest.Get().(*AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest 将 AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest(v *AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest.Put(v)
}
