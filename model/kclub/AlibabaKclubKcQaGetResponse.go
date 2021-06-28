package kclub

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
知识云-查询单个知识详情 APIResponse
alibaba.kclub.kc.qa.get

知识云-查询单个知识详情。通过租户id、问题id查询问题详情
*/
type AlibabaKclubKcQaGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_kclub_kc_qa_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *AlibabaKclubKcQaGetResult `json:"result,omitempty" xml:"