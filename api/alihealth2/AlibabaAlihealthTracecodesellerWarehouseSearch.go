package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthTracecodesellerWarehouseSearch 查询仓库api
// alibaba.alihealth.tracecodeseller.warehouse.search
//
// 查询仓库api
func AlibabaAlihealthTracecodesellerWarehouseSearch(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodesellerWarehouseSearchAPIRequest, resp *alihealth2.AlibabaAlihealthTracecodesellerWarehouseSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
