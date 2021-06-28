package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
基础设施资产标签生成 APIResponse
alibaba.ais.assets.tag.generate

提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code生成
*/
type AlibabaAisAssetsTagGenerateAPIResponse struct {
    model.CommonResponse
    AlibabaAisAssetsTagGenerateResponse
}

type AlibabaAisAssetsTagGenerateResponse struct {
    XMLName xml.Name `xml:"alibaba_ais_assets_tag_generate_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 最外层结果
    
    Result   *BaseRep `json:"result,omitempty" xml:"result,omitempty"`

    
}
