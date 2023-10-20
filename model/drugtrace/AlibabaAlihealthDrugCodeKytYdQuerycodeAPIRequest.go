package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest 查询追溯码对应的药品信息（药店） API请求
// alibaba.alihealth.drug.code.kyt.yd.querycode
//
// 此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest struct {
	model.Params
	// 码列表
	_codes []string
	// 企业唯一标识
	_refEntId string
}

// NewAlibabaAlihealthDrugCodeKytYdQuerycodeRequest 初始化AlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest对象
func NewAlibabaAlihealthDrugCodeKytYdQuerycodeRequest() *AlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest {
	return &AlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest) Reset() {
	r._codes = r._codes[:0]
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.yd.querycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodes is Codes Setter
// 码列表
func (r *AlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest) GetCodes() []string {
	return r._codes
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识
func (r *AlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugCodeKytYdQuerycodeRequest()
	},
}

// GetAlibabaAlihealthDrugCodeKytYdQuerycodeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest
func GetAlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest() *AlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest {
	return poolAlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest.Get().(*AlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest 将 AlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest(v *AlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeKytYdQuerycodeAPIRequest.Put(v)
}
