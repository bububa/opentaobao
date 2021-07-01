package tmallnr

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallnr"
)

/* 
区域零售订单获取取件码 
tmall.nr.order.logis.info

区域零售订单获取取件码，方便商家系统接入，获取取件码打印信息进行打印。
*/
func TmallNrOrderLogisInfo(clt *core.SDKClient, req *tmallnr.TmallNrOrderLogisInfoAPIRequest, session string) (*tmallnr.TmallNrOrderLogisInfoAPIResponse, error) {
    var resp tmallnr.TmallNrOrderLogisInfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
