package miniappcloud

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/miniappcloud"
)

/* 
云存储根据文件名反查地址 
taobao.miniapp.cloud.store.listfile

云存储中，根据文件名反查地址
*/
func TaobaoMiniappCloudStoreListfile(clt *core.SDKClient, req *miniappcloud.TaobaoMiniappCloudStoreListfileRequest, session string) (*miniappcloud.TaobaoMiniappCloudStoreListfileAPIResponse, error) {
    var resp miniappcloud.TaobaoMiniappCloudStoreListfileAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
