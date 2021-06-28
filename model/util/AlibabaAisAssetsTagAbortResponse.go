package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
基础设施资产标签废弃 APIResponse
alibaba.ais.assets.tag.abort

提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code未使用的废弃
*/
type AlibabaAisAssetsTagAbortAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAisAssetsTagAbortResponse `json:"alibaba_ais_assets_tag_abort_response,omitempty"` 
    AlibabaAisAssetsTagAbortResponse
}

/* model for simplify = false
type AlibabaAisAssetsTagAbortResponse struct {

    // 最外层结果
    
    Result  *struct {
        BaseRep  *BaseRep `json:"base_rep,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAisAssetsTagAbortResponse struct {

    // 最外层结果
    Result   *BaseRep `json:"result,omitempty"`

}
