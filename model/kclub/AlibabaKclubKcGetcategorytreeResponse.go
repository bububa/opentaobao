package kclub

import (
    "github.com/bububa/opentaobao/model"
)

/* 
知识云-查询租户下类目树 APIResponse
alibaba.kclub.kc.getcategorytree

知识云-查询租户下类目树。通过租户id、类型(外部类目、帮助中心类目、内部类目)。
*/
type AlibabaKclubKcGetcategorytreeAPIResponse struct {
    model.CommonResponse
    Response *AlibabaKclubKcGetcategorytreeResponse `json:"alibaba_kclub_kc_getcategorytree_response,omitempty"`
}

type AlibabaKclubKcGetcategorytreeResponse struct {

    // result
    Result   *AlibabaKclubKcGetcategorytreeResult `json:"result,omitempty"`

}
