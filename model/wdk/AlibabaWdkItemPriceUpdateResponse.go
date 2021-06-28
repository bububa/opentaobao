package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口商品中心价格修改 APIResponse
alibaba.wdk.item.price.update

修改门店商品售价和会员价
*/
type AlibabaWdkItemPriceUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_item_price_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // hsfResult
    
    HsfResult   *AlibabaWdkItemPriceUpdateResult `json:"hsf_result,omitempty" xml:"