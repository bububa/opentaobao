package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesListupout 查询货主/本企业上游企业出库单据信息
// alibaba.alihealth.drug.kyt.wes.listupout
//
// 查询货主/本企业上游企业出库单据信息
func AlibabaAlihealthDrugKytWesListupout(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesListupoutAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytWesListupoutAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytWesListupoutAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
