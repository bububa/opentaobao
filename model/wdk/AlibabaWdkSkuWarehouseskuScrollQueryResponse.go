package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
仓商品遍历接口(游标) APIResponse
alibaba.wdk.sku.warehousesku.scroll.query

提供仓商品数据接口查询
*/
type AlibabaWdkSkuWarehouseskuScrollQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkSkuWarehouseskuScrollQueryResponse `json:"alibaba_wdk_sku_warehousesku_scroll_query_response,omitempty"`
}

type AlibabaWdkSkuWarehouseskuScrollQueryResponse struct {

    // 返回结果
    Result   *ApiScrollPageResult `json:"result,omitempty"`

}