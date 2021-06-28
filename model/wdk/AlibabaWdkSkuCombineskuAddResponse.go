package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组合商品新增接口 APIResponse
alibaba.wdk.sku.combinesku.add

组合商品新增接口
*/
type AlibabaWdkSkuCombineskuAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_sku_combinesku_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用结果
    
    Result   *AlibabaWdkSkuCombineskuAddApiResults `json:"result,omitempty" xml:"