package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商户订单履约数据获取 APIResponse
alibaba.wdkorder.sharestock.fulfill.get

商户订单履约数据获取
*/
type AlibabaWdkorderSharestockFulfillGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkorderSharestockFulfillGetResponse `json:"alibaba_wdkorder_sharestock_fulfill_get_response,omitempty"` 
    AlibabaWdkorderSharestockFulfillGetResponse
}

/* model for simplify = false
type AlibabaWdkorderSharestockFulfillGetResponse struct {

    // 调用结果
    
    Result  *struct {
        MaochaoOrderFulfillQueryResult  *MaochaoOrderFulfillQueryResult `json:"maochao_order_fulfill_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkorderSharestockFulfillGetResponse struct {

    // 调用结果
    Result   *MaochaoOrderFulfillQueryResult `json:"result,omitempty"`

}
