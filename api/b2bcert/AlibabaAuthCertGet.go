package b2bcert

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/b2bcert"
)

/* 
获取证书数据 
alibaba.auth.cert.get

获取证书数据
*/
func AlibabaAuthCertGet(clt *core.SDKClient, req *b2bcert.AlibabaAuthCertGetAPIRequest, session string) (*b2bcert.AlibabaAuthCertGetAPIResponse, error) {
    var resp b2bcert.AlibabaAuthCertGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
