package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
共享库存订单投保消息获取 APIResponse
alibaba.wdkorder.sharestock.insurance.getorder

共享库存订单投保消息获取
*/
type AlibabaWdkorderSharestockInsuranceGetorderAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkorderSharestockInsuranceGetorderResponse `json:"alibaba_wdkorder_sharestock_insurance_getorder_response,omitempty"`
}

type AlibabaWdkorderSharestockInsuranceGetorderResponse struct {

    // 返回结果
    Result   *MaochaoOrderInsuranceQueryResult `json:"result,omitempty"`

}
