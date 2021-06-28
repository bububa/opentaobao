package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
专辑详情 APIResponse
alibaba.ailabs.aligenie.albums.get

给予厂商查询专辑下的音频详情
*/
type AlibabaAilabsAligenieAlbumsGetAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsAligenieAlbumsGetResponse
}

type AlibabaAilabsAligenieAlbumsGetResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_aligenie_albums_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
