package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDestbillCheck 直调审批
// alibaba.alihealth.drug.kyt.destbill.check
//
// 为药企提供直调单据审批操作
func AlibabaAlihealthDrugKytDestbillCheck(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDestbillCheckAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDestbillCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
