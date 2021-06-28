package baichuan

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/baichuan"
)

/* 
查询当天可添加的余量 
taobao.baichuan.item.subscribe.daily.left.query

查询当天可添加的余量
*/
func TaobaoBaichuanItemSubscribeDailyLeftQuery(clt *core.SDKClient, req *baichuan.TaobaoBaichuanItemSubscribeDailyLeftQueryRequest, session string) (*baichuan.TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse, error) {
    var resp baichuan.TaobaoBaichuanItemSubscribeDailyLeftQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
