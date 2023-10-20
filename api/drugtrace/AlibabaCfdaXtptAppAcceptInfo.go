package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabacfdaxtptappacceptinfo 协同平台数据下行接口
// alibaba.cfda.xtpt.app.accept.info
//
// 协同平台数据下行接口
func Alibabacfdaxtptappacceptinfo(clt *core.SDKClient, req *drugtrace.AlibabacfdaxtptappacceptinfoAPIRequest, session string) (*drugtrace.AlibabacfdaxtptappacceptinfoAPIResponse, error) {
	var resp drugtrace.AlibabacfdaxtptappacceptinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
