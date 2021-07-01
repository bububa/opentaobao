package miniappcloud

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/miniappcloud"
)

/* 
更新MongoDB中的数据 
taobao.miniapp.cloud.mongo.update

更新MongoDB中的数据
*/
func TaobaoMiniappCloudMongoUpdate(clt *core.SDKClient, req *miniappcloud.TaobaoMiniappCloudMongoUpdateAPIRequest, session string) (*miniappcloud.TaobaoMiniappCloudMongoUpdateAPIResponse, error) {
    var resp miniappcloud.TaobaoMiniappCloudMongoUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
