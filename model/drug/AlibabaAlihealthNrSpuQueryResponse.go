package drug

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取标品库标品信息 APIResponse
alibaba.alihealth.nr.spu.query

提供给ERP使用的，获取健康标品库信息
*/
type AlibabaAlihealthNrSpuQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlihealthNrSpuQueryResponse `json:"alibaba_alihealth_nr_spu_query_response,omitempty"`
}

type AlibabaAlihealthNrSpuQueryResponse struct {

    // 结果
    Result   *ResponseResult `json:"result,omitempty"`

}
