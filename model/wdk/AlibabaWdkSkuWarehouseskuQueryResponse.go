package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
仓商品查询接口(指定商品编码) APIResponse
alibaba.wdk.sku.warehousesku.query

提供指定仓商品编码查询
*/
type AlibabaWdkSkuWarehouseskuQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_sku_warehousesku_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *AlibabaWdkSkuWarehouseskuQueryApiResult `json:"result,omitempty" xml:"