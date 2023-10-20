package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthTracecodesellerWarehouseSearch 查询仓库api
// alibaba.alihealth.tracecodeseller.warehouse.search
//
// 查询仓库api
func AlibabaAlihealthTracecodesellerWarehouseSearch(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest, resp *alihealth2.AlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
