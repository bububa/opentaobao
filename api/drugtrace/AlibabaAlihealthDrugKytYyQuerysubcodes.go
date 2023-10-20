package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytYyQuerysubcodes 查询一个码的所有子码
// alibaba.alihealth.drug.kyt.yy.querysubcodes
//
// 单码的了码查询
func AlibabaAlihealthDrugKytYyQuerysubcodes(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytYyQuerysubcodesAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
