package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
仓商品遍历接口(游标) APIResponse
alibaba.wdk.sku.warehousesku.scroll.query

提供仓商品数据接口查询
*/
type AlibabaWdkSkuWarehouseskuScrollQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSkuWarehouseskuScrollQueryResponse
}

type AlibabaWdkSkuWarehouseskuScrollQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_sku_warehousesku_scroll_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *ApiScrollPageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
