package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesListparts 查询往来单位列表
// alibaba.alihealth.drug.kyt.wes.listparts
//
// 查询往来单位列表
func AlibabaAlihealthDrugKytWesListparts(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesListpartsAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytWesListpartsAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytWesListpartsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
