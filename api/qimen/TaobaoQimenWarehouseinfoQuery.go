package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenWarehouseinfoQuery 货主仓库资源查询接口
// taobao.qimen.warehouseinfo.query
//
// 货主仓库资源查询
func TaobaoQimenWarehouseinfoQuery(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenWarehouseinfoQueryAPIRequest, resp *qimen.TaobaoQimenWarehouseinfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
