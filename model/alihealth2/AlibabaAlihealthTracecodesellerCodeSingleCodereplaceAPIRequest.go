package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest 非药单码替换 API请求
// alibaba.alihealth.tracecodeseller.code.single.codereplace
//
// 提供非药追溯码单码替换功能
type AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest struct {
	model.Params
	// 企业id
	_entInfoId string
	// 新码
	_newCode string
	// 老码
	_oldCode string
}

// NewAlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest 初始化AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest对象
func NewAlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest() *AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest {
	return &AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.code.single.codereplace"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntInfoId is EntInfoId Setter
// 企业id
func (r *AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest) SetEntInfoId(_entInfoId string) error {
	r._entInfoId = _entInfoId
	r.Set("ent_info_id", _entInfoId)
	return nil
}

// GetEntInfoId EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest) GetEntInfoId() string {
	return r._entInfoId
}

// SetNewCode is NewCode Setter
// 新码
func (r *AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest) SetNewCode(_newCode string) error {
	r._newCode = _newCode
	r.Set("new_code", _newCode)
	return nil
}

// GetNewCode NewCode Getter
func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest) GetNewCode() string {
	return r._newCode
}

// SetOldCode is OldCode Setter
// 老码
func (r *AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest) SetOldCode(_oldCode string) error {
	r._oldCode = _oldCode
	r.Set("old_code", _oldCode)
	return nil
}

// GetOldCode OldCode Getter
func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest) GetOldCode() string {
	return r._oldCode
}
