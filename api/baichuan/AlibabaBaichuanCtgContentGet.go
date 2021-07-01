package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
百川内容平台内容获取 
alibaba.baichuan.ctg.content.get

百川内容平台内容获取
*/
func AlibabaBaichuanCtgContentGet(clt *core.SDKClient, req *baichuan.AlibabaBaichuanCtgContentGetAPIRequest, session string) (*baichuan.AlibabaBaichuanCtgContentGetAPIResponse, error) {
    var resp baichuan.AlibabaBaichuanCtgContentGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
