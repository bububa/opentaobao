package xiami

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取电台歌曲 APIResponse
alibaba.xiami.api.radio.songs.get

获取电台songs
*/
type AlibabaXiamiApiRadioSongsGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaXiamiApiRadioSongsGetResponse `json:"alibaba_xiami_api_radio_songs_get_response,omitempty"`
}

type AlibabaXiamiApiRadioSongsGetResponse struct {

    // 歌曲列表
    Data   []StandardSong `json:"data,omitempty"`

}
