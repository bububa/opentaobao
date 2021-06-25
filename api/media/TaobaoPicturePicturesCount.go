package media

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/media"
)

/* 
图片总数查询 
taobao.picture.pictures.count

图片总数查询
*/
func TaobaoPicturePicturesCount(clt *core.SDKClient, req *media.TaobaoPicturePicturesCountRequest, session string) (*media.TaobaoPicturePicturesCountResponse, error) {
    var resp media.TaobaoPicturePicturesCountAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
