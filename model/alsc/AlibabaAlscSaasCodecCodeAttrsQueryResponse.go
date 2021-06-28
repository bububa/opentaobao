package alsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
码业务属性查询 APIResponse
alibaba.alsc.saas.codec.code.attrs.query

码业务属性查询
*/
type AlibabaAlscSaasCodecCodeAttrsQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlscSaasCodecCodeAttrsQueryResponse `json:"alibaba_alsc_saas_codec_code_attrs_query_response,omitempty"` 
    AlibabaAlscSaasCodecCodeAttrsQueryResponse
}

/* model for simplify = false
type AlibabaAlscSaasCodecCodeAttrsQueryResponse struct {

    // 接口返回model
    
    Result  *struct {
        AlibabaAlscSaasCodecCodeAttrsQueryResult  *AlibabaAlscSaasCodecCodeAttrsQueryResult `json:"alibaba_alsc_saas_codec_code_attrs_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlscSaasCodecCodeAttrsQueryResponse struct {

    // 接口返回model
    Result   *AlibabaAlscSaasCodecCodeAttrsQueryResult `json:"result,omitempty"`

}
