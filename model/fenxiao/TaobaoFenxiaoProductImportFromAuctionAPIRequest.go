package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductImportFromAuctionAPIRequest 导入商品生成产品 API请求
// taobao.fenxiao.product.import.from.auction
//
// 供应商选择关联店铺的前台宝贝，导入生成产品
type TaobaoFenxiaoProductImportFromAuctionAPIRequest struct {
	model.Params
	// 导入产品需要支持的交易类型：[1 代销][ 2 经销 ][3 代销和经销]
	_tradeType int64
	// 店铺宝贝id
	_auctionId int64
	// 产品线id
	_productLineId int64
}

// NewTaobaoFenxiaoProductImportFromAuctionRequest 初始化TaobaoFenxiaoProductImportFromAuctionAPIRequest对象
func NewTaobaoFenxiaoProductImportFromAuctionRequest() *TaobaoFenxiaoProductImportFromAuctionAPIRequest {
	return &TaobaoFenxiaoProductImportFromAuctionAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoProductImportFromAuctionAPIRequest) Reset() {
	r._tradeType = 0
	r._auctionId = 0
	r._productLineId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductImportFromAuctionAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.import.from.auction"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductImportFromAuctionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoProductImportFromAuctionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeType is TradeType Setter
// 导入产品需要支持的交易类型：[1 代销][ 2 经销 ][3 代销和经销]
func (r *TaobaoFenxiaoProductImportFromAuctionAPIRequest) SetTradeType(_tradeType int64) error {
	r._tradeType = _tradeType
	r.Set("trade_type", _tradeType)
	return nil
}

// GetTradeType TradeType Getter
func (r TaobaoFenxiaoProductImportFromAuctionAPIRequest) GetTradeType() int64 {
	return r._tradeType
}

// SetAuctionId is AuctionId Setter
// 店铺宝贝id
func (r *TaobaoFenxiaoProductImportFromAuctionAPIRequest) SetAuctionId(_auctionId int64) error {
	r._auctionId = _auctionId
	r.Set("auction_id", _auctionId)
	return nil
}

// GetAuctionId AuctionId Getter
func (r TaobaoFenxiaoProductImportFromAuctionAPIRequest) GetAuctionId() int64 {
	return r._auctionId
}

// SetProductLineId is ProductLineId Setter
// 产品线id
func (r *TaobaoFenxiaoProductImportFromAuctionAPIRequest) SetProductLineId(_productLineId int64) error {
	r._productLineId = _productLineId
	r.Set("product_line_id", _productLineId)
	return nil
}

// GetProductLineId ProductLineId Getter
func (r TaobaoFenxiaoProductImportFromAuctionAPIRequest) GetProductLineId() int64 {
	return r._productLineId
}

var poolTaobaoFenxiaoProductImportFromAuctionAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoProductImportFromAuctionRequest()
	},
}

// GetTaobaoFenxiaoProductImportFromAuctionRequest 从 sync.Pool 获取 TaobaoFenxiaoProductImportFromAuctionAPIRequest
func GetTaobaoFenxiaoProductImportFromAuctionAPIRequest() *TaobaoFenxiaoProductImportFromAuctionAPIRequest {
	return poolTaobaoFenxiaoProductImportFromAuctionAPIRequest.Get().(*TaobaoFenxiaoProductImportFromAuctionAPIRequest)
}

// ReleaseTaobaoFenxiaoProductImportFromAuctionAPIRequest 将 TaobaoFenxiaoProductImportFromAuctionAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoProductImportFromAuctionAPIRequest(v *TaobaoFenxiaoProductImportFromAuctionAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoProductImportFromAuctionAPIRequest.Put(v)
}
