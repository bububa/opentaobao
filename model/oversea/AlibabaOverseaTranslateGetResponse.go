package oversea

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取文本翻译信息 API返回值 
alibaba.oversea.translate.get

根据传入的文本信息，获取其目标语言的翻译结果
*/
type AlibabaOverseaTranslateGetAPIResponse struct {
    model.CommonResponse
    AlibabaOverseaTranslateGetResponse
}

// 获取文本翻译信息 成功返回结果
type AlibabaOverseaTranslateGetResponse struct {
    XMLName xml.Name `xml:"alibaba_oversea_translate_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *SimpleTransResult `json:"result,omitempty" xml:"result,omitempty"`
}
