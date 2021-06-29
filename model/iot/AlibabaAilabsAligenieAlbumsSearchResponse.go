package iot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询专辑 API返回值 
alibaba.ailabs.aligenie.albums.search

搜索类目下的专辑信息
*/
type AlibabaAilabsAligenieAlbumsSearchAPIResponse struct {
    model.CommonResponse
    AlibabaAilabsAligenieAlbumsSearchResponse
}

// 查询专辑 成功返回结果
type AlibabaAilabsAligenieAlbumsSearchResponse struct {
    XMLName xml.Name `xml:"alibaba_ailabs_aligenie_albums_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AiCloudResult `json:"result,omitempty" xml:"result,omitempty"`
}
