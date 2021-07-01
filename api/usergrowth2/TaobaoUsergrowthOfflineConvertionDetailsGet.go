package usergrowth2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/usergrowth2"
)

/* 
淘系用增线下转化明细 
taobao.usergrowth.offline.convertion.details.get

淘系用增增长-线下拉新：为渠道提供返回拉新转化数据接口
*/
func TaobaoUsergrowthOfflineConvertionDetailsGet(clt *core.SDKClient, req *usergrowth2.TaobaoUsergrowthOfflineConvertionDetailsGetAPIRequest, session string) (*usergrowth2.TaobaoUsergrowthOfflineConvertionDetailsGetAPIResponse, error) {
    var resp usergrowth2.TaobaoUsergrowthOfflineConvertionDetailsGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
