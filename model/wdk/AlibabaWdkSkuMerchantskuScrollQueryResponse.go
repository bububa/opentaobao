package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家商品批量查询接口 APIResponse
alibaba.wdk.sku.merchantsku.scroll.query

提供主档商品数据接口查询
*/
type AlibabaWdkSkuMerchantskuScrollQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkSkuMerchantskuScrollQueryResponse `json:"alibaba_wdk_sku_merchantsku_scroll_query_response,omitempty"`
}

type AlibabaWdkSkuMerchantskuScrollQueryResponse struct {

    // 请求结果对象
    Result   *ApiScrollPageResult `json:"result,omitempty"`

}
