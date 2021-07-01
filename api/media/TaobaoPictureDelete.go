package media

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/media"
)

/* 
删除图片空间图片 
taobao.picture.delete

删除图片空间图片
*/
func TaobaoPictureDelete(clt *core.SDKClient, req *media.TaobaoPictureDeleteAPIRequest, session string) (*media.TaobaoPictureDeleteAPIResponse, error) {
    var resp media.TaobaoPictureDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
