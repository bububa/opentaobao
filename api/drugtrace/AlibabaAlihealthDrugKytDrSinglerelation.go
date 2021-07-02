package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrSinglerelation 多融单码关联关系查询
// alibaba.alihealth.drug.kyt.dr.singlerelation
//
// 单码关联关系查询
func AlibabaAlihealthDrugKytDrSinglerelation(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrSinglerelationAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDrSinglerelationAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytDrSinglerelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
