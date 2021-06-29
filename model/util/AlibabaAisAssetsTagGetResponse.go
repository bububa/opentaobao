package util

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
基础设施资产标签获取 APIResponse
alibaba.ais.assets.tag.get

提供浪潮，英业达等厂商供应阿里巴巴基础设施资产的标签QR code获取
*/
type AlibabaAisAssetsTagGetAPIResponse struct {
    model.CommonResponse
    AlibabaAisAssetsTagGetResponse
}

type AlibabaAisAssetsTagGetResponse struct {
    XMLName xml.Name `xml:"alibaba_ais_assets_tag_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 最外层结果
    
    Result   *BaseRep `json:"result,omitempty" xml:"result,omitempty"`

    
}
