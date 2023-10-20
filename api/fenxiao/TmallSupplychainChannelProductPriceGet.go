package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Tmallsupplychainchannelproductpriceget 渠道价格查询接口
// tmall.supplychain.channel.product.price.get
//
// 渠道价格查询接口
func Tmallsupplychainchannelproductpriceget(clt *core.SDKClient, req *fenxiao.TmallsupplychainchannelproductpricegetAPIRequest, session string) (*fenxiao.TmallsupplychainchannelproductpricegetAPIResponse, error) {
	var resp fenxiao.TmallsupplychainchannelproductpricegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
