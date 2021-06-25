package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新渠道商品 APIResponse
alibaba.wdk.sku.channelsku.update

批量更新渠道商品，商家通过Top接入
*/
type AlibabaWdkSkuChannelskuUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkSkuChannelskuUpdateResponse `json:"alibaba_wdk_sku_channelsku_update_response,omitempty"`
}

type AlibabaWdkSkuChannelskuUpdateResponse struct {

    // 调用结果
    Result   *AlibabaWdkSkuChannelskuUpdateApiResults `json:"result,omitempty"`

}
