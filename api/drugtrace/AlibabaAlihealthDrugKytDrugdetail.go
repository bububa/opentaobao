package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrugdetail 查询药品详细信息
// alibaba.alihealth.drug.kyt.drugdetail
//
// 查询药品详细信息
func AlibabaAlihealthDrugKytDrugdetail(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrugdetailAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDrugdetailAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytDrugdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
