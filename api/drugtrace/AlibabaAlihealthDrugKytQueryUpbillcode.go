package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytQueryUpbillcode 通过一个码查询上游出库单
// alibaba.alihealth.drug.kyt.query.upbillcode
//
// 一个查询上游出库单号的接口。企业在扫码入库时，接口通过扫到的码判定这个码对应的上游企业所属的出库单据号
func AlibabaAlihealthDrugKytQueryUpbillcode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytQueryUpbillcodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
