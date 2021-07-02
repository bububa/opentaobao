package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytSinglerelation 单码关联关系查询，通过一个码查询这个码下的所有子码
// alibaba.alihealth.drug.kyt.singlerelation
//
// 单码关联关系查询，通过一个码查询这个码下的所有子码。（只有做过入库的码，才能能进行查询）
func AlibabaAlihealthDrugKytSinglerelation(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytSinglerelationAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytSinglerelationAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytSinglerelationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
