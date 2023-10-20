package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiEcSearchProjectSearch 大麦电商对外搜索服务
// alibaba.damai.ec.search.project.search
//
// 大麦电商对外搜索服务
func AlibabaDamaiEcSearchProjectSearch(clt *core.SDKClient, req *damai.AlibabaDamaiEcSearchProjectSearchAPIRequest, resp *damai.AlibabaDamaiEcSearchProjectSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
