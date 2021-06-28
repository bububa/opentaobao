package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品池阶梯商品添加 APIResponse
alibaba.wdk.marketing.itempool.stair.additem

添加商品池阶梯商品
*/
type AlibabaWdkMarketingItempoolStairAdditemAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_itempool_stair_additem_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 添加商品返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"