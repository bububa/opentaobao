package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest 扫码营销码查询 API请求
// alibaba.alihealth.drug.code.kyt.smyx.querycode
//
// 此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest struct {
	model.Params
	// 码列表
	_codes []string
	// 企业唯一标识
	_refEntId string
}

// NewAlibabaAlihealthDrugCodeKytSmyxQuerycodeRequest 初始化AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest对象
func NewAlibabaAlihealthDrugCodeKytSmyxQuerycodeRequest() *AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest {
	return &AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest) Reset() {
	r._codes = r._codes[:0]
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.smyx.querycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodes is Codes Setter
// 码列表
func (r *AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest) GetCodes() []string {
	return r._codes
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识
func (r *AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugCodeKytSmyxQuerycodeRequest()
	},
}

// GetAlibabaAlihealthDrugCodeKytSmyxQuerycodeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest
func GetAlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest() *AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest {
	return poolAlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest.Get().(*AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest 将 AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest(v *AlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeKytSmyxQuerycodeAPIRequest.Put(v)
}
