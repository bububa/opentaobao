package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
派样商品库存变更同步接口 APIResponse
alibaba.wdk.bm.paiyang.supply.publish

淘鲜达接入第三方进行派样，第三方同步大仓和门店的库存变更信息。
*/
type AlibabaWdkBmPaiyangSupplyPublishAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkBmPaiyangSupplyPublishResponse `json:"alibaba_wdk_bm_paiyang_supply_publish_response,omitempty"` 
    AlibabaWdkBmPaiyangSupplyPublishResponse
}

/* model for simplify = false
type AlibabaWdkBmPaiyangSupplyPublishResponse struct {

    // 请求出参
    
    Result  *struct {
        BmResult  *BmResult `json:"bm_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkBmPaiyangSupplyPublishResponse struct {

    // 请求出参
    Result   *BmResult `json:"result,omitempty"`

}
