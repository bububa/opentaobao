package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaoproductimportfromauctionAPIRequest 导入商品生成产品 API请求
// taobao.fenxiao.product.import.from.auction
//
// 供应商选择关联店铺的前台宝贝，导入生成产品
type TaobaofenxiaoproductimportfromauctionAPIRequest struct {
	model.Params
	// 导入产品需要支持的交易类型：[1 代销][ 2 经销 ][3 代销和经销]
	_tradeType int64
	// 店铺宝贝id
	_auctionId int64
	// 产品线id
	_productLineId int64
}

// NewTaobaofenxiaoproductimportfromauctionRequest 初始化TaobaofenxiaoproductimportfromauctionAPIRequest对象
func NewTaobaofenxiaoproductimportfromauctionRequest() *TaobaofenxiaoproductimportfromauctionAPIRequest {
	return &TaobaofenxiaoproductimportfromauctionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaoproductimportfromauctionAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.import.from.auction"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaoproductimportfromauctionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaoproductimportfromauctionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeType is TradeType Setter
// 导入产品需要支持的交易类型：[1 代销][ 2 经销 ][3 代销和经销]
func (r *TaobaofenxiaoproductimportfromauctionAPIRequest) SetTradeType(_tradeType int64) error {
	r._tradeType = _tradeType
	r.Set("trade_type", _tradeType)
	return nil
}

// GetTradeType TradeType Getter
func (r TaobaofenxiaoproductimportfromauctionAPIRequest) GetTradeType() int64 {
	return r._tradeType
}

// SetAuctionId is AuctionId Setter
// 店铺宝贝id
func (r *TaobaofenxiaoproductimportfromauctionAPIRequest) SetAuctionId(_auctionId int64) error {
	r._auctionId = _auctionId
	r.Set("auction_id", _auctionId)
	return nil
}

// GetAuctionId AuctionId Getter
func (r TaobaofenxiaoproductimportfromauctionAPIRequest) GetAuctionId() int64 {
	return r._auctionId
}

// SetProductLineId is ProductLineId Setter
// 产品线id
func (r *TaobaofenxiaoproductimportfromauctionAPIRequest) SetProductLineId(_productLineId int64) error {
	r._productLineId = _productLineId
	r.Set("product_line_id", _productLineId)
	return nil
}

// GetProductLineId ProductLineId Getter
func (r TaobaofenxiaoproductimportfromauctionAPIRequest) GetProductLineId() int64 {
	return r._productLineId
}
