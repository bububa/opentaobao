package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取oss签名接口 APIResponse
alibaba.wt.cif.coop.osstoken.get

商家合作上传oss图片获取token接口
*/
type AlibabaWtCifCoopOsstokenGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWtCifCoopOsstokenGetResponse `json:"alibaba_wt_cif_coop_osstoken_get_response,omitempty"`
}

type AlibabaWtCifCoopOsstokenGetResponse struct {

    // result
    Result   *OssTokenGetResult `json:"result,omitempty"`

}
