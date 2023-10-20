package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// Alitripxhotelchannelordercreate 渠道分销创建订单接口
// alitrip.xhotel.channel.order.create
//
// 创建订单接口服务（如菲住等其他渠道分销提供）
func Alitripxhotelchannelordercreate(clt *core.SDKClient, req *xhotelonlineorder.AlitripxhotelchannelordercreateAPIRequest, session string) (*xhotelonlineorder.AlitripxhotelchannelordercreateAPIResponse, error) {
	var resp xhotelonlineorder.AlitripxhotelchannelordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
