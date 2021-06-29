package fenxiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
导入商品生成产品 API请求
taobao.fenxiao.product.import.from.auction

供应商选择关联店铺的前台宝贝，导入生成产品
*/
type TaobaoFenxiaoProductImportFromAuctionRequest struct {
    model.Params
    // 导入产品需要支持的交易类型：[1 代销][ 2 经销 ][3 代销和经销]
    tradeType   int64
    // 店铺宝贝id
    auctionId   int64
    // 产品线id
    productLineId   int64
}

// 初始化TaobaoFenxiaoProductImportFromAuctionRequest对象
func NewTaobaoFenxiaoProductImportFromAuctionRequest() *TaobaoFenxiaoProductImportFromAuctionRequest{
    return &TaobaoFenxiaoProductImportFromAuctionRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductImportFromAuctionRequest) GetApiMethodName() string {
    return "taobao.fenxiao.product.import.from.auction"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductImportFromAuctionRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TradeType Setter
// 导入产品需要支持的交易类型：[1 代销][ 2 经销 ][3 代销和经销]
func (r *TaobaoFenxiaoProductImportFromAuctionRequest) SetTradeType(tradeType int64) error {
    r.tradeType = tradeType
    r.Set("trade_type", tradeType)
    return nil
}

// TradeType Getter
func (r TaobaoFenxiaoProductImportFromAuctionRequest) GetTradeType() int64 {
    return r.tradeType
}
// AuctionId Setter
// 店铺宝贝id
func (r *TaobaoFenxiaoProductImportFromAuctionRequest) SetAuctionId(auctionId int64) error {
    r.auctionId = auctionId
    r.Set("auction_id", auctionId)
    return nil
}

// AuctionId Getter
func (r TaobaoFenxiaoProductImportFromAuctionRequest) GetAuctionId() int64 {
    return r.auctionId
}
// ProductLineId Setter
// 产品线id
func (r *TaobaoFenxiaoProductImportFromAuctionRequest) SetProductLineId(productLineId int64) error {
    r.productLineId = productLineId
    r.Set("product_line_id", productLineId)
    return nil
}

// ProductLineId Getter
func (r TaobaoFenxiaoProductImportFromAuctionRequest) GetProductLineId() int64 {
    return r.productLineId
}
