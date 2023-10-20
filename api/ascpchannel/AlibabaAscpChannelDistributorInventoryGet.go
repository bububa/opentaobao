package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// AlibabaAscpChannelDistributorInventoryGet 链渠道中心淘外库存查询
// alibaba.ascp.channel.distributor.inventory.get
//
// 此api为淘外分销的渠道产品库存查询标准api，淘外分销商专用
func AlibabaAscpChannelDistributorInventoryGet(clt *core.SDKClient, req *ascpchannel.AlibabaAscpChannelDistributorInventoryGetAPIRequest, session string) (*ascpchannel.AlibabaAscpChannelDistributorInventoryGetAPIResponse, error) {
	var resp ascpchannel.AlibabaAscpChannelDistributorInventoryGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
