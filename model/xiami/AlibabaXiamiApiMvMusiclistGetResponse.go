package xiami

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
乐馆mv列表 APIResponse
alibaba.xiami.api.mv.musiclist.get

乐馆mv列表
*/
type AlibabaXiamiApiMvMusiclistGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_xiami_api_mv_musiclist_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // mv_music_list_result
    
    MvMusicListResult   *AlibabaXiamiApiMvMusiclistGetStruct `json:"mv_music_list_result,omitempty" xml:"