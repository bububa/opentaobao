package xiami

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xiami"
)

/* 
获取电台歌曲 
alibaba.xiami.api.radio.songs.get

获取电台songs
*/
func AlibabaXiamiApiRadioSongsGet(clt *core.SDKClient, req *xiami.AlibabaXiamiApiRadioSongsGetAPIRequest, session string) (*xiami.AlibabaXiamiApiRadioSongsGetAPIResponse, error) {
    var resp xiami.AlibabaXiamiApiRadioSongsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
