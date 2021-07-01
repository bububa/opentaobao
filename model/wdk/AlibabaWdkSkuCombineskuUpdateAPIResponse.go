package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组合商品更新接口 API返回值 
alibaba.wdk.sku.combinesku.update

组合商品修改接口
*/
type AlibabaWdkSkuCombineskuUpdateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSkuCombineskuUpdateAPIResponseModel
}

// 组合商品更新接口 成功返回结果
type AlibabaWdkSkuCombineskuUpdateAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_wdk_sku_combinesku_update_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 调用结果
    Result   *AlibabaWdkSkuCombineskuUpdateApiResults `json:"result,omitempty" xml:"result,omitempty"`
}
