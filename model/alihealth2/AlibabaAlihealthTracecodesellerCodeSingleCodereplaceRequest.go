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
type AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest struct {
    model.Params
    // 企业id
    _entInfoId   string
    // 新码
    _newCode   string
    // 老码
    _oldCode   string
}

// 初始化AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest对象
func NewAlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest() *AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest{
    return &AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.code.single.codereplace"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EntInfoId Setter
// 企业id
func (r *AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest) SetEntInfoId(_entInfoId string) error {
    r._entInfoId = _entInfoId
    r.Set("ent_info_id", _entInfoId)
    return nil
}

// EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest) GetEntInfoId() string {
    return r._entInfoId
}
// NewCode Setter
// 新码
func (r *AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest) SetNewCode(_newCode string) error {
    r._newCode = _newCode
    r.Set("new_code", _newCode)
    return nil
}

// NewCode Getter
func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest) GetNewCode() string {
    return r._newCode
}
// OldCode Setter
// 老码
func (r *AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest) SetOldCode(_oldCode string) error {
    r._oldCode = _oldCode
    r.Set("old_code", _oldCode)
    return nil
}

// OldCode Getter
func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest) GetOldCode() string {
    return r._oldCode
}
