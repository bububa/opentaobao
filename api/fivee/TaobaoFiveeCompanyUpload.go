package fivee

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/fivee"
)

/* 
上传商信息接口 
taobao.fivee.company.upload

资质共享平台上传资质证照
*/
func TaobaoFiveeCompanyUpload(clt *core.SDKClient, req *fivee.TaobaoFiveeCompanyUploadAPIRequest, session string) (*fivee.TaobaoFiveeCompanyUploadAPIResponse, error) {
    var resp fivee.TaobaoFiveeCompanyUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
