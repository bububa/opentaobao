package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
查询app在设备上的安装信息 
alibaba.baichuan.aso.query

查询app在设备上的安装信息
*/
func AlibabaBaichuanAsoQuery(clt *core.SDKClient, req *baichuan.AlibabaBaichuanAsoQueryRequest, session string) (*baichuan.AlibabaBaichuanAsoQueryResponse, error) {
    var resp baichuan.AlibabaBaichuanAsoQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
