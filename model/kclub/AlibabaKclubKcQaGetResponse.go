package kclub

import (
    "github.com/bububa/opentaobao/model"
)

/* 
知识云-查询单个知识详情 APIResponse
alibaba.kclub.kc.qa.get

知识云-查询单个知识详情。通过租户id、问题id查询问题详情
*/
type AlibabaKclubKcQaGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaKclubKcQaGetResponse `json:"alibaba_kclub_kc_qa_get_response,omitempty"`
}

type AlibabaKclubKcQaGetResponse struct {

    // 返回结果
    Result   *AlibabaKclubKcQaGetResult `json:"result,omitempty"`

}
