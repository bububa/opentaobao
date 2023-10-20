package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrListparts 多融查询一个企业的往来单位
// alibaba.alihealth.drug.kyt.dr.listparts
//
// 查询往来单位列表
func AlibabaAlihealthDrugKytDrListparts(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrListpartsAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDrListpartsAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytDrListpartsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
