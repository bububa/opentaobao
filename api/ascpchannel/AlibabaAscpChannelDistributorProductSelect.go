package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpchanneldistributorproductselect 供应链渠道中心品的选品接口（淘外分销商专用）
// alibaba.ascp.channel.distributor.product.select
//
// 此api为淘外分销的品的选品标准api，淘外分销商专用
func Alibabaascpchanneldistributorproductselect(clt *core.SDKClient, req *ascpchannel.AlibabaascpchanneldistributorproductselectAPIRequest, session string) (*ascpchannel.AlibabaascpchanneldistributorproductselectAPIResponse, error) {
	var resp ascpchannel.AlibabaascpchanneldistributorproductselectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
