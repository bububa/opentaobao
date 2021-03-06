package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterContractsSearch 获取合同类的服务工单信息
// tmall.servicecenter.contracts.search
//
// 获取合同类的服务工单信息
func TmallServicecenterContractsSearch(clt *core.SDKClient, req *tmallservice.TmallServicecenterContractsSearchAPIRequest, session string) (*tmallservice.TmallServicecenterContractsSearchAPIResponse, error) {
	var resp tmallservice.TmallServicecenterContractsSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
