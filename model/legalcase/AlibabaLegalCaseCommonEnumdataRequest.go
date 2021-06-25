package legalcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取通用枚举值接口 APIRequest
alibaba.legal.case.common.enumdata

获取通用枚举值接口
*/
type AlibabaLegalCaseCommonEnumdataRequest struct {
    model.Params

    // bu
    key   string 

    // 语言
    lang   string 

}

func NewAlibabaLegalCaseCommonEnumdataRequest() *AlibabaLegalCaseCommonEnumdataRequest{
    return &AlibabaLegalCaseCommonEnumdataRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLegalCaseCommonEnumdataRequest) GetApiMethodName() string {
    return "alibaba.legal.case.common.enumdata"
}

func (r AlibabaLegalCaseCommonEnumdataRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLegalCaseCommonEnumdataRequest) SetKey(key string) error {
    r.key = key
    r.Set("key", key)
    return nil
}

func (r AlibabaLegalCaseCommonEnumdataRequest) GetKey() string {
    return r.key
}

func (r *AlibabaLegalCaseCommonEnumdataRequest) SetLang(lang string) error {
    r.lang = lang
    r.Set("lang", lang)
    return nil
}

func (r AlibabaLegalCaseCommonEnumdataRequest) GetLang() string {
    return r.lang
}

