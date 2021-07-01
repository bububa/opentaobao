package logistic

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/logistic"
)

/* 
预采购服务包查询接口 
alibaba.ele.fengniao.service.package.query

查询门店所在经纬度可用服务包的接口
*/
func AlibabaEleFengniaoServicePackageQuery(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoServicePackageQueryAPIRequest, session string) (*logistic.AlibabaEleFengniaoServicePackageQueryAPIResponse, error) {
    var resp logistic.AlibabaEleFengniaoServicePackageQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
