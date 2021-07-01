package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

/* AlitripXhotelChannelOrderMembertypeSync
酒店分销渠道会员类型同步
alitrip.xhotel.channel.order.membertype.sync

酒店分销渠道会员类型同步 */
func AlitripXhotelChannelOrderMembertypeSync(clt *core.SDKClient, req *xhotelonlineorder.AlitripXhotelChannelOrderMembertypeSyncAPIRequest, session string) (*xhotelonlineorder.AlitripXhotelChannelOrderMembertypeSyncAPIResponse, error) {
	var resp xhotelonlineorder.AlitripXhotelChannelOrderMembertypeSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
