package ioti

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
对混合云提供的刷图接口 API返回值 
alibaba.it.esl.eslimage.showimagecommon

混合云使用，提供给isv和我们混合云环境部署的应用刷图
*/
type AlibabaItEslEslimageShowimagecommonAPIResponse struct {
    model.CommonResponse
    AlibabaItEslEslimageShowimagecommonResponse
}

// 对混合云提供的刷图接口 成功返回结果
type AlibabaItEslEslimageShowimagecommonResponse struct {
    XMLName xml.Name `xml:"alibaba_it_esl_eslimage_showimagecommon_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // resultmsg
    Resultmsg   string `json:"resultmsg,omitempty" xml:"resultmsg,omitempty"`
}
