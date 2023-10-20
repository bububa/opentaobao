package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthTracecodesellerBillResultSearch 查询出入库单处理结果api
// alibaba.alihealth.tracecodeseller.bill.result.search
//
// 查询出入库单处理结果api
func AlibabaAlihealthTracecodesellerBillResultSearch(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerBillResultSearchAPIRequest, resp *alihealth2.AlibabaAlihealthTracecodesellerBillResultSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
