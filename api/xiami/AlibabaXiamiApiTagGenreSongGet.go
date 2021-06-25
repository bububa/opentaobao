package xiami

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xiami"
)

/* 
虾米-风格，流派歌曲 
alibaba.xiami.api.tag.genre.song.get

虾米-风格，流派歌曲
*/
func AlibabaXiamiApiTagGenreSongGet(clt *core.SDKClient, req *xiami.AlibabaXiamiApiTagGenreSongGetRequest, session string) (*xiami.AlibabaXiamiApiTagGenreSongGetResponse, error) {
    var resp xiami.AlibabaXiamiApiTagGenreSongGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
