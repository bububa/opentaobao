package tanx

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// TaobaoTanxBiddingrefusesGet tanx竞价失败反馈api
// taobao.tanx.biddingrefuses.get
//
// 竞价失败反馈根据创意id查询API提供
func TaobaoTanxBiddingrefusesGet(clt *core.SDKClient, req *tanx.TaobaoTanxBiddingrefusesGetAPIRequest, resp *tanx.TaobaoTanxBiddingrefusesGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
