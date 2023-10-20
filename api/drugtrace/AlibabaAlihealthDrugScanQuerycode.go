package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugScanQuerycode 查询药监码对应的有效期和包装规格
// alibaba.alihealth.drug.scan.querycode
//
// 查询药监码对应的有效期和包装规格
func AlibabaAlihealthDrugScanQuerycode(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugScanQuerycodeAPIRequest, resp *drugtrace.AlibabaAlihealthDrugScanQuerycodeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
