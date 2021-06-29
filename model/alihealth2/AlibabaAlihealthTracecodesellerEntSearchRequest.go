package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商家信息 APIRequest
alibaba.alihealth.tracecodeseller.ent.search

查询商家信息
*/
type AlibabaAlihealthTracecodesellerEntSearchRequest struct {
    model.Params

    // appkey
    skeyCode   string 

    // 商家名称
    name   string 

    // 淘宝名
    tbUserId   string 

}

func NewAlibabaAlihealthTracecodesellerEntSearchRequest() *AlibabaAlihealthTracecodesellerEntSearchRequest{
    return &AlibabaAlihealthTracecodesellerEntSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthTracecodesellerEntSearchRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.ent.search"
}

func (r AlibabaAlihealthTracecodesellerEntSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthTracecodesellerEntSearchRequest) SetSkeyCode(skeyCode string) error {
    r.skeyCode = skeyCode
    r.Set("skey_code", skeyCode)
    return nil
}

func (r AlibabaAlihealthTracecodesellerEntSearchRequest) GetSkeyCode() string {
    return r.skeyCode
}

func (r *AlibabaAlihealthTracecodesellerEntSearchRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r AlibabaAlihealthTracecodesellerEntSearchRequest) GetName() string {
    return r.name
}

func (r *AlibabaAlihealthTracecodesellerEntSearchRequest) SetTbUserId(tbUserId string) error {
    r.tbUserId = tbUserId
    r.Set("tb_user_id", tbUserId)
    return nil
}

func (r AlibabaAlihealthTracecodesellerEntSearchRequest) GetTbUserId() string {
    return r.tbUserId
}

