package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
仓商品遍历接口(游标) API返回值 
alibaba.wdk.sku.warehousesku.scroll.query

提供仓商品数据接口查询
*/
type AlibabaWdkSkuWarehouseskuScrollQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSkuWarehouseskuScrollQueryAPIResponseModel
}

// 仓商品遍历接口(游标) 成功返回结果
type AlibabaWdkSkuWarehouseskuScrollQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_sku_warehousesku_scroll_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *ApiScrollPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
