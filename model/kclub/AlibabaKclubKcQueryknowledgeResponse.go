package kclub

import (
    "github.com/bububa/opentaobao/model"
)

/* 
知识云-通用知识查询服务 APIResponse
alibaba.kclub.kc.queryknowledge

知识云-通用知识查询服务。通过租户id、类目id、知识类型、知识状态等条件查询类目。
*/
type AlibabaKclubKcQueryknowledgeAPIResponse struct {
    model.CommonResponse
    Response *AlibabaKclubKcQueryknowledgeResponse `json:"alibaba_kclub_kc_queryknowledge_response,omitempty"`
}

type AlibabaKclubKcQueryknowledgeResponse struct {

    // 结果
    Result   *AlibabaKclubKcQueryknowledgeResult `json:"result,omitempty"`

}
