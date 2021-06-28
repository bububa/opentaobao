package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
基础设施资产标签生成 APIResponse
alibaba.ais.assets.tag.generate

提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code生成
*/
type AlibabaAisAssetsTagGenerateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAisAssetsTagGenerateResponse `json:"alibaba_ais_assets_tag_generate_response,omitempty"` 
    AlibabaAisAssetsTagGenerateResponse
}

/* model for simplify = false
type AlibabaAisAssetsTagGenerateResponse struct {

    // 最外层结果
    
    Result  *struct {
        BaseRep  *BaseRep `json:"base_rep,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAisAssetsTagGenerateResponse struct {

    // 最外层结果
    Result   *BaseRep `json:"result,omitempty"`

}
