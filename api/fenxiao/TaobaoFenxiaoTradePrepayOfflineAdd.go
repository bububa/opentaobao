package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoTradePrepayOfflineAdd 线下预存款流水增加
// taobao.fenxiao.trade.prepay.offline.add
//
// 渠道分销供应商上传线下流水预存款（增加）
func TaobaoFenxiaoTradePrepayOfflineAdd(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoTradePrepayOfflineAddAPIRequest, resp *fenxiao.TaobaoFenxiaoTradePrepayOfflineAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
