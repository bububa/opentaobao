package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
品牌营销涉及到的商品的库存同步接口 APIResponse
alibaba.wdk.bm.stock.publish

用于操作sku的库存
*/
type AlibabaWdkBmStockPublishAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkBmStockPublishResponse `json:"alibaba_wdk_bm_stock_publish_response,omitempty"`
}

type AlibabaWdkBmStockPublishResponse struct {

    // 出参
    Result   *BmResult `json:"result,omitempty"`

}
