package nlp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
文本语言相似度 APIResponse
taobao.nlp.similarity

文本语言相似度计算，提供余弦距离、编辑距离和simHash三种相似度计算。返回文本相似度区间为0-1之间，0为完全不相似，1为完全相似。
*/
type TaobaoNlpSimilarityAPIResponse struct {
    model.CommonResponse
    Response *TaobaoNlpSimilarityResponse `json:"taobao_nlp_similarity_response,omitempty"`
}

type TaobaoNlpSimilarityResponse struct {

    // 返回结果
    Simresult   *SimResult `json:"simresult,omitempty"`

}
