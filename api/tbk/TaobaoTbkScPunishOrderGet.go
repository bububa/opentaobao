package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkscpunishorderget 淘宝客-服务商-处罚订单查询
// taobao.tbk.sc.punish.order.get
//
// 服务商使用。可通过该接口查询推广者账号下在处罚管理后台，可直接下载的处罚订单明细。请注意：若是基于账号(member)、媒体id(site)、推广位(adzone)、渠道方(relationid)维度的完整处罚，对应订单明细请根据处罚后台对应的处罚订单时间说明，在“推广订单明细”中筛选对应订单。
func Taobaotbkscpunishorderget(clt *core.SDKClient, req *tbk.TaobaotbkscpunishordergetAPIRequest, session string) (*tbk.TaobaotbkscpunishordergetAPIResponse, error) {
	var resp tbk.TaobaotbkscpunishordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
