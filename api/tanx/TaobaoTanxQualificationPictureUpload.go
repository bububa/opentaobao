package tanx

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tanx"
)

/* 
资质图片上传接口 
taobao.tanx.qualification.picture.upload

资质图片上传接口
*/
func TaobaoTanxQualificationPictureUpload(clt *core.SDKClient, req *tanx.TaobaoTanxQualificationPictureUploadRequest, session string) (*tanx.TaobaoTanxQualificationPictureUploadResponse, error) {
    var resp tanx.TaobaoTanxQualificationPictureUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
