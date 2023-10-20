package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Tmallsupplychainchannelproductdownshelf 产品下架
// tmall.supplychain.channel.product.downshelf
//
// 产品下架
func Tmallsupplychainchannelproductdownshelf(clt *core.SDKClient, req *fenxiao.TmallsupplychainchannelproductdownshelfAPIRequest, session string) (*fenxiao.TmallsupplychainchannelproductdownshelfAPIResponse, error) {
	var resp fenxiao.TmallsupplychainchannelproductdownshelfAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
