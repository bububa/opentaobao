package media

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/media"
)

/* 
替换图片 
taobao.picture.replace

替换一张图片，只替换图片数据，图片名称，图片分类等保持不变。
*/
func TaobaoPictureReplace(clt *core.SDKClient, req *media.TaobaoPictureReplaceRequest, session string) (*media.TaobaoPictureReplaceAPIResponse, error) {
    var resp media.TaobaoPictureReplaceAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
