package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterContractsSearch 获取合同类的服务工单信息
// tmall.servicecenter.contracts.search
//
// 获取合同类的服务工单信息
func TmallServicecenterContractsSearch(clt *core.SDKClient, req *tmallservice.TmallServicecenterContractsSearchAPIRequest, resp *tmallservice.TmallServicecenterContractsSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
