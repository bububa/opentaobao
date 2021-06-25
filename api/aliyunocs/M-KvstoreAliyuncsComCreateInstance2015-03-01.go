package aliyunocs

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliyunocs"
)

/* 
创建OCS实例 
m-kvstore.aliyuncs.com.CreateInstance.2015-03-01

创建OCS实例
*/
func M-KvstoreAliyuncsComCreateInstance2015-03-01(clt *core.SDKClient, req *aliyunocs.M-KvstoreAliyuncsComCreateInstance2015-03-01Request, session string) (*aliyunocs.M-KvstoreAliyuncsComCreateInstance2015-03-01Response, error) {
    var resp aliyunocs.M-KvstoreAliyuncsComCreateInstance2015-03-01APIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
