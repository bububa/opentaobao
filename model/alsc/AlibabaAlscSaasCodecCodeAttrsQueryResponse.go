package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
码业务属性查询 APIResponse
alibaba.alsc.saas.codec.code.attrs.query

码业务属性查询
*/
type AlibabaAlscSaasCodecCodeAttrsQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alsc_saas_codec_code_attrs_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回model
    
    Result   *AlibabaAlscSaasCodecCodeAttrsQueryResult `json:"result,omitempty" xml:"