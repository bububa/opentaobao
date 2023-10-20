package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// Alitripxhotelchannelordermembertypesync 酒店分销渠道会员类型同步
// alitrip.xhotel.channel.order.membertype.sync
//
// 酒店分销渠道会员类型同步
func Alitripxhotelchannelordermembertypesync(clt *core.SDKClient, req *xhotelonlineorder.AlitripxhotelchannelordermembertypesyncAPIRequest, session string) (*xhotelonlineorder.AlitripxhotelchannelordermembertypesyncAPIResponse, error) {
	var resp xhotelonlineorder.AlitripxhotelchannelordermembertypesyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
