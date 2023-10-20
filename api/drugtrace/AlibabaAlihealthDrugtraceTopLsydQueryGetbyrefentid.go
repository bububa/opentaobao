package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentid 根据企业唯一标识查看企业详细信息
// alibaba.alihealth.drugtrace.top.lsyd.query.getbyrefentid
//
// 根据企业唯一标识查看企业详细信息
func AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentid(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryGetbyrefentidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
