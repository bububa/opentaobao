package usergrowth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/usergrowth2"
)

/* 
手淘线下拉新业务 t+8转化明细数据 
taobao.usergrowth.offline.convertion.details.eight.get

手淘线下拉新业务 给合作渠道返回t+8转化明细数据
*/
func TaobaoUsergrowthOfflineConvertionDetailsEightGet(clt *core.SDKClient, req *usergrowth2.TaobaoUsergrowthOfflineConvertionDetailsEightGetAPIRequest, session string) (*usergrowth2.TaobaoUsergrowthOfflineConvertionDetailsEightGetAPIResponse, error) {
    var resp usergrowth2.TaobaoUsergrowthOfflineConvertionDetailsEightGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
