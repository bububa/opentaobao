package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
派样商品库存变更同步接口 APIResponse
alibaba.wdk.bm.paiyang.supply.publish

淘鲜达接入第三方进行派样，第三方同步大仓和门店的库存变更信息。
*/
type AlibabaWdkBmPaiyangSupplyPublishAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_bm_paiyang_supply_publish_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 请求出参
    
    Result   *BmResult `json:"result,omitempty" xml:"