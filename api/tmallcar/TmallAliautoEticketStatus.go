package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoEticketStatus 查询电子凭证状态
// tmall.aliauto.eticket.status
//
// 查询天猫汽车二轮车行业门店电子凭证状态
func TmallAliautoEticketStatus(clt *core.SDKClient, req *tmallcar.TmallAliautoEticketStatusAPIRequest, resp *tmallcar.TmallAliautoEticketStatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
