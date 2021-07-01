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
type TaobaoAuctionGovGetLatestbidAPIRequest struct {
    model.Params
    // 法院名称
    _courtName   string
    // 死否包含下属法院
    _containChild   bool
    // 获取最新出价条数
    _maxCount   int64
}

// 初始化TaobaoAuctionGovGetLatestbidAPIRequest对象
func NewTaobaoAuctionGovGetLatestbidRequest() *TaobaoAuctionGovGetLatestbidAPIRequest{
    return &TaobaoAuctionGovGetLatestbidAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAuctionGovGetLatestbidAPIRequest) GetApiMethodName() string {
    return "taobao.auction.gov.get.latestbid"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAuctionGovGetLatestbidAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CourtName Setter
// 法院名称
func (r *TaobaoAuctionGovGetLatestbidAPIRequest) SetCourtName(_courtName string) error {
    r._courtName = _courtName
    r.Set("court_name", _courtName)
    return nil
}

// CourtName Getter
func (r TaobaoAuctionGovGetLatestbidAPIRequest) GetCourtName() string {
    return r._courtName
}
// ContainChild Setter
// 死否包含下属法院
func (r *TaobaoAuctionGovGetLatestbidAPIRequest) SetContainChild(_containChild bool) error {
    r._containChild = _containChild
    r.Set("contain_child", _containChild)
    return nil
}

// ContainChild Getter
func (r TaobaoAuctionGovGetLatestbidAPIRequest) GetContainChild() bool {
    return r._containChild
}
// MaxCount Setter
// 获取最新出价条数
func (r *TaobaoAuctionGovGetLatestbidAPIRequest) SetMaxCount(_maxCount int64) error {
    r._maxCount = _maxCount
    r.Set("max_count", _maxCount)
    return nil
}

// MaxCount Getter
func (r TaobaoAuctionGovGetLatestbidAPIRequest) GetMaxCount() int64 {
    return r._maxCount
}
