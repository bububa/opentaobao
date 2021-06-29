package uscesl

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/uscesl"
)

/* 
新增价签通讯AP设备 
taobao.uscesl.biz.ap.add

根据门店和ap的MAC地址新增
*/
func TaobaoUsceslBizApAdd(clt *core.SDKClient, req *uscesl.TaobaoUsceslBizApAddRequest, session string) (*uscesl.TaobaoUsceslBizApAddAPIResponse, error) {
    var resp uscesl.TaobaoUsceslBizApAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
