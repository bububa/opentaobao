package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询专辑 APIResponse
alibaba.ailabs.aligenie.albums.search

搜索类目下的专辑信息
*/
type AlibabaAilabsAligenieAlbumsSearchAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAilabsAligenieAlbumsSearchResponse `json:"alibaba_ailabs_aligenie_albums_search_response,omitempty"` 
    AlibabaAilabsAligenieAlbumsSearchResponse
}

/* model for simplify = false
type AlibabaAilabsAligenieAlbumsSearchResponse struct {

    // result
    
    Result  *struct {
        AiCloudResult  *AiCloudResult `json:"ai_cloud_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAilabsAligenieAlbumsSearchResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
