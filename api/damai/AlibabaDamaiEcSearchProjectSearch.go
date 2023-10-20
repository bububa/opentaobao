package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaiecsearchprojectsearch 大麦电商对外搜索服务
// alibaba.damai.ec.search.project.search
//
// 大麦电商对外搜索服务
func Alibabadamaiecsearchprojectsearch(clt *core.SDKClient, req *damai.AlibabadamaiecsearchprojectsearchAPIRequest, session string) (*damai.AlibabadamaiecsearchprojectsearchAPIResponse, error) {
	var resp damai.AlibabadamaiecsearchprojectsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
