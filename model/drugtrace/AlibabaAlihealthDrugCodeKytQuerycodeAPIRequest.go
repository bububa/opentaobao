package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest 查询追溯码对应的药品信息 API请求
// alibaba.alihealth.drug.code.kyt.querycode
//
// 此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest struct {
	model.Params
	// 码列表【多个码用逗号分隔的字符串。要求数量在1000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
	_codes []string
	// 企业唯一标识
	_refEntId string
}

// NewAlibabaAlihealthDrugCodeKytQuerycodeRequest 初始化AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest对象
func NewAlibabaAlihealthDrugCodeKytQuerycodeRequest() *AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest {
	return &AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest) Reset() {
	r._codes = r._codes[:0]
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.querycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodes is Codes Setter
// 码列表【多个码用逗号分隔的字符串。要求数量在1000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
func (r *AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest) GetCodes() []string {
	return r._codes
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识
func (r *AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugCodeKytQuerycodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugCodeKytQuerycodeRequest()
	},
}

// GetAlibabaAlihealthDrugCodeKytQuerycodeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest
func GetAlibabaAlihealthDrugCodeKytQuerycodeAPIRequest() *AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest {
	return poolAlibabaAlihealthDrugCodeKytQuerycodeAPIRequest.Get().(*AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugCodeKytQuerycodeAPIRequest 将 AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeKytQuerycodeAPIRequest(v *AlibabaAlihealthDrugCodeKytQuerycodeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeKytQuerycodeAPIRequest.Put(v)
}
