package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenOrderprocessQuery 订单流水查询接口
// taobao.qimen.orderprocess.query
//
// ERP调用订单流水查询接口
func TaobaoQimenOrderprocessQuery(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenOrderprocessQueryAPIRequest, resp *qimen.TaobaoQimenOrderprocessQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
