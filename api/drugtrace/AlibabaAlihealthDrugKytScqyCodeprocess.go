package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytScqyCodeprocess 关联关系处理查询
// alibaba.alihealth.drug.kyt.scqy.codeprocess
//
// 关联关系处理查询
func AlibabaAlihealthDrugKytScqyCodeprocess(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytScqyCodeprocessAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
