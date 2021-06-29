package legalcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取通用枚举值接口 API请求
alibaba.legal.case.common.enumdata

获取通用枚举值接口
*/
type AlibabaLegalCaseCommonEnumdataRequest struct {
    model.Params
    // bu
    _key   string
    // 语言
    _lang   string
}

// 初始化AlibabaLegalCaseCommonEnumdataRequest对象
func NewAlibabaLegalCaseCommonEnumdataRequest() *AlibabaLegalCaseCommonEnumdataRequest{
    return &AlibabaLegalCaseCommonEnumdataRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseCommonEnumdataRequest) GetApiMethodName() string {
    return "alibaba.legal.case.common.enumdata"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseCommonEnumdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Key Setter
// bu
func (r *AlibabaLegalCaseCommonEnumdataRequest) SetKey(_key string) error {
    r._key = _key
    r.Set("key", _key)
    return nil
}

// Key Getter
func (r AlibabaLegalCaseCommonEnumdataRequest) GetKey() string {
    return r._key
}
// Lang Setter
// 语言
func (r *AlibabaLegalCaseCommonEnumdataRequest) SetLang(_lang string) error {
    r._lang = _lang
    r.Set("lang", _lang)
    return nil
}

// Lang Getter
func (r AlibabaLegalCaseCommonEnumdataRequest) GetLang() string {
    return r._lang
}
