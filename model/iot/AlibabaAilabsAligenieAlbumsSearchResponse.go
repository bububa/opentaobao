package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询专辑 APIResponse
alibaba.ailabs.aligenie.albums.search

搜索类目下的专辑信息
*/
type AlibabaAilabsAligenieAlbumsSearchAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsAligenieAlbumsSearchResponse
}

type AlibabaAilabsAligenieAlbumsSearchResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_aligenie_albums_search_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
