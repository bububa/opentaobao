package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductImportFromAuctionAPIRequest
导入商品生成产品 API请求
taobao.fenxiao.product.import.from.auction

供应商选择关联店铺的前台宝贝，导入生成产品 */
type TaobaoFenxiaoProductImportFromAuctionAPIRequest struct {
	model.Params
	// 导入产品需要支持的交易类型：[1 代销][ 2 经销 ][3 代销和经销]
	_tradeType int64
	// 店铺宝贝id
	_auctionId int64
	// 产品线id
	_productLineId int64
}

// New
