package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrBillcheck 疫苗追溯验证
// alibaba.alihealth.drug.kyt.dr.billcheck
//
// 各级疾控在入库完成后，需要做追溯信息验证
func AlibabaAlihealthDrugKytDrBillcheck(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrBillcheckAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDrBillcheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
