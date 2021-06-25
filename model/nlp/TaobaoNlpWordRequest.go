package nlp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
文本语言词法分析 APIRequest
taobao.nlp.word

提供文本语言处理中的词法分析功能,开放中文分词和词权重计算功能。
*/
type TaobaoNlpWordRequest struct {
    model.Params

    // 功能类型选择：1)wType=1时提供分词功能，type=0时为基本粒度，type=1时为混合粒度，type=3时为基本粒度和混合粒度共同输出；
    wType   int64 

    // 文本内容
    text   *Text 

}

func NewTaobaoNlpWordRequest() *TaobaoNlpWordRequest{
    return &TaobaoNlpWordRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoNlpWordRequest) GetApiMethodName() string {
    return "taobao.nlp.word"
}

func (r TaobaoNlpWordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoNlpWordRequest) SetWType(wType int64) error {
    r.wType = wType
    r.Set("w_type", wType)
    return nil
}

func (r TaobaoNlpWordRequest) GetWType() int64 {
    return r.wType
}

func (r *TaobaoNlpWordRequest) SetText(text *Text) error {
    r.text = text
    r.Set("text", text)
    return nil
}

func (r TaobaoNlpWordRequest) GetText() *Text {
    return r.text
}

