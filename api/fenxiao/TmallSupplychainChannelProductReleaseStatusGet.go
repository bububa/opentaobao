package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Tmallsupplychainchannelproductreleasestatusget 产品铺货状态查询
// tmall.supplychain.channel.product.release.status.get
//
// 巴拿马战役渠道产品状态查询
func Tmallsupplychainchannelproductreleasestatusget(clt *core.SDKClient, req *fenxiao.TmallsupplychainchannelproductreleasestatusgetAPIRequest, session string) (*fenxiao.TmallsupplychainchannelproductreleasestatusgetAPIResponse, error) {
	var resp fenxiao.TmallsupplychainchannelproductreleasestatusgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
