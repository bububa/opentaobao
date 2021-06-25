package media

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/media"
)

/* 
修改图片的分类 
taobao.picture.changecategory

把批量的图片移动到某个分类下
*/
func TaobaoPictureChangecategory(clt *core.SDKClient, req *media.TaobaoPictureChangecategoryRequest, session string) (*media.TaobaoPictureChangecategoryResponse, error) {
    var resp media.TaobaoPictureChangecategoryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
