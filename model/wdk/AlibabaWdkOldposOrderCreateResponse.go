package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘鲜达外部商户老pos机产生的订单同步进淘鲜达 APIResponse
alibaba.wdk.oldpos.order.create

淘鲜达外部商户老pos机产生的订单同步进淘鲜达
*/
type AlibabaWdkOldposOrderCreateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkOldposOrderCreateResponse `json:"alibaba_wdk_oldpos_order_create_response,omitempty"` 
    AlibabaWdkOldposOrderCreateResponse
}

/* model for simplify = false
type AlibabaWdkOldposOrderCreateResponse struct {

    // 结果
    
    Result  *struct {
        PosOrderCreateResult  *PosOrderCreateResult `json:"pos_order_create_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkOldposOrderCreateResponse struct {

    // 结果
    Result   *PosOrderCreateResult `json:"result,omitempty"`

}
