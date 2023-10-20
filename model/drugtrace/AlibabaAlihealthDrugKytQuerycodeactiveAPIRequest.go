package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest 查询码是否激活 API请求
// alibaba.alihealth.drug.kyt.querycodeactive
//
// 查询码是否激活
type AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest struct {
	model.Params
	// 码列表【多个码时用逗号拼接的字符串。要求数量在4000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
	_codes []string
	// 企业唯一标识
	_refEntId string
}

// NewAlibabaAlihealthDrugKytQuerycodeactiveRequest 初始化AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest对象
func NewAlibabaAlihealthDrugKytQuerycodeactiveRequest() *AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest {
	return &AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest) Reset() {
	r._codes = r._codes[:0]
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.querycodeactive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodes is Codes Setter
// 码列表【多个码时用逗号拼接的字符串。要求数量在4000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
func (r *AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest) GetCodes() []string {
	return r._codes
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识
func (r *AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugKytQuerycodeactiveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytQuerycodeactiveRequest()
	},
}

// GetAlibabaAlihealthDrugKytQuerycodeactiveRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest
func GetAlibabaAlihealthDrugKytQuerycodeactiveAPIRequest() *AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest {
	return poolAlibabaAlihealthDrugKytQuerycodeactiveAPIRequest.Get().(*AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytQuerycodeactiveAPIRequest 将 AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytQuerycodeactiveAPIRequest(v *AlibabaAlihealthDrugKytQuerycodeactiveAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytQuerycodeactiveAPIRequest.Put(v)
}
