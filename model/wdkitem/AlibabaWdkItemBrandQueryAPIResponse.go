package wdkitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌信息查询 API返回值 
alibaba.wdk.item.brand.query

品牌信息查询
*/
type AlibabaWdkItemBrandQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkItemBrandQueryAPIResponseModel
}

// 品牌信息查询 成功返回结果
type AlibabaWdkItemBrandQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_item_brand_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AlibabaWdkItemBrandQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
