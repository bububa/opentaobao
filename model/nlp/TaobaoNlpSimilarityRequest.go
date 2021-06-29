package nlp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
文本语言相似度 API请求
taobao.nlp.similarity

文本语言相似度计算，提供余弦距离、编辑距离和simHash三种相似度计算。返回文本相似度区间为0-1之间，0为完全不相似，1为完全相似。
*/
type TaobaoNlpSimilarityRequest struct {
    model.Params
    // 多文本内容
    texts   *Texts
}

// 初始化TaobaoNlpSimilarityRequest对象
func NewTaobaoNlpSimilarityRequest() *TaobaoNlpSimilarityRequest{
    return &TaobaoNlpSimilarityRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoNlpSimilarityRequest) GetApiMethodName() string {
    return "taobao.nlp.similarity"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoNlpSimilarityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Texts Setter
// 多文本内容
func (r *TaobaoNlpSimilarityRequest) SetTexts(texts *Texts) error {
    r.texts = texts
    r.Set("texts", texts)
    return nil
}

// Texts Getter
func (r TaobaoNlpSimilarityRequest) GetTexts() *Texts {
    return r.texts
}
