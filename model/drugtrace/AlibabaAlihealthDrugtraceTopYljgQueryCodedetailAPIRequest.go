package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest 根据码查询码信息 API请求
// alibaba.alihealth.drugtrace.top.yljg.query.codedetail
//
// 服务描述
// 此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息。
// 核查平台优先过滤非8开头的，长度非20位数字的码信息。
type AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest struct {
	model.Params
	// 码列表【多个码用逗号拼接的字符串。要求数量在1000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
	_codes []string
	// 企业唯一标识
	_refEntId string
}

// NewAlibabaAlihealthDrugtraceTopYljgQueryCodedetailRequest 初始化AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopYljgQueryCodedetailRequest() *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest {
	return &AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest) Reset() {
	r._codes = r._codes[:0]
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.yljg.query.codedetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCodes is Codes Setter
// 码列表【多个码用逗号拼接的字符串。要求数量在1000个码以下，但一般不要传这么多，如果网络不好很容易传输一半报错】
func (r *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest) SetCodes(_codes []string) error {
	r._codes = _codes
	r.Set("codes", _codes)
	return nil
}

// GetCodes Codes Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest) GetCodes() []string {
	return r._codes
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识
func (r *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugtraceTopYljgQueryCodedetailRequest()
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgQueryCodedetailRequest 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest
func GetAlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest() *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest {
	return poolAlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest.Get().(*AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest 将 AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest(v *AlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopYljgQueryCodedetailAPIRequest.Put(v)
}
