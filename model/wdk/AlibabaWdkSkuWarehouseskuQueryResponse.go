package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
仓商品查询接口(指定商品编码) APIResponse
alibaba.wdk.sku.warehousesku.query

提供指定仓商品编码查询
*/
type AlibabaWdkSkuWarehouseskuQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkSkuWarehouseskuQueryResponse `json:"alibaba_wdk_sku_warehousesku_query_response,omitempty"`
}

type AlibabaWdkSkuWarehouseskuQueryResponse struct {

    // 返回结果
    Result   *AlibabaWdkSkuWarehouseskuQueryApiResult `json:"result,omitempty"`

}
