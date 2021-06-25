package auction

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/auction"
)

/* 
获取实时(今日)统计数据 
taobao.auction.gov.data.realtime.get

提供查询当日法院及下属法院的拍卖统计数据
*/
func TaobaoAuctionGovDataRealtimeGet(clt *core.SDKClient, req *auction.TaobaoAuctionGovDataRealtimeGetRequest, session string) (*auction.TaobaoAuctionGovDataRealtimeGetResponse, error) {
    var resp auction.TaobaoAuctionGovDataRealtimeGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
