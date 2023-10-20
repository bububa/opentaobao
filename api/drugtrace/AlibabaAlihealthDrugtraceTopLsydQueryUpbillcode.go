package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopLsydQueryUpbillcode 通过一个码，查询这个码对应的上游企业出库单的单据号
// alibaba.alihealth.drugtrace.top.lsyd.query.upbillcode
//
// 一个查询上游出库单号的接口。企业在扫码入库时，接口通过扫到的码判定这个码对应的上游企业所属的出库单据号
func AlibabaAlihealthDrugtraceTopLsydQueryUpbillcode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
