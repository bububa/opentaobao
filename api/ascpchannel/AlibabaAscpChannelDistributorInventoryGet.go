package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpchanneldistributorinventoryget 链渠道中心淘外库存查询
// alibaba.ascp.channel.distributor.inventory.get
//
// 此api为淘外分销的渠道产品库存查询标准api，淘外分销商专用
func Alibabaascpchanneldistributorinventoryget(clt *core.SDKClient, req *ascpchannel.AlibabaascpchanneldistributorinventorygetAPIRequest, session string) (*ascpchannel.AlibabaascpchanneldistributorinventorygetAPIResponse, error) {
	var resp ascpchannel.AlibabaascpchanneldistributorinventorygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
