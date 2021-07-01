package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
仓商品查询接口(指定商品编码) API返回值 
alibaba.wdk.sku.warehousesku.query

提供指定仓商品编码查询
*/
type AlibabaWdkSkuWarehouseskuQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSkuWarehouseskuQueryAPIResponseModel
}

// 仓商品查询接口(指定商品编码) 成功返回结果
type AlibabaWdkSkuWarehouseskuQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_sku_warehousesku_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *AlibabaWdkSkuWarehouseskuQueryApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
