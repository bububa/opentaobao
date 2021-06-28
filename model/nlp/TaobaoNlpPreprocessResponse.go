package nlp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
文本语言预处理 APIResponse
taobao.nlp.preprocess

实现文本语言处理中的预处理功能，如实现文字繁简转换、文字转拼音、文字拆分、谐音同音字替换和形似字替换。
*/
type TaobaoNlpPreprocessAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoNlpPreprocessResponse `json:"nlp_preprocess_response,omitempty"` 
    TaobaoNlpPreprocessResponse
}

/* model for simplify = false
type TaobaoNlpPreprocessResponse struct {

    // 返回结果
    
    Processresult  *struct {
        ProcessResult  *ProcessResult `json:"process_result,omitempty"`
    } `json:"processresult,omitempty"`
    

}
*/

type TaobaoNlpPreprocessResponse struct {

    // 返回结果
    Processresult   *ProcessResult `json:"processresult,omitempty"`

}
