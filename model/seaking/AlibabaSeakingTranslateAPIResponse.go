package seaking

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
MT定制接口 API返回值 
alibaba.seaking.translate

MT定制接口
*/
type AlibabaSeakingTranslateAPIResponse struct {
    model.CommonResponse
    AlibabaSeakingTranslateAPIResponseModel
}

// MT定制接口 成功返回结果
type AlibabaSeakingTranslateAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_seaking_translate_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 译文
    Result   string `json:"result,omitempty" xml:"result,omitempty"`
}
