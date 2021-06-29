package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询音频 APIResponse
alibaba.ailabs.aligenie.tracks.search

搜索类目下的音频信息
*/
type AlibabaAilabsAligenieTracksSearchAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsAligenieTracksSearchResponse
}

type AlibabaAilabsAligenieTracksSearchResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_aligenie_tracks_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
