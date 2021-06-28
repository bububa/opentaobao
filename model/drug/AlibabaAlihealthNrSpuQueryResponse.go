package drug

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取标品库标品信息 APIResponse
alibaba.alihealth.nr.spu.query

提供给ERP使用的，获取健康标品库信息
*/
type AlibabaAlihealthNrSpuQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alihealth_nr_spu_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *ResponseResult `json:"result,omitempty" xml:"