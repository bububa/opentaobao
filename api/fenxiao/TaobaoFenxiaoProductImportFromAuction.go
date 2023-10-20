package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaofenxiaoproductimportfromauction 导入商品生成产品
// taobao.fenxiao.product.import.from.auction
//
// 供应商选择关联店铺的前台宝贝，导入生成产品
func Taobaofenxiaoproductimportfromauction(clt *core.SDKClient, req *fenxiao.TaobaofenxiaoproductimportfromauctionAPIRequest, session string) (*fenxiao.TaobaofenxiaoproductimportfromauctionAPIResponse, error) {
	var resp fenxiao.TaobaofenxiaoproductimportfromauctionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
