package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
非药单码替换 API请求
alibaba.alihealth.tracecodeseller.code.single.codereplace

提供非药追溯码单码替换功能
*/
type AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest struct {
    model.Params
    // 企业id
    _entInfoId   string
    // 新码
    _newCode   string
    // 老码
    _oldCode   string
}

// 初始化AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest对象
func NewAlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest() *AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest{
    return &AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.code.single.codereplace"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EntInfoId Setter
// 企业id
func (r *AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest) SetEntInfoId(_entInfoId string) error {
    r._entInfoId = _entInfoId
    r.Set("ent_info_id", _entInfoId)
    return nil
}

// EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest) GetEntInfoId() string {
    return r._entInfoId
}
// NewCode Setter
// 新码
func (r *AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest) SetNewCode(_newCode string) error {
    r._newCode = _newCode
    r.Set("new_code", _newCode)
    return nil
}

// NewCode Getter
func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest) GetNewCode() string {
    return r._newCode
}
// OldCode Setter
// 老码
func (r *AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest) SetOldCode(_oldCode string) error {
    r._oldCode = _oldCode
    r.Set("old_code", _oldCode)
    return nil
}

// OldCode Getter
func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceAPIRequest) GetOldCode() string {
    return r._oldCode
}
