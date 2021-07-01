package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest
通过一个码，查询这个码对应的上游企业出库单的单据号 API请求
alibaba.alihealth.drugtrace.top.yljg.query.upbillcode

一个查询上游出库单号的接口。企业在扫码入库时，接口通过扫到的码判定这个码对应的上游企业所属的出库单据号 */
type AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 企业ID （一般为要查询单据的收货企业）
	_refEntId string
}

// New
