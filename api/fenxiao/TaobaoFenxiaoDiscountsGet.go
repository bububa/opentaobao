package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoDiscountsGet 获取折扣信息
// taobao.fenxiao.discounts.get
//
// 查询折扣信息
func TaobaoFenxiaoDiscountsGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDiscountsGetAPIRequest, resp *fenxiao.TaobaoFenxiaoDiscountsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
