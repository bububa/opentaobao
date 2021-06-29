package xiamicontent

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xiamicontent"
)

/* 
获取副歌信息 
xiami.content.songs.audio.getrefrain

获取歌曲音频副歌
*/
func XiamiContentSongsAudioGetrefrain(clt *core.SDKClient, req *xiamicontent.XiamiContentSongsAudioGetrefrainRequest, session string) (*xiamicontent.XiamiContentSongsAudioGetrefrainAPIResponse, error) {
    var resp xiamicontent.XiamiContentSongsAudioGetrefrainAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
