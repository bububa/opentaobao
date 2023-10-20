package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpchanneldistributorproductdistribute 分销商基于渠道产品铺货到商品
// alibaba.ascp.channel.distributor.product.distribute
//
// 分销商基于渠道产品铺货到商品
func Alibabaascpchanneldistributorproductdistribute(clt *core.SDKClient, req *ascpchannel.AlibabaascpchanneldistributorproductdistributeAPIRequest, session string) (*ascpchannel.AlibabaascpchanneldistributorproductdistributeAPIResponse, error) {
	var resp ascpchannel.AlibabaascpchanneldistributorproductdistributeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
