package xiamicontent

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xiamicontent"
)

/* 
搜索歌曲列表 
xiami.content.songs.info.query

多维度查询歌曲列表
*/
func XiamiContentSongsInfoQuery(clt *core.SDKClient, req *xiamicontent.XiamiContentSongsInfoQueryAPIRequest, session string) (*xiamicontent.XiamiContentSongsInfoQueryAPIResponse, error) {
    var resp xiamicontent.XiamiContentSongsInfoQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
