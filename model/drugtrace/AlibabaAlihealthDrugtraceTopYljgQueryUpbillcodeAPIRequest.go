package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest 通过一个码，查询这个码对应的上游企业出库单的单据号 API请求
// alibaba.alihealth.drugtrace.top.yljg.query.upbillcode
//
// 一个查询上游出库单号的接口。企业在扫码入库时，接口通过扫到的码判定这个码对应的上游企业所属的出库单据号
type AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 企业ID （一般为要查询单据的收货企业）
	_refEntId string
}

// NewAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest 初始化AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest() *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest {
	return &AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest) Reset() {
	r._code = ""
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.yljg.query.upbillcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest) GetCode() string {
	return r._code
}

// SetRefEntId is RefEntId Setter
// 企业ID （一般为要查询单据的收货企业）
func (r *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest()
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest
func GetAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest() *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest {
	return poolAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest.Get().(*AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest 将 AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest(v *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest.Put(v)
}
