package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeNodenameGetAPIRequest 根据码获取机构名称 API请求
// alibaba.alihealth.drugcode.nodename.get
//
// 根据码获取机构名称
type AlibabaAlihealthDrugcodeNodenameGetAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 企业id
	_refEntId string
}

// NewAlibabaAlihealthDrugcodeNodenameGetRequest 初始化AlibabaAlihealthDrugcodeNodenameGetAPIRequest对象
func NewAlibabaAlihealthDrugcodeNodenameGetRequest() *AlibabaAlihealthDrugcodeNodenameGetAPIRequest {
	return &AlibabaAlihealthDrugcodeNodenameGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugcodeNodenameGetAPIRequest) Reset() {
	r._code = ""
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeNodenameGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.nodename.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeNodenameGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugcodeNodenameGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugcodeNodenameGetAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugcodeNodenameGetAPIRequest) GetCode() string {
	return r._code
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugcodeNodenameGetAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugcodeNodenameGetAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugcodeNodenameGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugcodeNodenameGetRequest()
	},
}

// GetAlibabaAlihealthDrugcodeNodenameGetRequest 从 sync.Pool 获取 AlibabaAlihealthDrugcodeNodenameGetAPIRequest
func GetAlibabaAlihealthDrugcodeNodenameGetAPIRequest() *AlibabaAlihealthDrugcodeNodenameGetAPIRequest {
	return poolAlibabaAlihealthDrugcodeNodenameGetAPIRequest.Get().(*AlibabaAlihealthDrugcodeNodenameGetAPIRequest)
}

// ReleaseAlibabaAlihealthDrugcodeNodenameGetAPIRequest 将 AlibabaAlihealthDrugcodeNodenameGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeNodenameGetAPIRequest(v *AlibabaAlihealthDrugcodeNodenameGetAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeNodenameGetAPIRequest.Put(v)
}
