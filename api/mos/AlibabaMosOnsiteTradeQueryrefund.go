package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosOnsiteTradeQueryrefund 退款查询
// alibaba.mos.onsite.trade.queryrefund
//
// 商户可使用该接口查询退款请求是否执行成功。
func AlibabaMosOnsiteTradeQueryrefund(clt *core.SDKClient, req *mos.AlibabaMosOnsiteTradeQueryrefundAPIRequest, resp *mos.AlibabaMosOnsiteTradeQueryrefundAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
