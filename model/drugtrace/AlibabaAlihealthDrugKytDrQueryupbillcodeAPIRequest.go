package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest 查询上游企业出库单据号 API请求
// alibaba.alihealth.drug.kyt.dr.queryupbillcode
//
// 疫苗温度合规补充需求-增加一个查询上游出库单号的接口。疾控在扫码入库时，接口通过扫到的码判定这个码对应所属的出库单据号
type AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 企业ID （一般为要查询单据的收货企业）
	_refEntId string
}

// NewAlibabaAlihealthDrugKytDrQueryupbillcodeRequest 初始化AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest对象
func NewAlibabaAlihealthDrugKytDrQueryupbillcodeRequest() *AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest {
	return &AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest) Reset() {
	r._code = ""
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.dr.queryupbillcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest) GetCode() string {
	return r._code
}

// SetRefEntId is RefEntId Setter
// 企业ID （一般为要查询单据的收货企业）
func (r *AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytDrQueryupbillcodeRequest()
	},
}

// GetAlibabaAlihealthDrugKytDrQueryupbillcodeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest
func GetAlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest() *AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest {
	return poolAlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest.Get().(*AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest 将 AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest(v *AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest.Put(v)
}
