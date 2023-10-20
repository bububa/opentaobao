package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopYljgDrugtable 查询药品目录信息
// alibaba.alihealth.drugtrace.top.yljg.drugtable
//
// 查询药品目录信息
func AlibabaAlihealthDrugtraceTopYljgDrugtable(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgDrugtableAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
