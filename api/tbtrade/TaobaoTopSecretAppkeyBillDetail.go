package tbtrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTopSecretAppkeyBillDetail 服务商解密账单查询
// taobao.top.secret.appkey.bill.detail
//
// 服务商解密账单查询,分页返回所有店铺的账单，每个店铺每天仅包含两条数据，当天产生的号租费 和 当天产生的通话费，仅对90天内的账单提供SLA保障。查询账单详情请使用taobao.top.secret.bill.detail接口。
func TaobaoTopSecretAppkeyBillDetail(ctx context.Context, clt *core.SDKClient, req *tbtrade.TaobaoTopSecretAppkeyBillDetailAPIRequest, resp *tbtrade.TaobaoTopSecretAppkeyBillDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
