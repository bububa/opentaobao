package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
添加商品池活动 APIResponse
alibaba.wdk.marketing.itempool.createactivity

添加商品池活动
*/
type AlibabaWdkMarketingItempoolCreateactivityAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_itempool_createactivity_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 创建活动返回结果
    
    Result   *MarketResult `json:"result,omitempty" xml:"