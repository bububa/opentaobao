package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytCodeprocess 关联关系处理查询
// alibaba.alihealth.drug.kyt.codeprocess
//
// 关联关系处理查询
func AlibabaAlihealthDrugKytCodeprocess(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytCodeprocessAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytCodeprocessAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
