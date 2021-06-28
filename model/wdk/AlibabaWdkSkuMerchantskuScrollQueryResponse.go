package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家商品批量查询接口 APIResponse
alibaba.wdk.sku.merchantsku.scroll.query

提供主档商品数据接口查询
*/
type AlibabaWdkSkuMerchantskuScrollQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_sku_merchantsku_scroll_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 请求结果对象
    
    Result   *ApiScrollPageResult `json:"result,omitempty" xml:"