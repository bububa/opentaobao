package aliyunocs

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aliyunocs"
)

/* 
查看Region列表 
m-kvstore.aliyuncs.com.DescribeRegions.2015-03-01

查看Region列表
*/
func M-KvstoreAliyuncsComDescribeRegions2015-03-01(clt *core.SDKClient, req *aliyunocs.M-KvstoreAliyuncsComDescribeRegions2015-03-01Request, session string) (*aliyunocs.M-KvstoreAliyuncsComDescribeRegions2015-03-01APIResponse, error) {
    var resp aliyunocs.M-KvstoreAliyuncsComDescribeRegions2015-03-01APIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
