package xiamiopen

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xiamiopen"
)

/* 
获取歌曲试听文件 
xiami.api.song.listenfile.get

获取歌曲试听文件
*/
func XiamiApiSongListenfileGet(clt *core.SDKClient, req *xiamiopen.XiamiApiSongListenfileGetRequest, session string) (*xiamiopen.XiamiApiSongListenfileGetAPIResponse, error) {
    var resp xiamiopen.XiamiApiSongListenfileGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
