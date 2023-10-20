package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest 查询一个码的所有子码 API请求
// alibaba.alihealth.drug.kyt.yy.querysubcodes
//
// 单码的了码查询
type AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest struct {
	model.Params
	// 码
	_codes []string
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
}

// NewAlibabaAlihealthDrugKytYyQuerysubcodesRequest 初始化AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest对象
func NewAlibabaAlihealthDrugKytYyQuerysubcodesRequest() *AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest {
	return &AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest) Reset() {
	r._codes = r._codes[:0]
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.yy.querysubcodes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodes is Codes Setter
// 码
func (r *AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest) GetCodes() []string {
	return r._codes
}

// SetRefEntId is RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytYyQuerysubcodesRequest()
	},
}

// GetAlibabaAlihealthDrugKytYyQuerysubcodesRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest
func GetAlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest() *AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest {
	return poolAlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest.Get().(*AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest 将 AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest(v *AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest.Put(v)
}
