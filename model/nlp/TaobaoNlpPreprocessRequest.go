package nlp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
文本语言预处理 APIRequest
taobao.nlp.preprocess

实现文本语言处理中的预处理功能，如实现文字繁简转换、文字转拼音、文字拆分、谐音同音字替换和形似字替换。
*/
type TaobaoNlpPreprocessRequest struct {
    model.Params

    // 文本内容
    text   *Text 

    // 谐音字转换、形似字转换需提供关键词进行替换，关键词之间以#分隔。keyword示例：兼职#招聘#微信、天猫#日结#微信#招聘#加微
    keyword   string 

    // 1)繁简字转换：func_type=1，对应type =1 繁转简 type=2 简转繁；2)拆分字转换：func_type =2，对应type=1 文字拆分 type=2 拆分字合并；3)文字转拼音：func_type =3，对应type=1 文字转拼音 type=2 拼音+声调；4)谐音同音字替换：func_type =4，对应type=1 谐音字替换 type=2 同音字替换；5)形似字替换：func_type =5，对应type=1 形似字替换;
    funcType   int64 

}

func NewTaobaoNlpPreprocessRequest() *TaobaoNlpPreprocessRequest{
    return &TaobaoNlpPreprocessRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoNlpPreprocessRequest) GetApiMethodName() string {
    return "taobao.nlp.preprocess"
}

func (r TaobaoNlpPreprocessRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoNlpPreprocessRequest) SetText(text *Text) error {
    r.text = text
    r.Set("text", text)
    return nil
}

func (r TaobaoNlpPreprocessRequest) GetText() *Text {
    return r.text
}

func (r *TaobaoNlpPreprocessRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

func (r TaobaoNlpPreprocessRequest) GetKeyword() string {
    return r.keyword
}

func (r *TaobaoNlpPreprocessRequest) SetFuncType(funcType int64) error {
    r.funcType = funcType
    r.Set("func_type", funcType)
    return nil
}

func (r TaobaoNlpPreprocessRequest) GetFuncType() int64 {
    return r.funcType
}

