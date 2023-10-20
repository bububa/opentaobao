package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Tmallsupplychainchannelproductquantityget 渠道库存查询接口
// tmall.supplychain.channel.product.quantity.get
//
// 渠道库存查询接口
func Tmallsupplychainchannelproductquantityget(clt *core.SDKClient, req *fenxiao.TmallsupplychainchannelproductquantitygetAPIRequest, session string) (*fenxiao.TmallsupplychainchannelproductquantitygetAPIResponse, error) {
	var resp fenxiao.TmallsupplychainchannelproductquantitygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
