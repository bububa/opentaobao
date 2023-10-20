package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoTradePrepayOfflineReduce 渠道分销供应商上传线下流水预存款（减少）
// taobao.fenxiao.trade.prepay.offline.reduce
//
// 渠道分销供应商上传线下流水预存款（减少）
func TaobaoFenxiaoTradePrepayOfflineReduce(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoTradePrepayOfflineReduceAPIRequest, resp *fenxiao.TaobaoFenxiaoTradePrepayOfflineReduceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
