package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新商品 APIResponse
alibaba.wdk.sku.update

开放商品更新接口
*/
type AlibabaWdkSkuUpdateAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkSkuUpdateResponse `json:"alibaba_wdk_sku_update_response,omitempty"`
}

type AlibabaWdkSkuUpdateResponse struct {

    // 执行结果
    Result   *AlibabaWdkSkuUpdateApiResults `json:"result,omitempty"`

}
