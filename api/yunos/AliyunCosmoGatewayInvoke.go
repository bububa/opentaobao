package yunos

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/yunos"
)

/* 
alios cosmo服务调用 
aliyun.cosmo.gateway.invoke

AliOS cosmo服务分发平台对外调用接口
*/
func AliyunCosmoGatewayInvoke(clt *core.SDKClient, req *yunos.AliyunCosmoGatewayInvokeRequest, session string) (*yunos.AliyunCosmoGatewayInvokeResponse, error) {
    var resp yunos.AliyunCosmoGatewayInvokeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
