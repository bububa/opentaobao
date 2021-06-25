package auction

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/auction"
)

/* 
获取司法拍卖最新出价数据 
taobao.auction.gov.get.latestbid

获取司法拍卖最新出价数据
*/
func TaobaoAuctionGovGetLatestbid(clt *core.SDKClient, req *auction.TaobaoAuctionGovGetLatestbidRequest, session string) (*auction.TaobaoAuctionGovGetLatestbidResponse, error) {
    var resp auction.TaobaoAuctionGovGetLatestbidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
