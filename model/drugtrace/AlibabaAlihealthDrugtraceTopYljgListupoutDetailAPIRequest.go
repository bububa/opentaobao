package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest 上游出库单单据明细查询 API请求
// alibaba.alihealth.drugtrace.top.yljg.listupout.detail
//
// 查询上游出库单明细(带追溯码信息)
type AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest struct {
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

// NewAlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest 初始化AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest() *AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest {
	return &AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) Reset() {
	r._refEntId = ""
	r._billCode = ""
	r._fromRefUserId = ""
	r._toRefUserId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.yljg.listupout.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBillCode is BillCode Setter
// 单据编码
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetFromRefUserId is FromRefUserId Setter
// 发货企业renEntId
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) SetFromRefUserId(_fromRefUserId string) error {
	r._fromRefUserId = _fromRefUserId
	r.Set("from_ref_user_id", _fromRefUserId)
	return nil
}

// GetFromRefUserId FromRefUserId Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) GetFromRefUserId() string {
	return r._fromRefUserId
}

// SetToRefUserId is ToRefUserId Setter
// 收货企业refEntId
func (r *AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) SetToRefUserId(_toRefUserId string) error {
	r._toRefUserId = _toRefUserId
	r.Set("to_ref_user_id", _toRefUserId)
	return nil
}

// GetToRefUserId ToRefUserId Getter
func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) GetToRefUserId() string {
	return r._toRefUserId
}

var poolAlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest()
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest
func GetAlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest() *AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest {
	return poolAlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest.Get().(*AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest 将 AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest(v *AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest.Put(v)
}
