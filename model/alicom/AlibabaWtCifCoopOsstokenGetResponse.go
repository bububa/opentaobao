package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取oss签名接口 API返回值 
alibaba.wt.cif.coop.osstoken.get

商家合作上传oss图片获取token接口
*/
type AlibabaWtCifCoopOsstokenGetAPIResponse struct {
    model.CommonResponse
    AlibabaWtCifCoopOsstokenGetResponse
}

// 获取oss签名接口 成功返回结果
type AlibabaWtCifCoopOsstokenGetResponse struct {
    XMLName xml.Name `xml:"alibaba_wt_cif_coop_osstoken_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *OssTokenGetResult `json:"result,omitempty" xml:"result,omitempty"`
}
