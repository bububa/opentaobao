package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductImportFromAuction 导入商品生成产品
// taobao.fenxiao.product.import.from.auction
//
// 供应商选择关联店铺的前台宝贝，导入生成产品
func TaobaoFenxiaoProductImportFromAuction(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductImportFromAuctionAPIRequest, resp *fenxiao.TaobaoFenxiaoProductImportFromAuctionAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
