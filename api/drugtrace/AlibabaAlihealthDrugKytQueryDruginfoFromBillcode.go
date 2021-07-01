package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

/* AlibabaAlihealthDrugKytQueryDruginfoFromBillcode
根据单据编号查询单据明细
alibaba.alihealth.drug.kyt.query.druginfo.from.billcode

根据单据编号查询单据明细 */
func AlibabaAlihealthDrugKytQueryDruginfoFromBillcode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
