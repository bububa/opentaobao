package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// Alibabaalihealthdrugkytweslistpartsbyagent 物流代货主查找往来单位接口
// alibaba.alihealth.drug.kyt.wes.listparts.byagent
//
// 代理企业查询往来单位列表
func Alibabaalihealthdrugkytweslistpartsbyagent(clt *core.SDKClient, req *drugtrace.AlibabaalihealthdrugkytweslistpartsbyagentAPIRequest, session string) (*drugtrace.AlibabaalihealthdrugkytweslistpartsbyagentAPIResponse, error) {
	var resp drugtrace.AlibabaalihealthdrugkytweslistpartsbyagentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
