package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
基础设施资产标签获取 APIResponse
alibaba.ais.assets.tag.get

提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code获取
*/
type AlibabaAisAssetsTagGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAisAssetsTagGetResponse `json:"alibaba_ais_assets_tag_get_response,omitempty"`
}

type AlibabaAisAssetsTagGetResponse struct {

    // 最外层结果
    Result   *BaseRep `json:"result,omitempty"`

}
