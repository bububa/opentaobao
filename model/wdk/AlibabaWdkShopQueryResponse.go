package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
门店查询接口 APIResponse
alibaba.wdk.shop.query

根据门店code查询门店信息
*/
type AlibabaWdkShopQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_shop_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用结果
    
    Result   *AlibabaWdkShopQueryApiResults `json:"result,omitempty" xml:"