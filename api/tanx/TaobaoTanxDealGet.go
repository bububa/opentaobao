package tanx

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// TaobaoTanxDealGet 对外部dsp提供交易id查询接口
// taobao.tanx.deal.get
//
// 对外部dsp提供交易id查询接口
func TaobaoTanxDealGet(clt *core.SDKClient, req *tanx.TaobaoTanxDealGetAPIRequest, resp *tanx.TaobaoTanxDealGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
