package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest 通过一个码查询上游出库单 API请求
// alibaba.alihealth.drug.kyt.query.upbillcode
//
// 一个查询上游出库单号的接口。企业在扫码入库时，接口通过扫到的码判定这个码对应的上游企业所属的出库单据号
type AlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 企业REF_ENT_ID （当前企业的唯一标识）
	_refEntId string
}

// NewAlibabaAlihealthDrugKytQueryUpbillcodeRequest 初始化AlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest对象
func NewAlibabaAlihealthDrugKytQueryUpbillcodeRequest() *AlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest {
	return &AlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest) Reset() {
	r._code = ""
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.query.upbillcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest) GetCode() string {
	return r._code
}

// SetRefEntId is RefEntId Setter
// 企业REF_ENT_ID （当前企业的唯一标识）
func (r *AlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytQueryUpbillcodeRequest()
	},
}

// GetAlibabaAlihealthDrugKytQueryUpbillcodeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest
func GetAlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest() *AlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest {
	return poolAlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest.Get().(*AlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest 将 AlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest(v *AlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest.Put(v)
}
