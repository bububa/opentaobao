package media

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/media"
)

/* 
获取商家视频列表 
taobao.media.video.list

用于获取授权商家的视频列表
*/
func TaobaoMediaVideoList(clt *core.SDKClient, req *media.TaobaoMediaVideoListRequest, session string) (*media.TaobaoMediaVideoListAPIResponse, error) {
    var resp media.TaobaoMediaVideoListAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
