package xiami

import (
    "github.com/bububa/opentaobao/model"
)

/* 
艺人详情 APIResponse
alibaba.xiami.api.artist.detail.get

艺人详情
*/
type AlibabaXiamiApiArtistDetailGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaXiamiApiArtistDetailGetResponse `json:"alibaba_xiami_api_artist_detail_get_response,omitempty"` 
    AlibabaXiamiApiArtistDetailGetResponse
}

/* model for simplify = false
type AlibabaXiamiApiArtistDetailGetResponse struct {

    // 返回数据
    
    Data  *struct {
        StandardArtist  *StandardArtist `json:"standard_artist,omitempty"`
    } `json:"data,omitempty"`
    

}
*/

type AlibabaXiamiApiArtistDetailGetResponse struct {

    // 返回数据
    Data   *StandardArtist `json:"data,omitempty"`

}
