package subuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/subuser"
)

// TaobaoSubuserFullinfoGet 获取指定账户子账号的详细信息
// taobao.subuser.fullinfo.get
//
// 获取指定账户子账号的详细信息，其中包括子账号的账号信息以及员工、部门、职务信息（只能通过主账号登陆并查询属于该主账号下的某个子账号详细信息）
func TaobaoSubuserFullinfoGet(clt *core.SDKClient, req *subuser.TaobaoSubuserFullinfoGetAPIRequest, resp *subuser.TaobaoSubuserFullinfoGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
