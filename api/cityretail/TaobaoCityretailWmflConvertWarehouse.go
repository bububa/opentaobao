package cityretail

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cityretail"
)

// TaobaoCityretailWmflConvertWarehouse 同城零售完美履约转仓
// taobao.cityretail.wmfl.convert.warehouse
//
// 同城零售完美履约转仓
func TaobaoCityretailWmflConvertWarehouse(ctx context.Context, clt *core.SDKClient, req *cityretail.TaobaoCityretailWmflConvertWarehouseAPIRequest, resp *cityretail.TaobaoCityretailWmflConvertWarehouseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
