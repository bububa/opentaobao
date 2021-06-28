package nlp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
文本语言词法分析 APIResponse
taobao.nlp.word

提供文本语言处理中的词法分析功能,开放中文分词和词权重计算功能。
*/
type TaobaoNlpWordAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"nlp_word_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回词法分析的结果
    
    Wordresult   *WordResult `json:"wordresult,omitempty" xml:"