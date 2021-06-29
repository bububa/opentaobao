package auction

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取司法拍卖最新出价数据 API请求
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

// 初始化TaobaoAuctionGovGetLatestbidRequest对象
func NewTaobaoAuctionGovGetLatestbidRequest() *TaobaoAuctionGovGetLatestbidRequest{
    return &TaobaoAuctionGovGetLatestbidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAuctionGovGetLatestbidRequest) GetApiMethodName() string {
    return "taobao.auction.gov.get.latestbid"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAuctionGovGetLatestbidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CourtName Setter
// 法院名称
func (r *TaobaoAuctionGovGetLatestbidRequest) SetCourtName(courtName string) error {
    r.courtName = courtName
    r.Set("court_name", courtName)
    return nil
}

// CourtName Getter
func (r TaobaoAuctionGovGetLatestbidRequest) GetCourtName() string {
    return r.courtName
}
// ContainChild Setter
// 死否包含下属法院
func (r *TaobaoAuctionGovGetLatestbidRequest) SetContainChild(containChild bool) error {
    r.containChild = containChild
    r.Set("contain_child", containChild)
    return nil
}

// ContainChild Getter
func (r TaobaoAuctionGovGetLatestbidRequest) GetContainChild() bool {
    return r.containChild
}
// MaxCount Setter
// 获取最新出价条数
func (r *TaobaoAuctionGovGetLatestbidRequest) SetMaxCount(maxCount int64) error {
    r.maxCount = maxCount
    r.Set("max_count", maxCount)
    return nil
}

// MaxCount Getter
func (r TaobaoAuctionGovGetLatestbidRequest) GetMaxCount() int64 {
    return r.maxCount
}
