package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家商品批量查询接口 API返回值 
alibaba.wdk.sku.merchantsku.scroll.query

提供主档商品数据接口查询
*/
type AlibabaWdkSkuMerchantskuScrollQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSkuMerchantskuScrollQueryAPIResponseModel
}

// 商家商品批量查询接口 成功返回结果
type AlibabaWdkSkuMerchantskuScrollQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_sku_merchantsku_scroll_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 请求结果对象
    Result   *ApiScrollPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
