package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
homs员工信息核对查询服务 
alibaba.wdk.hrworkbench.cdpemps.query

给盒马可靠软件服务商Cdp系统，做非阿里编员工数据一致性核对检查
*/
func AlibabaWdkHrworkbenchCdpempsQuery(clt *core.SDKClient, req *wdk.AlibabaWdkHrworkbenchCdpempsQueryAPIRequest, session string) (*wdk.AlibabaWdkHrworkbenchCdpempsQueryAPIResponse, error) {
    var resp wdk.AlibabaWdkHrworkbenchCdpempsQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
