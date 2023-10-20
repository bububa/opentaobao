package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDestbillList 直调单据查询
// alibaba.alihealth.drug.kyt.destbill.list
//
// 为药企提供直调单据查询功能
func AlibabaAlihealthDrugKytDestbillList(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDestbillListAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDestbillListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
