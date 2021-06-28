package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
共享库存订单投保后回传保单号 APIResponse
alibaba.wdkorder.sharestock.insurance.callback

共享库存订单投保消息获取
*/
type AlibabaWdkorderSharestockInsuranceCallbackAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkorderSharestockInsuranceCallbackResponse `json:"alibaba_wdkorder_sharestock_insurance_callback_response,omitempty"` 
    AlibabaWdkorderSharestockInsuranceCallbackResponse
}

/* model for simplify = false
type AlibabaWdkorderSharestockInsuranceCallbackResponse struct {

    // 系统自动生成
    
    Result  *struct {
        MaochaoOrderInsuranceCallbackResult  *MaochaoOrderInsuranceCallbackResult `json:"maochao_order_insurance_callback_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkorderSharestockInsuranceCallbackResponse struct {

    // 系统自动生成
    Result   *MaochaoOrderInsuranceCallbackResult `json:"result,omitempty"`

}
