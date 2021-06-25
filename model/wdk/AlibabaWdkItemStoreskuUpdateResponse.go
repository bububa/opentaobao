package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
五道口商品中心门店商品修改 APIResponse
alibaba.wdk.item.storesku.update

五道口商品中心门店商品修改
*/
type AlibabaWdkItemStoreskuUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkItemStoreskuUpdateResponse `json:"alibaba_wdk_item_storesku_update_response,omitempty"`
}

type AlibabaWdkItemStoreskuUpdateResponse struct {

    // result
    Result   *AlibabaWdkItemStoreskuUpdateResult `json:"result,omitempty"`

}
