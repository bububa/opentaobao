package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商家信息 API请求
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

// 初始化AlibabaAlihealthTracecodesellerEntSearchRequest对象
func NewAlibabaAlihealthTracecodesellerEntSearchRequest() *AlibabaAlihealthTracecodesellerEntSearchRequest{
    return &AlibabaAlihealthTracecodesellerEntSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerEntSearchRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.ent.search"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerEntSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkeyCode Setter
// appkey
func (r *AlibabaAlihealthTracecodesellerEntSearchRequest) SetSkeyCode(skeyCode string) error {
    r.skeyCode = skeyCode
    r.Set("skey_code", skeyCode)
    return nil
}

// SkeyCode Getter
func (r AlibabaAlihealthTracecodesellerEntSearchRequest) GetSkeyCode() string {
    return r.skeyCode
}
// Name Setter
// 商家名称
func (r *AlibabaAlihealthTracecodesellerEntSearchRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r AlibabaAlihealthTracecodesellerEntSearchRequest) GetName() string {
    return r.name
}
// TbUserId Setter
// 淘宝名
func (r *AlibabaAlihealthTracecodesellerEntSearchRequest) SetTbUserId(tbUserId string) error {
    r.tbUserId = tbUserId
    r.Set("tb_user_id", tbUserId)
    return nil
}

// TbUserId Getter
func (r AlibabaAlihealthTracecodesellerEntSearchRequest) GetTbUserId() string {
    return r.tbUserId
}
