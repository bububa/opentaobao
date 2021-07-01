package media

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/media"
)

/* 
图片获取 
taobao.picture.pictures.get

图片空间对外的图片获取接口，该接口只针对分页获取，获取某一页的图片，该接口不支持总数的查询asd
*/
func TaobaoPicturePicturesGet(clt *core.SDKClient, req *media.TaobaoPicturePicturesGetAPIRequest, session string) (*media.TaobaoPicturePicturesGetAPIResponse, error) {
    var resp media.TaobaoPicturePicturesGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
