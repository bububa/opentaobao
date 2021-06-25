package media

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/media"
)

/* 
获取图片分类信息 
taobao.picture.category.get

获取图片分类信息
*/
func TaobaoPictureCategoryGet(clt *core.SDKClient, req *media.TaobaoPictureCategoryGetRequest, session string) (*media.TaobaoPictureCategoryGetResponse, error) {
    var resp media.TaobaoPictureCategoryGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
