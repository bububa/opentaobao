package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytScqyCodeprocess 关联关系处理查询
// alibaba.alihealth.drug.kyt.scqy.codeprocess
//
// 关联关系处理查询
func AlibabaAlihealthDrugKytScqyCodeprocess(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytScqyCodeprocessAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytScqyCodeprocessAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
