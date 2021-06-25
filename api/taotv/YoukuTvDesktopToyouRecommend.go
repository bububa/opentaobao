package taotv

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/taotv"
)

/* 
TV桌面为你推荐接口 
youku.tv.desktop.toyou.recommend

提供为你推荐数据
*/
func YoukuTvDesktopToyouRecommend(clt *core.SDKClient, req *taotv.YoukuTvDesktopToyouRecommendRequest, session string) (*taotv.YoukuTvDesktopToyouRecommendResponse, error) {
    var resp taotv.YoukuTvDesktopToyouRecommendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
