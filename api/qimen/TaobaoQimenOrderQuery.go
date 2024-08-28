package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenOrderQuery 根据收件人信息查询交易单号接口
// taobao.qimen.order.query
//
// WMS 调用该接口，根据收件人信息查询平台交易订单号。
func TaobaoQimenOrderQuery(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenOrderQueryAPIRequest, resp *qimen.TaobaoQimenOrderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
