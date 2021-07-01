package xiami

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xiami"
)

/* 
虾米音乐－风格，流派专辑列表 
alibaba.xiami.api.tag.genre.album.get

虾米音乐－风格，流派专辑列表
*/
func AlibabaXiamiApiTagGenreAlbumGet(clt *core.SDKClient, req *xiami.AlibabaXiamiApiTagGenreAlbumGetAPIRequest, session string) (*xiami.AlibabaXiamiApiTagGenreAlbumGetAPIResponse, error) {
    var resp xiami.AlibabaXiamiApiTagGenreAlbumGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
