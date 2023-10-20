package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytCodereplaceAPIRequest 单码替换 API请求
// alibaba.alihealth.drug.kyt.codereplace
//
// 单码替换
type AlibabaAlihealthDrugKytCodereplaceAPIRequest struct {
	model.Params
	// 企业ref_ent_id（码申请企业）
	_refEntId string
	// 替换后的追溯码
	_newCode string
	// 被替换的追溯码
	_oldCode string
}

// NewAlibabaAlihealthDrugKytCodereplaceRequest 初始化AlibabaAlihealthDrugKytCodereplaceAPIRequest对象
func NewAlibabaAlihealthDrugKytCodereplaceRequest() *AlibabaAlihealthDrugKytCodereplaceAPIRequest {
	return &AlibabaAlihealthDrugKytCodereplaceAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytCodereplaceAPIRequest) Reset() {
	r._refEntId = ""
	r._newCode = ""
	r._oldCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytCodereplaceAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.codereplace"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytCodereplaceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytCodereplaceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ref_ent_id（码申请企业）
func (r *AlibabaAlihealthDrugKytCodereplaceAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytCodereplaceAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetNewCode is NewCode Setter
// 替换后的追溯码
func (r *AlibabaAlihealthDrugKytCodereplaceAPIRequest) SetNewCode(_newCode string) error {
	r._newCode = _newCode
	r.Set("new_code", _newCode)
	return nil
}

// GetNewCode NewCode Getter
func (r AlibabaAlihealthDrugKytCodereplaceAPIRequest) GetNewCode() string {
	return r._newCode
}

// SetOldCode is OldCode Setter
// 被替换的追溯码
func (r *AlibabaAlihealthDrugKytCodereplaceAPIRequest) SetOldCode(_oldCode string) error {
	r._oldCode = _oldCode
	r.Set("old_code", _oldCode)
	return nil
}

// GetOldCode OldCode Getter
func (r AlibabaAlihealthDrugKytCodereplaceAPIRequest) GetOldCode() string {
	return r._oldCode
}

var poolAlibabaAlihealthDrugKytCodereplaceAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytCodereplaceRequest()
	},
}

// GetAlibabaAlihealthDrugKytCodereplaceRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytCodereplaceAPIRequest
func GetAlibabaAlihealthDrugKytCodereplaceAPIRequest() *AlibabaAlihealthDrugKytCodereplaceAPIRequest {
	return poolAlibabaAlihealthDrugKytCodereplaceAPIRequest.Get().(*AlibabaAlihealthDrugKytCodereplaceAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytCodereplaceAPIRequest 将 AlibabaAlihealthDrugKytCodereplaceAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytCodereplaceAPIRequest(v *AlibabaAlihealthDrugKytCodereplaceAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytCodereplaceAPIRequest.Put(v)
}
