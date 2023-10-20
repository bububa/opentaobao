package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytScqySinglerelation 单码关联关系查询
// alibaba.alihealth.drug.kyt.scqy.singlerelation
//
// 单码关联关系查询，通过一个码查询这个码下的所有子码。（只有做过入库的码，才能能进行查询）
func AlibabaAlihealthDrugKytScqySinglerelation(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytScqySinglerelationAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytScqySinglerelationAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytScqySinglerelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
