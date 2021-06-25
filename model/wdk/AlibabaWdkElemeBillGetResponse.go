package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
饿了么日维度对账单查询 APIResponse
alibaba.wdk.eleme.bill.get

查询饿了么日维度对账单信息
*/
type AlibabaWdkElemeBillGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkElemeBillGetResponse `json:"alibaba_wdk_eleme_bill_get_response,omitempty"`
}

type AlibabaWdkElemeBillGetResponse struct {

    // 返回结果
    Result   *AlibabaWdkElemeBillGetApiResult `json:"result,omitempty"`

}
