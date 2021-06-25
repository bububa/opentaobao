package xiami

import (
    "github.com/bububa/opentaobao/model"
)

/* 
乐馆mv列表 APIResponse
alibaba.xiami.api.mv.musiclist.get

乐馆mv列表
*/
type AlibabaXiamiApiMvMusiclistGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaXiamiApiMvMusiclistGetResponse `json:"alibaba_xiami_api_mv_musiclist_get_response,omitempty"`
}

type AlibabaXiamiApiMvMusiclistGetResponse struct {

    // mv_music_list_result
    MvMusicListResult   *AlibabaXiamiApiMvMusiclistGetStruct `json:"mv_music_list_result,omitempty"`

}
