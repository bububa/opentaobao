package iot

import (
    "github.com/bububa/opentaobao/model"
)

/* 
专辑详情 APIResponse
alibaba.ailabs.aligenie.albums.get

给予厂商查询专辑下的音频详情
*/
type AlibabaAilabsAligenieAlbumsGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAilabsAligenieAlbumsGetResponse `json:"alibaba_ailabs_aligenie_albums_get_response,omitempty"` 
    AlibabaAilabsAligenieAlbumsGetResponse
}

/* model for simplify = false
type AlibabaAilabsAligenieAlbumsGetResponse struct {

    // result
    
    Result  *struct {
        AiCloudResult  *AiCloudResult `json:"ai_cloud_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAilabsAligenieAlbumsGetResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
