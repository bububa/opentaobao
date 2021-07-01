package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

/* TaobaoFenxiaoProductImportFromAuction
导入商品生成产品
taobao.fenxiao.product.import.from.auction

供应商选择关联店铺的前台宝贝，导入生成产品 */
func TaobaoFenxiaoProductImportFromAuction(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductImportFromAuctionAPIRequest, session string) (*fenxiao.TaobaoFenxiaoProductImportFromAuctionAPIResponse, error) {
	var resp fenxiao.TaobaoFenxiaoProductImportFromAuctionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
