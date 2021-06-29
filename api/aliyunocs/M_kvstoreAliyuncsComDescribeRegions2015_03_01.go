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
func M_kvstoreAliyuncsComDescribeRegions2015_03_01(clt *core.SDKClient, req *aliyunocs.M_kvstoreAliyuncsComDescribeRegions2015_03_01Request, session string) (*aliyunocs.M_kvstoreAliyuncsComDescribeRegions2015_03_01APIResponse, error) {
    var resp aliyunocs.M_kvstoreAliyuncsComDescribeRegions2015_03_01APIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
