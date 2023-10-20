package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentid 根据企业唯一标识查看企业详细信息
// alibaba.alihealth.drugtrace.top.yljg.query.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
func AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentid(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
