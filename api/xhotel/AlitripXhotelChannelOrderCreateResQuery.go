package xhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotel"
)

// Alitripxhotelchannelordercreateresquery 分销订单查询订单创建结果
// alitrip.xhotel.channel.order.create.res.query
//
// 针对分销渠道订单，在调用创建订单接口失败1分钟后，调用此接口，用以确认订单是否创建成功。
func Alitripxhotelchannelordercreateresquery(clt *core.SDKClient, req *xhotel.AlitripxhotelchannelordercreateresqueryAPIRequest, session string) (*xhotel.AlitripxhotelchannelordercreateresqueryAPIResponse, error) {
	var resp xhotel.AlitripxhotelchannelordercreateresqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
