package ascpchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpchannel"
)

// Alibabaascpchanneldistributorproductdetail 获取供应链渠道中心品的详情接口（淘外分销商专用）
// alibaba.ascp.channel.distributor.product.detail
//
// 此api为淘外分销的品批量查询标准api，淘外分销商专用
func Alibabaascpchanneldistributorproductdetail(clt *core.SDKClient, req *ascpchannel.AlibabaascpchanneldistributorproductdetailAPIRequest, session string) (*ascpchannel.AlibabaascpchanneldistributorproductdetailAPIResponse, error) {
	var resp ascpchannel.AlibabaascpchanneldistributorproductdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
