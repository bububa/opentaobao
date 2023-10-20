package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpchanneldistributorpriceget 链渠道中心淘外分销价格查询(分销商专用)
// alibaba.ascp.channel.distributor.price.get
//
// 此api为淘外分销的渠道产品价格查询标准api，淘外分销商专用
func Alibabaascpchanneldistributorpriceget(clt *core.SDKClient, req *ascpchannel.AlibabaascpchanneldistributorpricegetAPIRequest, session string) (*ascpchannel.AlibabaascpchanneldistributorpricegetAPIResponse, error) {
	var resp ascpchannel.AlibabaascpchanneldistributorpricegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
