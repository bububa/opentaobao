package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
组合商品新增接口 APIResponse
alibaba.wdk.sku.combinesku.add

组合商品新增接口
*/
type AlibabaWdkSkuCombineskuAddAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkSkuCombineskuAddResponse `json:"alibaba_wdk_sku_combinesku_add_response,omitempty"`
}

type AlibabaWdkSkuCombineskuAddResponse struct {

    // 调用结果
    Result   *AlibabaWdkSkuCombineskuAddApiResults `json:"result,omitempty"`

}
