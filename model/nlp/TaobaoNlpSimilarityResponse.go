package nlp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
文本语言相似度 APIResponse
taobao.nlp.similarity

文本语言相似度计算，提供余弦距离、编辑距离和simHash三种相似度计算。返回文本相似度区间为0-1之间，0为完全不相似，1为完全相似。
*/
type TaobaoNlpSimilarityAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"nlp_similarity_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Simresult   *SimResult `json:"simresult,omitempty" xml:"