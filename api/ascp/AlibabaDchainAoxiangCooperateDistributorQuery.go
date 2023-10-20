package ascp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascp"
)

// Alibabadchainaoxiangcooperatedistributorquery 商家关系查询分销商
// alibaba.dchain.aoxiang.cooperate.distributor.query
//
// 商家关系查询分销商
func Alibabadchainaoxiangcooperatedistributorquery(clt *core.SDKClient, req *ascp.AlibabadchainaoxiangcooperatedistributorqueryAPIRequest, session string) (*ascp.AlibabadchainaoxiangcooperatedistributorqueryAPIResponse, error) {
	var resp ascp.AlibabadchainaoxiangcooperatedistributorqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
