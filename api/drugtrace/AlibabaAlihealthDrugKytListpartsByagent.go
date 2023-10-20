package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytListpartsByagent 物流代货主查往来单位接口
// alibaba.alihealth.drug.kyt.listparts.byagent
//
// 代理企业查询往来单位列表
func AlibabaAlihealthDrugKytListpartsByagent(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytListpartsByagentAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytListpartsByagentAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
