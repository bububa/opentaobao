package aliyun

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliyun"
)

/* 
配置LoadBalancer的别名。 
slb.aliyuncs.com.SetLoadBalancerName.2013-02-21

配置LoadBalancer的别名。
*/
func SlbAliyuncsComSetLoadBalancerName2013-02-21(clt *core.SDKClient, req *aliyun.SlbAliyuncsComSetLoadBalancerName2013-02-21Request, session string) (*aliyun.SlbAliyuncsComSetLoadBalancerName2013-02-21Response, error) {
    var resp aliyun.SlbAliyuncsComSetLoadBalancerName2013-02-21APIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
