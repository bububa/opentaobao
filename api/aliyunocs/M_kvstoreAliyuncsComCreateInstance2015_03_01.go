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
func M_kvstoreAliyuncsComCreateInstance2015_03_01(clt *core.SDKClient, req *aliyunocs.M_kvstoreAliyuncsComCreateInstance2015_03_01APIRequest, session string) (*aliyunocs.M_kvstoreAliyuncsComCreateInstance2015_03_01APIResponse, error) {
    var resp aliyunocs.M_kvstoreAliyuncsComCreateInstance2015_03_01APIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
