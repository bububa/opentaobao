package nazca

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/nazca"
)

/* 
获取文件秘钥 
alibaba.nazca.token.filesecret.get

获取文件秘钥
*/
func AlibabaNazcaTokenFilesecretGet(clt *core.SDKClient, req *nazca.AlibabaNazcaTokenFilesecretGetAPIRequest, session string) (*nazca.AlibabaNazcaTokenFilesecretGetAPIResponse, error) {
    var resp nazca.AlibabaNazcaTokenFilesecretGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
