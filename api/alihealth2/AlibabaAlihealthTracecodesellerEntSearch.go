package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthTracecodesellerEntSearch 查询商家信息
// alibaba.alihealth.tracecodeseller.ent.search
//
// 查询商家信息
func AlibabaAlihealthTracecodesellerEntSearch(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerEntSearchAPIRequest, resp *alihealth2.AlibabaAlihealthTracecodesellerEntSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
