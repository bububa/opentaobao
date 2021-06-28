package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
热门歌曲 APIResponse
alibaba.xiami.api.artist.hotSongs.get

热门歌曲
*/
type AlibabaXiamiApiArtistHotSongsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_xiami_api_artist_hotSongs_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Data   *HotSongsResult `json:"data,omitempty" xml:"