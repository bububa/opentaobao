package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店查询接口 API返回值 
alibaba.wdk.shop.query

根据门店code查询门店信息
*/
type AlibabaWdkShopQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkShopQueryResponse
}

// 门店查询接口 成功返回结果
type AlibabaWdkShopQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_shop_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 调用结果
    Result   *AlibabaWdkShopQueryApiResults `json:"result,omitempty" xml:"result,omitempty"`
}
