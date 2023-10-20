package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytlistpartsbyagent 物流代货主查往来单位接口
// alibaba.alihealth.drug.kyt.listparts.byagent
//
// 代理企业查询往来单位列表
func Alibabaalihealthdrugkytlistpartsbyagent(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytlistpartsbyagentAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytlistpartsbyagentAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytlistpartsbyagentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
