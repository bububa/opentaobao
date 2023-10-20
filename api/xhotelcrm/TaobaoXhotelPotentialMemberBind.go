package xhotelcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelcrm"
)

// TaobaoXhotelPotentialMemberBind 飞猪酒店商家会员绑定
// taobao.xhotel.potential.member.bind
//
// 支持互通商家发起会员绑定
func TaobaoXhotelPotentialMemberBind(clt *core.SDKClient, req *xhotelcrm.TaobaoXhotelPotentialMemberBindAPIRequest, resp *xhotelcrm.TaobaoXhotelPotentialMemberBindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
