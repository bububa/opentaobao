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
    Response *AlibabaAilabsAligenieAlbumsGetResponse `json:"alibaba_ailabs_aligenie_albums_get_response,omitempty"`
}

type AlibabaAilabsAligenieAlbumsGetResponse struct {

    // result
    Result   *AiCloudResult `json:"result,omitempty"`

}
