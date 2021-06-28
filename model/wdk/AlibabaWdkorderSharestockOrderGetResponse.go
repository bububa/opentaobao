package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
猫超商户订单拉取 APIResponse
alibaba.wdkorder.sharestock.order.get

商户拉取猫超订单数据
*/
type AlibabaWdkorderSharestockOrderGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkorderSharestockOrderGetResponse `json:"alibaba_wdkorder_sharestock_order_get_response,omitempty"` 
    AlibabaWdkorderSharestockOrderGetResponse
}

/* model for simplify = false
type AlibabaWdkorderSharestockOrderGetResponse struct {

    // 调用结果
    
    Result  *struct {
        MaochaoOrderQueryResult  *MaochaoOrderQueryResult `json:"maochao_order_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkorderSharestockOrderGetResponse struct {

    // 调用结果
    Result   *MaochaoOrderQueryResult `json:"result,omitempty"`

}
