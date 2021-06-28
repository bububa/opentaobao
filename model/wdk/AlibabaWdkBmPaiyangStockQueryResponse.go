package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
派样商品门店库存查询接口 APIResponse
alibaba.wdk.bm.paiyang.stock.query

淘鲜达接入第三方进行派样，第三方查询派样商品的门店库存信息。
*/
type AlibabaWdkBmPaiyangStockQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkBmPaiyangStockQueryResponse `json:"alibaba_wdk_bm_paiyang_stock_query_response,omitempty"` 
    AlibabaWdkBmPaiyangStockQueryResponse
}

/* model for simplify = false
type AlibabaWdkBmPaiyangStockQueryResponse struct {

    // 请求出参
    
    Result  *struct {
        BmResult  *BmResult `json:"bm_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkBmPaiyangStockQueryResponse struct {

    // 请求出参
    Result   *BmResult `json:"result,omitempty"`

}
