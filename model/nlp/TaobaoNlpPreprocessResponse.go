package nlp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
文本语言预处理 APIResponse
taobao.nlp.preprocess

实现文本语言处理中的预处理功能，如实现文字繁简转换、文字转拼音、文字拆分、谐音同音字替换和形似字替换。
*/
type TaobaoNlpPreprocessAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"nlp_preprocess_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Processresult   *ProcessResult `json:"processresult,omitempty" xml:"