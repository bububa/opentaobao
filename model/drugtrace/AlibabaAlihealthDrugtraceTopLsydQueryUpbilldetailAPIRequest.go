package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest 根据单据号查询单据的详情信息【注意：查询的是本企业的单据】 API请求
// alibaba.alihealth.drugtrace.top.lsyd.query.upbilldetail
//
// 根据单据号查询单据的详情信息【注意：这个接口查询的是本企业的单据，如果是查询上游的单据明细信息，使用xxxxxxx.listupout.detail接口】。
type AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest struct {
	model.Params
	// 单据号码
	_billCode string
	// 本企业refEntId
	_refEntId string
}

// NewAlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest 初始化AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest() *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest {
	return &AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest) Reset() {
	r._billCode = ""
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.lsyd.query.upbilldetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBillCode is BillCode Setter
// 单据号码
func (r *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetRefEntId is RefEntId Setter
// 本企业refEntId
func (r *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest()
	},
}

// GetAlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailRequest 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest
func GetAlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest() *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest {
	return poolAlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest.Get().(*AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest)
}

// ReleaseAlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest 将 AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest(v *AlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopLsydQueryUpbilldetailAPIRequest.Put(v)
}
