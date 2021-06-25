package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询音频 APIResponse
alibaba.ailabs.aligenie.tracks.search

搜索类目下的音频信息
*/
type AlibabaAilabsAligenieTracksSearchAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAilabsAligenieTracksSearchResponse `json:"alibaba_ailabs_aligenie_tracks_search_response,omitempty"`
}

type AlibabaAilabsAligenieTracksSearchResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
