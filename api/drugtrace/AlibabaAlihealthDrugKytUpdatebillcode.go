package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytUpdatebillcode 零售修改出入库单追溯码
// alibaba.alihealth.drug.kyt.updatebillcode
//
// 零售修改出入库单追溯码
func AlibabaAlihealthDrugKytUpdatebillcode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytUpdatebillcodeAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytUpdatebillcodeAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytUpdatebillcodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
