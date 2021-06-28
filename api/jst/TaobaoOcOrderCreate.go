package jst

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/jst"
)

/* 
创建OC订单 
taobao.oc.order.create

创建OC订单接口
*/
func TaobaoOcOrderCreate(clt *core.SDKClient, req *jst.TaobaoOcOrderCreateRequest, session string) (*jst.TaobaoOcOrderCreateAPIResponse, error) {
    var resp jst.TaobaoOcOrderCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
