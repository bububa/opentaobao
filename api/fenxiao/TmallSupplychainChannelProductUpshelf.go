package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Tmallsupplychainchannelproductupshelf 产品上架
// tmall.supplychain.channel.product.upshelf
//
// 上架渠道产品
func Tmallsupplychainchannelproductupshelf(clt *core.SDKClient, req *fenxiao.TmallsupplychainchannelproductupshelfAPIRequest, session string) (*fenxiao.TmallsupplychainchannelproductupshelfAPIResponse, error) {
	var resp fenxiao.TmallsupplychainchannelproductupshelfAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
