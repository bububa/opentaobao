package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
非药单码替换 APIRequest
alibaba.alihealth.tracecodeseller.code.single.codereplace

提供非药追溯码单码替换功能
*/
type AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest struct {
    model.Params

    // 企业id
    entInfoId   string 

    // 新码
    newCode   string 

    // 老码
    oldCode   string 

}

func NewAlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest() *AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest{
    return &AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.code.single.codereplace"
}

func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest) SetEntInfoId(entInfoId string) error {
    r.entInfoId = entInfoId
    r.Set("ent_info_id", entInfoId)
    return nil
}

func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest) GetEntInfoId() string {
    return r.entInfoId
}

func (r *AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest) SetNewCode(newCode string) error {
    r.newCode = newCode
    r.Set("new_code", newCode)
    return nil
}

func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest) GetNewCode() string {
    return r.newCode
}

func (r *AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest) SetOldCode(oldCode string) error {
    r.oldCode = oldCode
    r.Set("old_code", oldCode)
    return nil
}

func (r AlibabaAlihealthTracecodesellerCodeSingleCodereplaceRequest) GetOldCode() string {
    return r.oldCode
}

