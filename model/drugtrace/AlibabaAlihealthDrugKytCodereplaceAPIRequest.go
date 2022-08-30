package drugtrace

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytCodereplaceAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.codereplace"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytCodereplaceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
