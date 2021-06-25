package auction

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取司法拍卖最新出价数据 APIRequest
taobao.auction.gov.get.latestbid

获取司法拍卖最新出价数据
*/
type TaobaoAuctionGovGetLatestbidRequest struct {
    model.Params

    // 法院名称
    courtName   string 

    // 死否包含下属法院
    containChild   bool 

    // 获取最新出价条数
    maxCount   int64 

}

func NewTaobaoAuctionGovGetLatestbidRequest() *TaobaoAuctionGovGetLatestbidRequest{
    return &TaobaoAuctionGovGetLatestbidRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAuctionGovGetLatestbidRequest) GetApiMethodName() string {
    return "taobao.auction.gov.get.latestbid"
}

func (r TaobaoAuctionGovGetLatestbidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAuctionGovGetLatestbidRequest) SetCourtName(courtName string) error {
    r.courtName = courtName
    r.Set("court_name", courtName)
    return nil
}

func (r TaobaoAuctionGovGetLatestbidRequest) GetCourtName() string {
    return r.courtName
}

func (r *TaobaoAuctionGovGetLatestbidRequest) SetContainChild(containChild bool) error {
    r.containChild = containChild
    r.Set("contain_child", containChild)
    return nil
}

func (r TaobaoAuctionGovGetLatestbidRequest) GetContainChild() bool {
    return r.containChild
}

func (r *TaobaoAuctionGovGetLatestbidRequest) SetMaxCount(maxCount int64) error {
    r.maxCount = maxCount
    r.Set("max_count", maxCount)
    return nil
}

func (r TaobaoAuctionGovGetLatestbidRequest) GetMaxCount() int64 {
    return r.maxCount
}

