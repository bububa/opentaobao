package media

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/media"
)

/* 
查询某个用户图片空间的使用情况 
taobao.vas.service.getServTimes

查询某个用户图片空间的使用情况
*/
func TaobaoVasServiceGetServTimes(clt *core.SDKClient, req *media.TaobaoVasServiceGetServTimesRequest, session string) (*media.TaobaoVasServiceGetServTimesResponse, error) {
    var resp media.TaobaoVasServiceGetServTimesAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
